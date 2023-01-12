package main

import (
	"fmt"
)
const LoginToken string ="fsasdfghjjhgfd" //Public since first letter("L") is caps

func main(){

	
	var username string ="mainga"
    fmt.Println(username)
	fmt.Printf("variable "+username+" is of type: %T \n",username)

	var isLoggedin bool =true
    fmt.Println(isLoggedin)
	fmt.Printf("variable is of type: %T \n",isLoggedin)
	
	var smallValue uint8 =255
    fmt.Println(smallValue)
	fmt.Printf("variable is of type: %T \n",smallValue)

	var smallFloat float32 =255.67765456787654334567
    fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n",smallFloat)
     //default values 
	var anotherValue int
    fmt.Println(anotherValue)
	fmt.Printf("variable is of type: %T \n",anotherValue)

	//implicit type
	var website = "wayamoney.com"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n",website)

	//no var style

	numberOfUsers := 300000.000
	fmt.Println(numberOfUsers)
	fmt.Printf("variable is of type: %T \n",numberOfUsers)


    username2 :="mainga"
    fmt.Println(username2)
	fmt.Printf("variable "+username2+" is of type: %T \n",username2)

	fmt.Println(LoginToken)
	fmt.Printf("variable "+LoginToken+" is of type: %T \n",LoginToken)
}