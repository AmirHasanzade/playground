package mypackage

import (
	"fmt"
)

var PublicVar = "Amir Hasanzade"

type Person struct {
	FirstName string
	LastName  string
	age       int // private
}

func (p Person) ToString() string {
	return fmt.Sprintf("Person { FirstName: %q, LastName: %q, Age: %d }", p.FirstName, p.LastName, p.age)
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func Greet() {
	fmt.Println("Hello there")
}
