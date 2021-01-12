package main

import (
	"fmt"
	"strings"
)

func main() {
	kk := "pwdxwwnn"
	m := make(map[int]string)
	for s, v := range strings.Split(kk, "") {
		m[s] = v
	}

	//字串長度

	//比較不重複的長度

	count := len(m)
	fmt.Println(count)

	checklen := 1

	start := 0
	end := 0

	isRun := true
	//m[0] ~ m[1],m[1] ~ m[2],m[2] ~ m[3]
	for isRun {
		//fmt.Printf("進來checklen:%d \n", checklen)
		fmt.Printf("start:%d end+checklen:%d\n", start, end+checklen)
		//fmt.Println(checkMapNoRepeat(m, start+checklen, end+checklen))

		if checkMapNoRepeat(m, start, end+checklen) {
			fmt.Printf("最大長度:%d \n", checklen+1)
			resultStr := ""
			for k := start; k <= end+checklen; k++ {
				resultStr = resultStr + m[k]
			}
			fmt.Printf("resultStr:%v \n", resultStr)

		}
		start++
		end++
		if checklen >= count {
			isRun = false
		}

		if end+checklen >= count {
			start = 0
			end = 0
			checklen++
			continue
		}

	}

	//是否有重複

	fmt.Println(m)

}

func campare(v1 string, v2 string) bool {
	return v1 == v2
}

func checkMapNoRepeat(m map[int]string, start int, end int) bool {

	for i := start - 1; i <= end; i++ {
		for j := start; j <= end; j++ {
			if i != j {
				if !campare(m[i], m[j]) {
					continue
				} else {
					return false
				}
			}

		}
	}
	return true
}
