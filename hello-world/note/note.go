package note

import "fmt"

func SayHelloWorld() {
	fmt.Println("hello world")
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// func getTwoValue(n1, n2 int) (int, int) {
// 	res := n1 + n2
// 	diff := n1 - n2
// 	return res, diff
// }

func getTwoValuePretty(n1, n2 int) (res, diff int) {
	res = n1 + n2
	diff = n1 - n2
	return
}

func CallFunction() {
	res := getSum(2, 3)
	fmt.Println("res=", res)
	res, diff := getTwoValuePretty(4, 3)
	fmt.Println("res=", res, "diff=", diff)

	res, diff = func(n1, n2 int) (res, diff int) {
		res = n1 + n2
		diff = n1 - n2
		return
	}(4, 3)
	fmt.Println("res=", res, "diff=", diff)
	// fmt.Printf("getTwoValuePretty function's value=%v, and type=%T", getTwoValuePretty, getTwoValuePretty)
}
