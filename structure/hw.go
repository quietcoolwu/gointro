package main

import "fmt"

type A struct {
	B
	// C
	Name string
}

type B struct {
	C
	Name string
}

type C struct {
	Name string
}

func main() {
	a := A{
		Name: "A",
		B: B{
			Name: "B",
			C: C{
				Name: "C",
			},
		},
	}
	// ambiguous name
	fmt.Println(a.Name, a.B.Name, a.B.C.Name)
}
