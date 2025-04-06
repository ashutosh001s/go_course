package basics

import "fmt"

//keep package names short , lowercase
type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//PascalCase used for types : structs , interfaces , enums
	//Eg : CalculateArea , UserInfo, NewHTTPRequest

	//snake_case used for : variables , file names
	//Eg: user_id , first_name, http_request

	//UPPERCASE used for : constants
	//Eg : PI, LIMIT

	//mixedCase or camelCase used for variables
	//Eg : javaScript , htmlDocument , isValid

	const MAX_RETRIES = 5

	var empoyeeId = 1001
	fmt.Println(empoyeeId)
}
