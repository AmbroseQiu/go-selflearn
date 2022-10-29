package main

import (
	"fmt"

	subpack2 "example.com/golang_tutorial/pack1"
	subpack "example.com/golang_tutorial/pack2"
)

func init() {
	fmt.Println("init in main 1")
}

func init() {
	fmt.Println("init in main 2")
}

func init() {
	fmt.Println("init in main 3")
}

func main() {
	fmt.Println("hello program")
	subpack.CallG()
	subpack2.CallFunc()
	subpack.Print2DArray()
}
