package main

import "fmt"

type empry interface {
}

type Connector interface {
	Connect()
}

type USB interface {
	Connector
	Name() string
}

type PhoneConnector struct {
	name string
}

func (phone PhoneConnector) Name() string {
	return phone.name
}

func (phone PhoneConnector) Connect() {
	fmt.Println("Connected: ", phone.name)
}

func Disconnect(dock interface{}) {
	// if phone, ok := dock.(PhoneConnector); ok {
	// 	fmt.Println("Disconnected: ", phone.name)
	// 	return
	// }

	switch device := dock.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected: ", device.name)
	default:
		fmt.Println("Unknown Device")
	}

}

type TVConnector struct {
	name string
	size int
}

func (tv TVConnector) Connect() {
	fmt.Println("Connected: ", tv.name)
}

func TestPhone() {
	// var dock1 USB
	dock1 := PhoneConnector{
		name: "Phone1",
	}
	fmt.Println(dock1.Name())
	dock1.Connect()
	Disconnect(dock1)

	dock2 := Connector(dock1)
	dock2.Connect()

	// tv := TVConnector{
	// 	name: "TV1",
	// 	size: 43,
	// }
	//
	// dock3 := USB(tv)
	// dock3.Connect()

}

func main() {
	TestPhone()
}
