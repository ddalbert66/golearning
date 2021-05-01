package main

import "fmt"

func main() {
	names := []string{"台北", "台中", "高雄", "台南"}

	var names2 [3]string

	names2[0] = names[0]
	names2[1] = names[1]
	names2[2] = names[2]

	fmt.Println(names2)
	fmt.Println(names[0:4])
	names3 := names[0:4]

	names2[1] = "HEllo"
	names3[1] = "Hello"
	fmt.Println("修改後")
	fmt.Println(names2)
	fmt.Println(names[0:4])

	sub1 := "helloworld"[2:]
	fmt.Println(sub1)

	namesCopy := make([]string, len(names))
	copy(namesCopy, names[:])

	fmt.Println(namesCopy)
}
