package main

import "fmt"

type IAttack interface {
	Attack()
}

type HumanLowLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是:", a.name, ",等級為:", a.level)
}

func main() {
	var player IAttack

	lowLevel := HumanLowLevel{
		name:  "David",
		level: 1,
	}

	lowLevel.Attack()

	player = &lowLevel
	player.Attack()
}
