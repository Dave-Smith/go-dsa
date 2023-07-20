package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head, tail *Node
}

func (list *LinkedList) Add(value int) *LinkedList {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
	}
	list.tail = newNode

	return list
}

func randomNum(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func CreateList(size int) *LinkedList {
	list := &LinkedList{}
	for i := 0; i < size; i++ {
		list.Add(i)
	}

	return list
}

func PrintList(list LinkedList) {
	node := list.head
	for node != nil {
		fmt.Printf("%d,", node.value)
		node = node.next
	}
}

func CreateSlice(length int) []int {
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice = append(slice, i)
	}
	return slice
}

func RemoveFromList(list *LinkedList, index int) *LinkedList {
	node := list.head

	for i := 1; i < index; i++ {
		node = node.next
	}

	fmt.Println("Removing node with index ", node.value)
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node = nil

	return list
}

func (list *LinkedList) Split() (left LinkedList, right LinkedList) {
	slow := list.head
	fast := list.head

	for fast.next != nil {
		slow = slow.next
		fast = fast.next
		if fast.next != nil {
			fast = fast.next
		}
	}

	rightHead := slow.next
	slow.next = nil
	left = LinkedList{head: list.head, tail: slow}
	right = LinkedList{head: rightHead, tail: fast}
	return
}

func RemoveFromSlice(slice []int, index int) []int {
	left := slice[:index]
	right := slice[index+1:]
	return append(left, right...)
}
