package mypackage

import "fmt"

func Helper() {
	fmt.Println("***** Helper *****")

	person := Person{FirstName: "Amir", LastName: "Hasanzade", age: 35}
	fmt.Println(person.ToString())
}
