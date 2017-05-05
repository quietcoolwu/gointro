package main

import "fmt"

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		fmt.Println(i)
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer closure i = ", i) }()
		fs[i] = func() { fmt.Println("clousure i = ", i) }
	}

	for _, f := range fs {
		f()
	}
}
