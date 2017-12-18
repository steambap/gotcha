package main

import (
	"fmt"
)

func main() {
	store := &Store{}
	store.State = &State{}
	fmt.Print(store.State)
	store.Dispatch(&Action{Name: "setName", Payload: "newName"})
	store.Dispatch(&Action{Name: "setId", Payload: 1010})
	fmt.Print(store.State)
}
