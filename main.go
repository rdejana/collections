package main

import (
	"collections/lists"
	"fmt"
)

func main() {
	list := lists.NewListListV1[int]()
	list.Append(0)
	list.Append(1)
	list.Append(2)

	list.Print()
	fmt.Printf("%d\n", list.Length())
	v, ok := list.Get(0)
	fmt.Printf("%d %t\n", v, ok)
	v, ok = list.Get(1)
	fmt.Printf("%d %t\n", v, ok)
	v, ok = list.Get(2)
	fmt.Printf("%d %t\n", v, ok)
	v, ok = list.Get(3)
	fmt.Printf("%d %t\n", v, ok)

}
