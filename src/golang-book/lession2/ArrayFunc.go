package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	fmt.Println(averageNum())
}

func averageNum() float64 {

	var nums [5]float64
	nums[0] = 99
	nums[1] = 88
	nums[2] = 87
	nums[3] = 85
	nums[4] = 55

	var result float64 = 0
	for _, num := range nums {
		result = result + num
	}
	return result / float64(len(nums))

}
