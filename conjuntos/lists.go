package main

import (
	"container/list"
	"fmt"
)

func main() {
	numeros := list.New()
	el := numeros.PushBack(1)
	numeros.PushFront(0)
	numeros.PushBack(2)

	for e := numeros.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("___________________________________")

	numeros.Remove(el)
	for e := numeros.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
