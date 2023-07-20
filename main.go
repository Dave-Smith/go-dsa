package main

import (
	"fmt"
)

func main() {
	num := randomNum(100)
	list := CreateList(100)
	list = RemoveFromList(list, num)
	fmt.Println("Number:", num)
	left, right := list.Split()
	fmt.Println("left list")
	PrintList(left)
	fmt.Println("right list")
	PrintList(right)
}
