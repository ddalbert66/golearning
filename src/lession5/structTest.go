package main

import "fmt"

type persion struct {
	name string
	age  int
	eat  string
	sex  string
}

type runner interface {
	drink() string
}

type boss struct {
	persion
}

func main() {

	b := new(boss)
	b.name = "boss"
	fmt.Println(drinkNow(b))

}

func (this boss) drink() string {
	return this.name + " drink water"

}
func drinkNow(r runner) string {
	return r.drink()

}
