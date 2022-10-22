package package_module_name

import "fmt"

var g = func() int {
	fmt.Println("init static var")
	return 0
}

func init() {
	fmt.Println("init in sub 1")
}

func init() {
	fmt.Println("init in sub 2")
}

func init() {
	fmt.Println("init in sub 3")
}

func CallG() {
	g()
}

func Print2DArray() {
	two_d_arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	for i, arr := range two_d_arr {
		for j, val := range arr {
			fmt.Printf("a[%v][%v] = %v\t", i, j, val)
		}
		fmt.Println()
	}

	s1 := two_d_arr[:3]

	for i, arr := range s1 {
		for j, val := range arr {
			fmt.Printf("s1[%v][%v] = %v\t", i, j, val)
		}
		fmt.Println()
	}

	s2 := append(s1, []int{2, 3}, []int{4, 5}, []int{6, 7})

	for i, arr := range s2 {
		for j, val := range arr {
			fmt.Printf("s2[%v][%v] = %v\t", i, j, val)
		}
		fmt.Println()
	}

}
