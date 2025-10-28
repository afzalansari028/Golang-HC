package main

import (
	"fmt"
	"reflect"
)

const LoginToken string = "asdfghijk" //Public

func main() {
	// fmt.Println("Variables")
	var username string = "Afzal"
	fmt.Println(username)
	fmt.Printf("Variables of type : %T \n", username)
	// a := fmt.Sprintf("Variables of type : %T \n", username)  type checking
	// fmt.Println(a)
	b := reflect.TypeOf(username)
	fmt.Println("type:", b)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables of type : %T \n", isLoggedIn)

	var smallVal int64 = 1234678897545678289
	fmt.Println(smallVal)
	fmt.Printf("Variables of type : %T \n", smallVal)

	var smallFloat float64 = 123.46788975456782898976636372764646367282827462624274824728748278472343747242274
	fmt.Println(smallFloat)
	fmt.Printf("Variables of type : %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	var anotherVariableString string
	fmt.Println(anotherVariableString)
	fmt.Println("String by default takes :", anotherVariableString)

	// implicit type

	var website = "learnGolang.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	//no var style

	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
