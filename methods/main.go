package main

import "fmt"

type TZ int

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Prt()
	// fmt.Println(a.Name)

	b := B{}
	b.Prt()
	// fmt.Println(b.Name)

	// c := TZ{}
	// c.Prt()
	var c TZ
	c.Prt()
}

func (x *TZ) Prt() {
	fmt.Println(x, "TZ")
}

func (x *A) Prt() {
	x.Name = "AA"
	fmt.Println(x, x.Name)
}

func (x B) Prt() {
	x.Name = "BB"
	fmt.Println(x, x.Name)
}
