package main

import (
	"fmt"

	"amirhasanzade.com/playground/mypackage"
)

func main() {
	fmt.Println("PublicVar: " + mypackage.PublicVar)
	mypackage.Greet()

	var amir mypackage.Person = mypackage.Person{FirstName: "Amir 1", LastName: "Hasanzade"}
	amir.SetAge(35)
	fmt.Println(amir.ToString())

	amir2 := mypackage.Person{FirstName: "Amir 2", LastName: "Hasanzade"}
	amir2.SetAge(35)
	fmt.Println(amir2.ToString())

	amir3 := mypackage.Person{}
	amir3.FirstName = "Amir 3"
	amir3.LastName = "Hasanzade"
	amir3.SetAge(35)
	fmt.Println(amir3.ToString())

	mypackage.Helper()
}
