package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays")
	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "mangoe"
	fruitlist[2] = "avocado"
	fruitlist[3] = "pineapple"

	fmt.Println("list of fruits :", fruitlist)
	fmt.Println("list of fruits :", len(fruitlist))

	var veglist = [3]string{"kales", "cabbage", "saga"}

	fmt.Println("list of veggy :", veglist)
	fmt.Println("list of veggy :", len(veglist))

}
