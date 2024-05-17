package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"log"
	"time"
)

type Ticket struct {
	From string `json:"From"`
	Message string `json:"Message"`
	Type int `json:"Type"`
	To uint64 `json:"To"`
}

// 定义结构体来匹配JSON中的"TurnAuthServers"数组
type TurnAuthServer struct {
	Username string   `json:"Username"`
	Password string   `json:"Password"`
	Urls     []string `json:"Urls"`
}

// 定义顶层的结构体
type Config struct {
	ExpirationInSeconds int             `json:"ExpirationInSeconds"`
	TurnAuthServers     []TurnAuthServer `json:"TurnAuthServers"`
}


func main() {
	messageChan := make(chan []byte)
	go func() {
		c, _, err := websocket.DefaultDialer.Dial("ws://10.212.14.201:8899/33333/2147594971/TvLsOb7azH9hG01c22D0GQ==/ijYBBJpXL6JO5Ih2axFDnA==", nil)
		if err != nil {
			log.Println("dial:", err)
			return
		}
		defer c.Close()

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			messageChan <- message
			close(messageChan)
		}
	}()
	s := &Config{}
	for message := range messageChan {
		t := &Ticket{}
		err := json.Unmarshal(message, t)
		if err != nil {
			log.Println("json unmarshal:", err)
		} else {
			fmt.Println(t)
		}

		err = json.Unmarshal([]byte(t.Message),s )
		if err != nil {
			log.Println("json unmarshal:", err)
		} else {
			fmt.Println(s)
		}
	}
	// 创建一个新的PeerConnection
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{s.TurnAuthServers[0].Urls[0]}, // Google 公共 STUN 服务器
			},
			{
				URLs:           []string{s.TurnAuthServers[0].Urls[1]}, // TURN 服务器地址
				Username:       "threelines",                 // TURN 服务器用户名
				Credential:     "threelinestest@163",               // TURN 服务器密码
				CredentialType: webrtc.ICECredentialTypePassword, // 认证类型
			},
		},
		SDPSemantics:       webrtc.SDPSemanticsUnifiedPlan, // SDP 语义设置为 Unified Plan
		ICETransportPolicy: webrtc.ICETransportPolicyAll,   // ICE 传输策略
	}
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	// 创建一个DataChannel
	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		panic(err)
	}
	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		fmt.Printf("ICE connection state has changed to %s\n", state.String())
	})
	// 当DataChannel打开时
	dataChannel.OnOpen(func() {
		fmt.Println("Data channel is open")
		for {
			time.Sleep(5 * time.Second)
			// 发送消息
			text := "Hello from Golang WebRTC client"
			fmt.Printf("Sending: %s\n", text)
			if err := dataChannel.SendText(text); err != nil {
				panic(err)
			}
		}
	})
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		d.OnOpen(func() {
			fmt.Printf("Data channel '%s'-'%d' open.\n", d.Label(), d.ID())
		})

		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	})
	// 当DataChannel收到消息时
	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Received: %s\n", string(msg.Data))
	})

	// 阻塞等待命令行输入，以保持程序运行
	select {}
}
