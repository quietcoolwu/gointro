package main

import "fmt"

func main() {
	// fmt.Println('a')
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

func B() {
	f := func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}
	defer f()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}
