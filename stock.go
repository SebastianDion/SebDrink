package main

import "fmt"

type stock struct{
	sugar int
	ice int
	tea int
	coffee int
	matcha int
}

var inventory = stock{
	sugar:  100, 
	ice:    100, 
	tea:    50, 
	coffee: 50,  
	matcha: 50,  
}

func displayStock() {
	var ok string
	fmt.Println("Current Stock Levels:")
	fmt.Println("Sugar:", inventory.sugar)
	fmt.Println("Ice:", inventory.ice)
	fmt.Println("Tea:", inventory.tea)
	fmt.Println("Coffee:", inventory.coffee)
	fmt.Println("Matcha:", inventory.matcha)
	fmt.Println("type ok if you're done")
	fmt.Scanf("%s",&ok)
	
}

func addStock(){
	var option int
	fmt.Println("Which material you want to add?")
	fmt.Println("1. Sugar")
	fmt.Println("2. Ice")
	fmt.Println("3. Tea")
	fmt.Println("4. Coffee")
	fmt.Println("5. Matcha")
	fmt.Print(">")
	fmt.Scanf("%d", &option)
	
	switch option{
	case 1:{
		inventory.sugar++
	}
	case 2:{
		inventory.ice++
	}
	case 3:{
		inventory.tea++
	}
	case 4:{
		inventory.coffee++
	}
	case 5:{
		inventory.matcha++
	}
	}
}

