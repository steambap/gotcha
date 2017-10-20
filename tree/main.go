package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func NewTree(root int, values... int) *Tree {
	tree := &Tree{nil, root, nil}

	for _, value := range values {
		Insert(tree, value)
	}

	return tree
}

func Insert(t *Tree, item int) *Tree {
	if t == nil {
		return &Tree{nil, item, nil}
	} else {
		v := t.Value

		if item > v {
			return &Tree{t.Left, v, Insert(t.Right, item)}
		} else if item < v {
			return &Tree{Insert(t.Left, item), v, t.Right}
		} else {
			return t
		}
	}
}

func (t *Tree) print() {
	if t.Left != nil {
		t.Left.print()
	}
	fmt.Print(t.Value, " ")
	if t.Right != nil {
		t.Right.print()
	}
}

func main() {
	rand.Seed(0)
	tree := &Tree{nil, 10, nil}
	for i:=0; i< 30; i++ {
		tree = Insert(tree, rand.Intn(30))
	}

	tree.print()
}
