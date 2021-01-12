package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("array[2:4] = ", array[2:4])
	fmt.Println("array[:4] = ", array[:4])
	fmt.Println("array[2:] = ", array[2:])
	fmt.Println("array[:] = ", array[:])
	s1 := array[2:4]
	s2 := array[:4]
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(array)

	fmt.Println("reslice")
	s3 := array[:]
	fmt.Println(s3)
	s3 = s3[2:3]
	fmt.Println(s3)
	s3 = s3[0:1]
	fmt.Println(s3)

	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d", s3, len(s3), cap(s3))

	s4 := array[2:6]
	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)

	fmt.Println("s4=%v,s5=%v,s6=%v,s7=%v", s4, s5, s6, s7)
	fmt.Println("array=%v", array)
}
