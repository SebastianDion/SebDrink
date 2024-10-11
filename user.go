package main

import (
	"math/rand"
	"time"
	"fmt"
)

func generateRandomID() string {
	rand.Seed(time.Now().UnixNano())  
    return fmt.Sprintf("%05d", rand.Intn(100000)) // Generates a random number between 0 and 99,999
}

type User struct{
	name string
	id string
	userDrink drink
}

func adminOption(){
	var option int
	var Password string
	fmt.Print("Enter Password:")
	fmt.Scanf("%s", &Password)
	if Password == "admin123"{
		fmt.Println("Welcome to Admin Page!")
		fmt.Println("======================")
		fmt.Println("1. Add Stock?")
		fmt.Println("or")
		fmt.Println("2. Display Stock?")
		fmt.Print(">")
		fmt.Scanf("%d", &option)
		switch option{
		case 1:{addStock()
			return}
		case 2:{displayStock()
			return}
		default:{
			fmt.Println("Invalid Choice!!")
		}
		}
		
		
	}else{
	 fmt.Println("Invalid Password! Bye :p")
	}
}

func userInfo() User{
	var user User
	fmt.Print("Enter your name: ")
	fmt.Scan(&user.name)
	user.id = "SEB-" + generateRandomID()
	return user
}
