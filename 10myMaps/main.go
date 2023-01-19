package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to my maps ")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "rubyonrails"
	languages["PY"] = "python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS is abbriviation to : ", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages after deletion : ", languages)

	//loops are intresting in golang

	for key, value := range languages {
		fmt.Printf("for key : %v value is : %v \n ", key, value)
	}

}
