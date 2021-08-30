package main

import "fmt"

type People struct {
	name   string
	age    int
	gender string
}

//type 1 is better than type2
type android struct {
	People
	phoneType string
}

//typ2
type android2 struct {
	People    People
	phoneType string
}

func (p *People) talk() {
	fmt.Println("hi my name is ", p.name)
}

func main() {
	p := People{"sheldon", 18, "male"}
	phone := new(android)
	phone.People = p
	phone.talk()
	p.talk()

	phone2 := new(android2)
	phone2.People = p
	phone2.People.talk()

}
