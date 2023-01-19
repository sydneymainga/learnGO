package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("welcome to switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you can now open : value of dice is 1")

	case 2:
		fmt.Println("you can now move to spot 2 : value of dice is 2")

	case 3:
		fmt.Println("you can now move to spot 3 : value of dice is 3")
	case 4:
		fmt.Println("you can now move to spot 4 : value of dice is 4")
	case 5:
		fmt.Println("you can now move to spot 5 : value of dice is 5")
	case 6:
		fmt.Println("you can role dice now : value of dice is 6")

	}
}
