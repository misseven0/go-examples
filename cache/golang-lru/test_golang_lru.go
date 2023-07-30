package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
)

func Print[T any](t T) {
	fmt.Printf("printing type: %T\n", t)
}
func main() {
	l, _ := lru.New[int, interface{}](128)
	for i := 0; i < 256; i++ {
		l.Add(i, nil)
	}
	if l.Len() != 128 {
		panic(fmt.Sprintf("bad len: %v", l.Len()))
	}
}
