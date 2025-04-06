package basics

import "fmt"

var midddleName string = "Came"

func main() {
	//var age int32
	var name string = "John"
	var name2 = "Tom"

	//count := 24
	lastName := "Smith"

	// Default values
	//Numeric Types: 0
	//Boolean Types : false
	//String Type: ""
	//Pointer ,slices , maps ,functiosn, structs : nil

	// <=== SCOPE ====>
	n := getName(name, lastName)
	n2 := getName(name2, lastName)

	fmt.Println(n)
	fmt.Println(n2)

}

func getName(firstName string, lastName string) string {
	firstName = "" //it should not be possible
	return firstName + " " + lastName
}
