package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

type Human struct {
	Name string
	Age  int
	Sex  int
}

type Teacher struct {
	Human
	StaffID int
}

type Student struct {
	Human
	StuID int
}

type AnonymousPerson struct {
	string
	int
}

func main() {
	Alice := &Person{
		Name: "Alice",
	}
	// Alice.Name = "Alice"
	Alice.Age = 19
	fmt.Println(Alice)
	// ptr -> ref copy
	changeAlice(Alice)
	fmt.Println(Alice)

	Bob := &struct {
		Name string
		Age  int
	}{
		Name: "Bob",
		Age:  20,
	}
	fmt.Println(Bob)

	Cathy := &Person{
		Name: "Cathy",
		Age:  25,
	}

	Cathy.Contact.Phone = "18212345678"
	Cathy.Contact.City = "Shanghai"
	fmt.Println(Cathy)

	// var Cathy_2 Person
	Cathy2 := Cathy
	fmt.Println(Cathy2 == Cathy)

	Aman := AnonymousPerson{"Aman", 999}
	Aman1 := AnonymousPerson{"Aman", 99}
	fmt.Println(Aman == Aman1)

	t1 := &Teacher{
		Human: Human{
			Name: "Fuck",
			Sex:  1,
			Age:  30,
		},
		StaffID: 10086,
	}

	t2 := &Student{
		Human: Human{
			Name: "Jack",
			Sex:  0,
			Age:  18,
		},
		StuID: 100,
	}

	t2.Human.Age = 19
	t2.Age = 18

	fmt.Println(t1, t2)

}

func changeAlice(p *Person) {
	p.Age = 20
	p.Name = "Alice2"
	fmt.Println("Alice changed", p)
}
