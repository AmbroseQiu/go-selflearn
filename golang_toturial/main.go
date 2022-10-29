package main

import (
	subpack "example.com/golang_tutorial/pack1"
	subpack2 "example.com/golang_tutorial/pack2"
)

// func init() {
// 	fmt.Println("init in main 1")
// }

// func init() {
// 	fmt.Println("init in main 2")
// }

// func init() {
// 	fmt.Println("init in main 3")
// }

func main() {
	subpack2.CallG()
	subpack.CallFunc()
	// subpack.Print2DArray()
}
