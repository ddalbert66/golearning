package main

import (
	"fmt"
	"strconv"
)

type race interface {
	run() int
}

type Person struct {
	name string
	rank int
}

func (p Person) run() int {
	return p.rank
}

func Allrun(players []Person) string {
	var exl string = ""
	for _, player := range players {
		exl = exl + ",跑者:" + player.name + ",排名:" + string(player.run())
	}
	return exl
}

func main() {

	players := []Person{}
	for i := 0; i < 47; i++ {
		n := Person{name: "runer" + strconv.Itoa(i), rank: i}
		players = append(players, n)
	}
	Allrun(players)

	for _, player := range players {
		fmt.Println("name:", player.name, "run:", player.run())
	}

}
