package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slice ")
	var fruitlist = []string{"apple", "guava", "tomatoe"}
	fmt.Printf("type of fruitlist is : %T \n", fruitlist)

	fruitlist = append(fruitlist, "pineapple", "banana")

	fmt.Println("new list after appending : ", fruitlist)

	fruitlist = append(fruitlist[1:]) //fruitlist[1:3] //fruitlist[:3]
	fmt.Println(fruitlist)

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 334
	highscores[2] = 434
	highscores[3] = 534

	highscores = append(highscores, 555, 666, 777)

	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	//remove values from slice based on index

	var courses = []string{"java", "golang", "python", "c#", "c++"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
