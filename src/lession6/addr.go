package main

func adder() func(value int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

// func main() {
// 	a := adder2(0)
// 	fmt.Printf("xxxx%t \n", a)
// 	for i := 0; i < 10; i++ {
// 		var s int
// 		s, a = a(i)
// 		fmt.Printf("0+1+.....+ %d = %d \n", s, a)
// 	}
// }
