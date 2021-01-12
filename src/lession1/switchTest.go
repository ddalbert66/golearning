package main

import "fmt"

func grade(score int) string {
	g := ""
	switch {
	case score > 100 || score < 0:
		panic(fmt.Sprintf("error Score %d", score))
	case score < 60:
		g = "F"
	case score > 60:
		g = "D"
	case score > 70:
		g = "C"
	case score > 80:
		g = "B"
	case score > 90:
		g = "A"
	}
	return g

}

//func main() {

//	fmt.Println(grade(60))
//	fmt.Println(grade(80))
//	fmt.Println(grade(40))
//	fmt.Println(grade(101))
//}
