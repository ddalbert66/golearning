package main

import "fmt"

func arrayCopy(a [5]int) {
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func arrayCopy2(a []int) {

	fmt.Println(a)

}

func main() {
	var array1 [5]int
	array2 := [3]int{2, 4, 6}
	array3 := [...]int{7, 8, 9, 10, 11}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	arrayCopy(array3)
	//for i := 0; i < len(array3); i++ {
	//	fmt.Println(array3[i])
	//}

	for i, v := range array3 {
		arrayCopy2(v)
	}
}
