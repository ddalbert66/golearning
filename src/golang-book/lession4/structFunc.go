package main

import (
	"fmt"
	"math"
)

type circle struct {
	name      string
	testPoint *string
	x         float64
	y         float64
	z         float64
}

func main() {
	var cr circle
	fmt.Println("x:", cr.x, ",y:", cr.y, ",z:", cr.z, ",testPoint:", cr.testPoint, ",name:", cr.name)
	cr1 := new(circle)
	fmt.Println("x:", cr1.x, ",y:", cr1.y, ",z:", cr1.z, ",testPoint:", cr1.testPoint, ",name:", cr1.name)
	cr2 := circle{x: 10, y: 10, z: 12}
	fmt.Println("x:", cr2.x, ",y:", cr2.y, ",z:", cr2.z, ",testPoint:", cr2.testPoint, ",name:", cr2.name)
	cr3 := circle{"name", nil, 10, 10, 12}
	fmt.Println("x:", cr3.x, ",y:", cr3.y, ",z:", cr3.z, ",testPoint:", cr3.testPoint, ",name:", cr3.name)

	fmt.Println(circleArea(cr2))

}

func circleArea(cr circle) float64 {
	return math.Pi * cr.z * cr.z
}
