package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello(name string) {
	fmt.Println("Hello,", name, ". My name is", u.Name)
}

func main() {
	u := User{1, "Tom", 12}
	info(u)

	m := Manager{User{1, "Tom", 12}, "Manager1"}
	indexinfo(m)

	x := 123
	fmt.Printf("%#v\n", x)
	simpleReflectWrite(x)

	u = User{2, "Jack", 20}
	setterReflections(&u)
	fmt.Println(u)

	u.Hello("Wukai")
	dynamicReflections(u)
}

func dynamicReflections(o interface{}) {
	v := reflect.ValueOf(o)
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("Dick")}
	mv.Call(args)
}

func setterReflections(o interface{}) {
	v := reflect.ValueOf(o)

	if (v.Kind() != reflect.Ptr) || (!v.Elem().CanSet()) {
		fmt.Println("Can't Set! Back to Orignal", o)
		return
	}
	v = v.Elem()

	// Fail Silently
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Bad Field Name!")
		return
	}
	if f.Kind() == reflect.String {
		fmt.Println("Good Field Name:", f)
		f.SetString("What?")
	}
}

func indexinfo(m interface{}) {
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.Field(1))
	// 传入slice获取字段
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

}

func simpleReflectWrite(x int) {
	v := reflect.ValueOf(&x)
	fmt.Printf("%#v\n", v)
	v.Elem().SetInt(999)
	fmt.Println(x)
}

func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())
	if t.Kind() != reflect.Struct {
		fmt.Println("Error! Parameter Not a struct!")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
