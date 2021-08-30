package main

import "fmt"

func main() {

	// var nums = []float64{12, 13, 45, 56}
	// fmt.Println(averge(nums))
	// x, y := f()
	// fmt.Println(x, y)

	// xs := []int{1, 2, 3}
	// fmt.Println(add(xs...))

	doublex := 2
	doubleNum := func() int {
		doublex = doublex * 2
		return doublex
	}

	addOneNum := addOne()

	fmt.Println(addOneNum())
	fmt.Println(addOneNum())
	fmt.Println(doubleNum())
}

func addOne() func() int {
	total := 0
	return func() (x int) {
		total = total + 1
		return total
	}
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func f() (int, int) {
	return 5, 6
}

func averge(avg []float64) float64 {
	resultnum := 0.0
	for _, num := range avg {
		resultnum += num
	}
	resultnum = resultnum / float64(len(avg))
	return resultnum
}
