package main

import (
	"data-structures/list"
	"fmt"
)

func main() {
	list := list.New()
	list.AddValue(1)
	fmt.Println("Hello from ", list)
}
