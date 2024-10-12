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

func payment(){
	var choice int
	var ok string
	fmt.Println("Choose Payment:")
	fmt.Println("1. Cash")
	fmt.Println("2. Cashless")
	fmt.Print(">")
	fmt.Scanf("%d", &choice)
	switch choice{
	case 1:{
		fmt.Println("Payed")
		return
	}
case 2:{
	fmt.Println("QR Code")
	fmt.Println("type 'ok' when done")
	fmt.Scanf("%s", ok)
	fmt.Println("Done")
	return
}
default:{
	fmt.Println("Payment Failed")
}
	}
}
