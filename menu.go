package main

import (
	"fmt"
)

type drink struct {
	drinkType   string
	SugarLevel  string
	Temperature string
	
}

func order() {

	var choice int
	
	for{
	fmt.Println("Welcome to SebDrink")
	fmt.Println("=====================")
	fmt.Println("\tMenu")
	fmt.Println("=====================")
	menu()
	fmt.Println("Pick a Number")
	fmt.Print(">")
	fmt.Scanf("%d", &choice)
	if choice== 4 {
		break 
	}
	drinkChoice(choice)
	}
	fmt.Println("Thank You for Ordering SebDrink! ^-^ ")

}

func menu() {
	fmt.Println("1. Tea    \t3$")
	fmt.Println("2. Coffee \t5$")
	fmt.Println("3. Matcha \t10$")
	fmt.Println("4. EXIT")
}



func drinkChoice(number int) {
	var drinkChoice drink

	switch number {
	case 1:
		if inventory.tea > 0 {
			drinkChoice.drinkType = "Tea"
			inventory.tea-- // Decrease tea stock
		} else {
			fmt.Println("Sorry, we're out of Tea!")
			return
		}
	case 2:
		if inventory.coffee > 0 {
			drinkChoice.drinkType = "Coffee"
			inventory.coffee-- // Decrease coffee stock
		} else {
			fmt.Println("Sorry, we're out of Coffee!")
			return
		}
	case 3:
		if inventory.matcha > 0 {
			drinkChoice.drinkType = "Matcha"
			inventory.matcha-- // Decrease matcha stock
		} else {
			fmt.Println("Sorry, we're out of Matcha!")
			return
		}
	case 19122002:
		adminOption()
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
		return
	}
	
	// Now get the configuration for the drink
	configuredDrink := config() // Get the configured drink
	drinkChoice.SugarLevel = configuredDrink.SugarLevel // Set sugar level
	drinkChoice.Temperature = configuredDrink.Temperature // Set temperature

	if configuredDrink.SugarLevel != "Less Sugar" {
		inventory.sugar-- 
	}
	
	user := userInfo() // Get user info
	user.userDrink = drinkChoice // Set user's drink
	receipt(user) // Pass user to receipt
}


func config() drink {
	var choiceSugar int
	var choiceTemp int
	fmt.Println("1. Less sugar")
	fmt.Println("2. Normal")
	fmt.Println("3. Extra Sugar")
	fmt.Println(">")
	fmt.Scanf("%d", &choiceSugar)
	fmt.Println("1. Hot")
	fmt.Println("2. Cold")
	fmt.Print("> ")
	fmt.Scanf("%d", &choiceTemp)

	var config drink

	
	switch choiceSugar {
	case 1:
		
		break
	case 2:
		if inventory.sugar > 0 {
			inventory.sugar-- 
		} else {
			fmt.Println("Sorry, we're out of Sugar!")
		}
	case 3:
		if inventory.sugar > 1 {
			inventory.sugar -= 2  
		} else {
			fmt.Println("Sorry, we're out of Sugar!")
		}
	default:
		fmt.Println("Invalid choice. Please try again.")
	}

	switch choiceTemp {
	case 1:
		config.Temperature = "Hot"
	case 2:
		config.Temperature = "Cold"
	default:
		fmt.Println("Invalid choice. Please try again.")
	}

	return config

}

func receipt(user User) {
	var ok string
    fmt.Println("User Name:", user.name)
    fmt.Println("User ID:", user.id)
    fmt.Println("Drink Type:", user.userDrink.drinkType)
    fmt.Println("Sugar Level:", user.userDrink.SugarLevel)
    fmt.Println("Temperature:", user.userDrink.Temperature)
	fmt.Println("type ok if you're done")
	fmt.Scanf("%s",&ok)
}

