package main

import "fmt"

func lengthOfRepeatingSubStr(s string) int {
	lastOccoured := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastOccoured[ch] >= start {
			start = lastOccoured[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccoured[ch] = i
	}
	return maxLength

}

func main() {
	lastOccoured := make(map[byte]int)
	s := "xyxxxxxxxxxx"
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		fmt.Printf("lastOccoured[ch] 1:%d \n", lastOccoured[ch])
		if lastOccoured[ch] >= start {
			start = lastOccoured[ch] + 1

			fmt.Println(start)
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccoured[ch] = i
		fmt.Printf("start:%d \n", start)
		fmt.Printf("ch:%d \n", ch)
		fmt.Printf("lastOccoured[ch] 2:%d \n", lastOccoured[ch])
		fmt.Printf("maxLength:%d \n", maxLength)
		fmt.Println()

	}

	fmt.Println(lengthOfRepeatingSubStr("xykjkjkjsssss"))
}
