package main

import "fmt"

type persion struct {
	name string
	age  int
	eat  string
	sex  string
}

func main() {

	sheldon := persion{
		name: "sheldon",
		age:  18,
		eat:  "food",
		sex:  "male",
	}

	fmt.Println(sheldon.name)
	sheldon.eatFood()
	fmt.Println(sheldon.name)
	sheldon.eatFood2()
	fmt.Println(sheldon.name)

}

func (this persion) eatFood() {
	fmt.Println(this.name, "eat rise")
	this.name = "ricky"
}

func (this *persion) eatFood2() {

	fmt.Println(this.name, "eat rise")
	this.name = "ricky"
}
