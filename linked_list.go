package main

import (
	"strings"
)

type BiDLLNode struct {
	data     string
	previous *BiDLLNode
	next     *BiDLLNode
}

type BiDirectionalLinkedList struct {
	head *BiDLLNode
	tail *BiDLLNode
}

func (list *BiDirectionalLinkedList) PushFront(data string) {
	newNode := &BiDLLNode{data: data, next: list.head}
	if list.head != nil {
		list.head.previous = newNode
	}
	list.head = newNode
	if list.tail == nil {
		list.tail = newNode
	}
}

func (list *BiDirectionalLinkedList) PushBack(data string) {
	newNode := &BiDLLNode{data: data, previous: list.tail}
	if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode
	if list.head == nil {
		list.head = newNode
	}
}

func (list *BiDirectionalLinkedList) PopFront() (string, bool) {
	if list.head == nil {
		return "", false
	}
	data := list.head.data
	list.head = list.head.next
	if list.head != nil {
		list.head.previous = nil
	} else {
		list.tail = nil
	}
	return data, true
}

func (list *BiDirectionalLinkedList) PopBack() (string, bool) {
	if list.tail == nil {
		return "", false
	}
	data := list.tail.data
	list.tail = list.tail.previous
	if list.tail != nil {
		list.tail.next = nil
	} else {
		list.head = nil
	}
	return data, true
}

func (list *BiDirectionalLinkedList) getListFront() string {
	if list.head == nil {
		return "[]"
	}

	listElements := make([]string, 0)

	current := list.head
	for current != nil {
		listElements = append(listElements, current.data)
		current = current.next
	}
	listElements = append(listElements, "nil")

	return "[" + strings.Join(listElements, ", ") + "]"
}

func (list *BiDirectionalLinkedList) getListBack() string {
	if list.head == nil {
		return "[]"
	}

	listElements := make([]string, 0)

	current := list.tail
	for current != nil {
		//fmt.Printf("%v -> ", current.data)
		listElements = append(listElements, current.data)
		current = current.previous
	}
	listElements = append(listElements, "nil")

	return "[" + strings.Join(listElements, ", ") + "]"
}
