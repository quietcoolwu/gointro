package main

import "fmt"

type xs int

func main() {
	testIncrease()

	testslice()

	ans := testmap()

	fmt.Println(ans, len(ans))

}

func testIncrease() {
	var a xs = 3
	fmt.Println(a)
	a.Increase(20)
	fmt.Println(a)
}

func testslice() {
	// b := [10]int{}
	b := [...][3]int{{2, 3, 4}, {5, 6, 7}}
	fmt.Println(b, len(b))

	ar := [...]string{"a", "b", "c", "d", "e"}
	sub1 := ar[2:]
	sub2 := ar[:4]
	fmt.Println(sub1, sub2)
	sub1[1] = "z"
	fmt.Println(ar, sub1, sub2)
	fmt.Println(cap(sub1), cap(sub2))
	sub1 = append(sub1, "f")
	sub1 = append(sub1, "g")
	sub2 = append(sub2, "h")
	sub2 = append(sub2, "i")
	sub2 = append(sub2, "j")
	sub2 = append(sub2, "k")

	fmt.Println(ar, sub1, sub2)
	fmt.Println(cap(sub1), cap(sub2))
}

func testmap() map[string]int {
	numbers := make(map[string]int)
	numbers["a"] = 1
	numbers["b"] = 2
	numbers["b"] = 3
	x, ok := numbers["b"]
	fmt.Println(x, ok)
	delete(numbers, "b")

	return numbers
}

func (x *xs) Increase(inc int) {
	// force type conversion
	// *x += inc
	*x += xs(inc)
}
