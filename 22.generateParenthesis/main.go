package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var tmpMap = map[string]bool{"()": true}
	for i := 2; i <= n; i++ {
		tmpMap = backtrace(tmpMap)
	}
	var result []string
	for s, _ := range tmpMap {
		result = append(result, s)
	}
	return result
}

func backtrace(tmpMap map[string]bool) map[string]bool {
	newTmpMap := map[string]bool{}
	for a, _ := range tmpMap {
		for i := 0; i < len(a); i++ {
			str := a[:i] + "()" + a[i:]
			_, ok := newTmpMap[str]
			if !ok {
				newTmpMap[str] = true
			}
		}
	}

	return newTmpMap

}

func generateParenthesis1(n int) []string {
	trials := []*trial{}
	completeTrials := []*trial{}
	trials = append(trials, &trial{
		usedLefts:  0,
		needRights: 0,
		thesis:     "",
	})

	for len(trials) > 0 {
		newTrials := []*trial{}
		for _, oldTrial := range trials {
			if oldTrial.usedLefts < n {
				// 尝试加上一个左
				newTrial := &trial{
					usedLefts:  oldTrial.usedLefts + 1,
					needRights: oldTrial.needRights + 1,
					thesis:     oldTrial.thesis + "(",
				}
				newTrials = append(newTrials, newTrial)
			}
			if oldTrial.needRights > 0 {
				// 尝试加上一个右
				newTrial := &trial{
					usedLefts:  oldTrial.usedLefts,
					needRights: oldTrial.needRights - 1,
					thesis:     oldTrial.thesis + ")",
				}
				newTrials = append(newTrials, newTrial)
			}
			if oldTrial.usedLefts == n && oldTrial.needRights == 0 {
				completeTrials = append(completeTrials, oldTrial)
			}
		}
		trials = newTrials
	}
	// 返回所有尝试
	result := make([]string, len(completeTrials))
	for i, oldTrial := range completeTrials {
		result[i] = oldTrial.thesis
	}
	return result
}

type trial struct {
	usedLefts  int
	needRights int
	thesis     string
}
