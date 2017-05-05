package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m1)
	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[v] = k
	}

	fmt.Println(m2)
	fmt.Println(ex1(2, 3, 4, 5))

	a, b := 1, 2
	ex1(a, b)
	fmt.Println(a, b)

	s1 := []int{7, 8, 9}
	ex2(s1)
	fmt.Println(s1)

	a = 10
	// ex3(&a)
	ex4(a)
	fmt.Println(a)

	z := func() {
		fmt.Println("Func A")
	}
	z()

	f := closure(10)
	fmt.Println(f(10), f(20))

	for i := 0; i < 3; i++ {
		// value copy
		defer fmt.Println(i)
		// reference copy
		defer func() {
			fmt.Println("closure", i)
		}()
	}
	defer fmt.Println("b")
	defer fmt.Println("c")
}

func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

func ex1(a ...int) (e, f, g rune) {
	a[0] = 3
	a[1] = 4
	fmt.Println(a)
	e, f, g = 'A', 'B', 'V'
	return e, f, g
}

func ex2(s []int) (e, f, g rune) {
	s[0] = 3
	s[1] = 4
	fmt.Println(s)
	e, f, g = 'A', 'B', 'V'
	return e, f, g
}

func ex3(s *int) {
	*s = 2
	fmt.Println(*s)
}

func ex4(s int) {
	s = 2
	fmt.Println(s)
}
