package main

import "fmt"

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Prt()
	a.name = "123"
	a.Prt()
	// fmt.Println(a.name)

	b := B{}
	b.Prt()
	// fmt.Println(b.Name)

	// method value
	var c TZ
	c.Prt()

	// method exp
	(*TZ).Prt(&c)
}

type TZ int

type A struct {
	name string
}

func (x *TZ) Prt() {
	fmt.Println(x, "TZ")
}

func (x *A) Prt() {
	fmt.Println(x, x.name)
	x.name = "AA"
	fmt.Println(x, x.name)
}

func (x B) Prt() {
	x.Name = "BB"
	fmt.Println(x, x.Name)
}
