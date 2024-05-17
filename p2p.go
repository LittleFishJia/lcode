package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"os"

	"strconv"
	"time"

	"github.com/pion/webrtc/v3"
)

// 信令服消息
type Message struct {
	Type    int    `json:"Type"`
	To      uint64 `json:"To"`
	From    string `json:"From"`
	Message string `json:"Message"`
}
type Candidate struct {
	Candidate string `json:"candidate"`
}

func main() {
	clientId := os.Args[1]
	otherClientId, _ := strconv.Atoi(os.Args[2])
	role := os.Args[3] // Add role argument: "offer" or "answer"

	c, _, err := websocket.DefaultDialer.Dial("ws://10.212.14.201:8899/"+clientId+"/2147594970/TvLsOb7azH9hG01c22D0GQ==/ijYBBJpXL6JO5Ih2axFDnA==", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	msg := make(chan string)

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs:       []string{"stun:42.186.72.12:3478", "turn:42.186.72.12:5349"},
				Username:   "threelines",
				Credential: "threelinestest@163",
			},
		},
	}
	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				//log.Println("read:", err)
				return
			}
			log.Printf("received: %s", message)
			tmp := &Message{}
			err = json.Unmarshal(message, tmp)
			if err != nil {
				return
			}
			if tmp.To != 0 {
				//log.Printf("received: %s", tmp.Message)
				msg <- tmp.Message
			}

		}
	}()

	dataChannel, err := pc.CreateDataChannel("chat", nil)
	if err != nil {
		panic(err)
	}
	var sdp string
	// 设置数据通道事件处理程序
	dataChannel.OnOpen(func() {
		fmt.Println("Data channel opened")
		if role == "offer" {
			go func() {
				for {
					n := rand.Intn(5)
					time.Sleep(time.Second * time.Duration(n))
					message := fmt.Sprintf("ping %d", time.Now().UnixNano()/1e6)
					dataChannel.SendText(message)
				}
			}()
		}
	})
	pc.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())
		// Register text message handling
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
			if role == "answer" {
				message := fmt.Sprintf("pong %d", time.Now().UnixNano()/1e6)
				dataChannel.SendText(message)
			}
			if role == "offer" {
				if string(msg.Data)[:4] == "pong" {
					timestamp, _ := strconv.ParseInt(string(msg.Data[5:]), 10, 64)
					elapsed := time.Now().UnixNano()/1e6 - timestamp
					fmt.Printf("RTT: %d ms\n", elapsed)
				}
			}
		})
	})

	dataChannel.OnError(func(err error) {
		fmt.Println("Data channel error:", err)
	})
	pc.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		log.Printf("Peer Connection State has changed: %s", state)
	})
	pc.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed to %s\n", state.String())
	})

	pc.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			candidateJSON, _ := json.Marshal(candidate.ToJSON())
			// 发送candidate到信令服务器
			message := struct {
				Type    int
				Message string
				To      uint64
			}{
				Type:    1,
				Message: string(candidateJSON),
				To:      uint64(otherClientId),
			}
			msgbuf, err := json.Marshal(message)
			err = c.WriteMessage(websocket.TextMessage, msgbuf)
			if err != nil {
				log.Println("write:", err)
				return
			} else {
				log.Println("send candiateJson", string(msgbuf))
			}
		}
	})
	if role == "offer" {
		offer, err := pc.CreateOffer(nil)
		if err != nil {
			panic(err)
		}
		err = pc.SetLocalDescription(offer)
		if err != nil {
			panic(err)
		}

		offerJSON, err := json.Marshal(offer)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			ticker := time.NewTicker(time.Second * 5)
			defer ticker.Stop()
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					message := struct {
						Type    int
						Message string
						To      uint64
					}{
						Type:    1,
						Message: string(offerJSON),
						To:      uint64(otherClientId),
					}
					msgbuf, err := json.Marshal(message)
					if err == nil {
						err = c.WriteMessage(websocket.TextMessage, msgbuf)
						if err != nil {
							log.Println("write:", err)
							return
						}
					} else {
						log.Println("write:", err)
					}
				}
			}
		}()

		first := true
		for message := range msg {
			sdp = message
			sdpStruct := struct {
				Type string `json:"type"`
				SDP  string `json:"sdp"`
			}{}
			json.Unmarshal([]byte(sdp), &sdpStruct)
			fmt.Println(sdpStruct)
			if sdpStruct.Type == "answer" && first {
				first = false
				// 处理 Answer
				remoteDesc := webrtc.SessionDescription{
					Type: webrtc.SDPTypeAnswer,
					SDP:  sdpStruct.SDP,
				}
				log.Printf("SignalingState signaling state for creating an answer: %s", pc.SignalingState())
				// 设置远程描述
				if err := pc.SetRemoteDescription(remoteDesc); err != nil {
					log.Fatalf("Failed to set remote description: %v", err)
					return
				}

			} else {
				tmp1 := &Candidate{}
				err = json.Unmarshal([]byte(message), tmp1)
				var candidateInit webrtc.ICECandidateInit
				if err := json.Unmarshal([]byte(message), &candidateInit); err != nil {
					log.Printf("Failed to unmarshal ICE candidate: %s", err)
					return
				}

				// 添加ICE候选到PeerConnection
				if err := pc.AddICECandidate(candidateInit); err != nil {
					log.Printf("Failed to add ICE candidate: %s", err)
					return
				}
			}
		}

	} else if role == "answer" {
		first := true
		fmt.Println("Test")
		for message := range msg {
			sdp = message
			sdpStruct := struct {
				Type string `json:"type"`
				SDP  string `json:"sdp"`
			}{}

			err = json.Unmarshal([]byte(sdp), &sdpStruct)
			if err != nil {
				log.Fatalf("Error unmarshalling JSON: %v", err)
			}
			if sdpStruct.Type == "offer" && first {
				first = false
				remoteDesc := webrtc.SessionDescription{
					Type: webrtc.SDPTypeOffer,
					SDP:  sdpStruct.SDP,
				}
				log.Printf("SignalingState signaling state for creating an answer: %s", pc.SignalingState())
				// 设置远程描述
				err = pc.SetRemoteDescription(remoteDesc)
				if err != nil {
					log.Fatalf("Failed to set remote description: %v", err)
					return
				}

				// 确保当前状态允许创建 Answer
				if pc.SignalingState() == webrtc.SignalingStateHaveRemoteOffer {
					answer, err := pc.CreateAnswer(nil)
					if err != nil {
						log.Fatalf("Failed to create answer: %v", err)
						return
					}

					err = pc.SetLocalDescription(answer)
					if err != nil {
						log.Fatalf("Failed to set local description: %v", err)
						return
					}
					fmt.Printf("Answer SDP:\n%s\n", answer.SDP)

					answerJSON, err := json.Marshal(answer)
					if err != nil {
						log.Fatal(err)
					}

					message := struct {
						Type    int
						Message string
						To      uint64
					}{
						Type:    1,
						Message: string(answerJSON),
						To:      uint64(otherClientId),
					}
					msgbuf, err := json.Marshal(message)
					if err == nil {
						err = c.WriteMessage(websocket.TextMessage, msgbuf)
						if err != nil {
							log.Println("write:", err)
							return
						}
					} else {
						log.Println("write:", err)
					}
				} else {
					log.Printf("Invalid signaling state for creating an answer: %s", pc.SignalingState())
				}
				pc.OnICECandidate(func(candidate *webrtc.ICECandidate) {
					if candidate != nil {
						candidateJSON, _ := json.Marshal(candidate.ToJSON())
						// 发送candidate到信令服务器
						message := struct {
							Type    int
							Message string
							To      uint64
						}{
							Type:    1,
							Message: string(candidateJSON),
							To:      uint64(otherClientId),
						}
						msgbuf, err := json.Marshal(message)
						err = c.WriteMessage(websocket.TextMessage, msgbuf)
						if err != nil {
							log.Println("write:", err)
							return
						} else {
							log.Println("send candiateJson", string(msgbuf))
						}
					}
				})
			} else {
				tmp1 := &Candidate{}
				err = json.Unmarshal([]byte(message), tmp1)
				var candidateInit webrtc.ICECandidateInit
				if err := json.Unmarshal([]byte(message), &candidateInit); err != nil {
					log.Printf("Failed to unmarshal ICE candidate: %s", err)
					return
				}

				// 添加ICE候选到PeerConnection
				if err := pc.AddICECandidate(candidateInit); err != nil {
					log.Printf("Failed to add ICE candidate: %s", err)
					return
				}
			}
		}

	}
	select {
	case <-done:
		log.Println("Exiting due to WebSocket close/error.")
	}
}
func checkConnectionType(pc *webrtc.PeerConnection) {
	stats := pc.GetStats()

	for _, s := range stats {

		switch v := s.(type) {
		case webrtc.ICECandidatePairStats:

			if v.State == webrtc.StatsICECandidatePairStateSucceeded && v.Nominated {
				localCandidate := stats[v.LocalCandidateID].(webrtc.ICECandidateStats)
				remoteCandidate := stats[v.RemoteCandidateID].(webrtc.ICECandidateStats)
				fmt.Printf("Local Candidate Type: %s Ip: %s port:%d\n", localCandidate.CandidateType, localCandidate.IP, localCandidate.Port)
				fmt.Printf("Remote Candidate Type:  %s Ip: %s port:%d\n", remoteCandidate.CandidateType, remoteCandidate.IP,
					remoteCandidate.Port)

				// 判断连接类型
				if localCandidate.CandidateType == webrtc.ICECandidateTypeRelay || remoteCandidate.CandidateType == webrtc.ICECandidateTypeRelay {
					fmt.Println("The connection is established via TURN server.")
				} else {
					fmt.Println("The connection is a direct P2P connection.")
				}
			}
		}
	}
}

// 在主函数或适当的地方调用handleOffer
