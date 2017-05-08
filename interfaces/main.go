package main

import "fmt"

type empty interface {
}

type Connector interface {
	Connect()
}

type USBConnector interface {
	Connector
	Name() string
}

type Phone struct {
	name   string
	screen float64
}

type TV struct {
	name string
	size float64
}

func (p Phone) Name() string {
	return p.name
}

func (p Phone) Connect() {
	fmt.Println("Phone Connected: ", p.name, p.screen)
}

func (tv TV) Connect() {
	fmt.Println("TV Connected: ", tv.name, tv.size)
}

func Disconnect(dock interface{}) {

	switch device := dock.(type) {
	case Phone:
		fmt.Println("Disconnected Phone: ", device.name, device.screen)
	case TV:
		fmt.Println("Disconnected TV: ", device.name, device.size)
	default:
		fmt.Println("Unknown Device")
	}
}

func TestPhone() {
	// var dock1 USB
	phone1 := Phone{"Phone1", 5.5}

	// Issues on copy
	dock2 := USBConnector(phone1)
	// dock2 = Connector(phone1)
	fmt.Println(dock2.Name())
	phone1.name = "Phone2"
	dock2.Connect()
	Disconnect(dock2)

	dock3 := Connector(phone1)
	dock3 = USBConnector(phone1)
	// Why can't exec?
	// fmt.Println(dock3.Name())
	dock3.Connect()
	Disconnect(dock3)

}

func TestNilInterface() {
	var a interface{}
	fmt.Println(a == nil)

	var p *int = nil
	a = p
	fmt.Println(a == nil)
}

func main() {
	TestPhone()
	TestNilInterface()
}
