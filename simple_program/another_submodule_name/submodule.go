package subpackage

import "fmt"

func subpack2() {
	fmt.Println("call subpack2")
}

func CallFunc() {
	subpack2()
}
