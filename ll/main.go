package main

import (
	"fmt"
)

type LinkList struct {
	Head int
	Tail *LinkList
}

func NewLinkList(head int, tails... int) *LinkList {
	linkList := &LinkList{head, nil}

	if len(tails) > 0 {
		linkList.Tail = NewLinkList(tails[0], tails[1:]...)
	}

	return linkList
}

func (l *LinkList) Add (item int) {
	if l.Tail == nil {
		l.Tail = NewLinkList(item)
	} else {
		l.Tail.Add(item)
	}
}

func (l *LinkList) Has (item int) bool {
	if l.Head == item {
		return true
	}
	if l.Tail == nil {
		return false
	}

	return l.Tail.Has(item)
}

func (l *LinkList) print () {
	fmt.Printf("%d -> ", l.Head)
	if l.Tail != nil {
		l.Tail.print()
	}
}

func main() {
	var root = NewLinkList(1)
	root.Add(5)
	root.Add(10)
	root.Add(0)
	root.Add(100)

	root.print()
}