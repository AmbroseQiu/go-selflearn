package main

import (
	"fmt"
	subpack2 "package_module_name/another_submodule_name"
	subpack "package_module_name/submodule_name"
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
