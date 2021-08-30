package main

import "fmt"

func main() {
	fmt.Println(fibFun(8))
}

func fibFun(n int) int {
	result := 0
	n0 := 0
	n1 := 1
	for i := 0; i <= n; i++ {
		if i == 0 {
			result = n0
		} else if i == 1 {
			result = n1
		} else {
			if i%2 == 1 {
				n0 = n0 + n1

				result = n0
			} else if i%2 == 0 {
				n1 = n0 + n1

				result = n1
			}
		}

	}

	return result
}
