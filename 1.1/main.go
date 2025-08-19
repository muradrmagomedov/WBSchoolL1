package main

import "fmt"

type Human struct {
	FirstName string
	LastName  string
}

func (h Human) Hello() string {
	return "Hello world"
}

func (h Human) Greeting(name string) {
	fmt.Printf("Hello %s, my name is %s", name, h.FirstName)
}

type Action struct {
	Human
}

func main() {
	action := Action{}
	action.FirstName = "John"
	action.LastName = "Dow"
	action.Greeting("Peter")
}
