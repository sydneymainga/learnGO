package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")
	//no inheritance in golang :no super :no parent

	sydney := User{"sydney", "sydneymainga6@gmail.com", true, 26}
	fmt.Println(sydney)
	fmt.Printf("sydney details are : %+v\n", sydney)
	fmt.Printf("name is :% v and email is :%v", sydney.Name, sydney.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
