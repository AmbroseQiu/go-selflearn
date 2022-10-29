package pack2

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
	key := selectByKey("please", "use", "this", "func")
	fmt.Println("key = ", key)
	tryMap()
}

func selectByKey(text ...string) (key int) {
	for i, v := range text {
		fmt.Printf("%v: %s\n", i+1, v)
	}
	fmt.Print("type key: ")
	// fmt.Scanln(&key)
	return key
}

func tryMap() {
	var m1 map[string]string
	fmt.Println("m1 == nil ? ", m1 == nil)

	m1 = make(map[string]string, 2)

	m1["how r u"] = "I fine thank u"

	m1["a for apple"] = "b for ball"

	for key, value := range m1 {
		fmt.Printf("key=%v,value=%v\n", key, value)
	}

	value, ret := m1["who are u"]
	if ret {
		fmt.Println(value)
	} else {
		fmt.Println("Ambrose")
	}

	delete(m1, "a for apple")
	for key, value := range m1 {
		fmt.Printf("key=%v,value=%v\n", key, value)
	}
}
