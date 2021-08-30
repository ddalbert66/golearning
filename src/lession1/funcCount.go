package main

import "fmt"

func main() {
	floatArray := []float64{77, 88, 99, 33, 44}
	var total float64 = 0.0

	for _, fl := range floatArray {
		total += fl
	}

	fmt.Println(total / float64(len(floatArray)))
}
