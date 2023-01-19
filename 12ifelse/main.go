package main

import "fmt"

func main() {

	fmt.Println("welcome to controll flow statements")

	loginCount := 23
	result := ""

	if loginCount < 10 {
		result = "regular user"
	} else {
		result = "frequent user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

}
