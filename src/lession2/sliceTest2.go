package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len(s)=%d,cap(s)=%d \n", len(s), cap(s))
}

func main() {
	var nuils []int
	nuils = nil
	var nilMap = make(map[int][]string)
	nilMap = nil

	fmt.Printf("slice nil v:%v ,map nil v:%v \n", nuils, nilMap)
	fmt.Println(nuils, nilMap)
	var s []int

	for i := 1; i < 10; i++ {
		printSlice(s)
		s = append(s, i*2+1)
	}
	fmt.Println(s)

	fmt.Println("copy slice")
	s1 := make([]int, 5, 6)
	s2 := []int{3, 4, 5, 6}
	fmt.Printf("s1=%v,s2=%v", s1, s2)
	copy(s1, s2)
	fmt.Printf("s1=%v,s2=%v \n", s1, s2)

	fmt.Println("delete slice middle")
	s1 = append(s1[:3], s1[4:]...)
	fmt.Printf("s1=%v \n", s1)
	printSlice(s1)

	fmt.Println("pop slice front")
	front := s1[0]
	s1 = s1[1:]
	fmt.Printf("s1=%v \n", front)
	fmt.Printf("s1=%v \n", s1)
	printSlice(s1)

	fmt.Println("pop slice end")
	lens1 := len(s1)
	end := s1[lens1-1]
	s1 = s1[:lens1-1]
	fmt.Printf("s1=%v \n", end)
	fmt.Printf("s1=%v \n", s1)
	printSlice(s1)
}
