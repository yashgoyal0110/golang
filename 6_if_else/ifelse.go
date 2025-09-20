package main

import "fmt"

func main() {
	age := 16

	if age >= 18 {
		fmt.Println("person is an adult")
	} else {
		fmt.Println("person is not an adult")
	}

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is a kid")
	}

	var role = "admin"
	var hasPermissions = false

	if role == "admin" && hasPermissions {
		fmt.Println("yes")
	}

	if age := 20; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	// go does not have ternary, you will have to use normal if else

}
