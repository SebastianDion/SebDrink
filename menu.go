package main

import (
	"fmt"
)

type drink struct {
	drinkType   string
	SugarLevel  string
	Temperature string
	
	price int
	
}


func order() {
	var choice int
	var totalPrice int

	for {
		
		fmt.Println("Welcome to SebDrink")
		fmt.Print("Your Total Bill: ")
		fmt.Print(totalPrice)
		fmt.Println("$")
		fmt.Println("=====================")
		fmt.Println("\tMenu")
		fmt.Println("=====================")
		menu()
		fmt.Println("Pick a Number")
		fmt.Print(">")
		fmt.Scanf("%d", &choice)
		if choice == 4 {
			break
		}

		// Get the drink from drinkChoice and update totalPrice
		drink := drinkChoice(choice)
		if drink.drinkType != "" {
			totalPrice += drink.price
		}
	}
	fmt.Printf("Final Price: %d$\n", totalPrice) // Print the final total price
	payment()
	

	fmt.Println("Thank You for Ordering SebDrink! ^-^ ")

}


func menu() {
	fmt.Println("1. Tea    \t3$")
	fmt.Println("2. Coffee \t5$")
	fmt.Println("3. Matcha \t10$")
	fmt.Println("4. EXIT")
}



func drinkChoice(number int) drink {
	var drinkChoice drink

	switch number {
	case 1:
		if inventory.tea > 0 {
			drinkChoice.drinkType = "Tea"
			drinkChoice.price = 3
			inventory.tea-- // Decrease tea stock
		} else {
			fmt.Println("Sorry, we're out of Tea!")
			return drink{} // Return an empty drink
		}
	case 2:
		if inventory.coffee > 0 {
			drinkChoice.drinkType = "Coffee"
			drinkChoice.price = 5
			inventory.coffee-- // Decrease coffee stock
		} else {
			fmt.Println("Sorry, we're out of Coffee!")
			return drink{} // Return an empty drink
		}
	case 3:
		if inventory.matcha > 0 {
			drinkChoice.drinkType = "Matcha"
			drinkChoice.price = 10
			inventory.matcha-- // Decrease matcha stock
		} else {
			fmt.Println("Sorry, we're out of Matcha!")
			return drink{} // Return an empty drink
		}
	case 19122002:
		adminOption()
		return drink{} // Return an empty drink
	default:
		fmt.Println("Invalid choice. Please try again.")
		return drink{} // Return an empty drink
	}

	// Now get the configuration for the drink
	configuredDrink := config() // Get the configured drink
	drinkChoice.SugarLevel = configuredDrink.SugarLevel // Set sugar level
	drinkChoice.Temperature = configuredDrink.Temperature // Set temperature

	if configuredDrink.SugarLevel != "Less Sugar" {
		inventory.sugar--
	}
	user := userInfo() // Get user info2
	
	user.userDrink = drinkChoice // Set user's drink
	receipt(user) 
	return drinkChoice // Return the configured drink
}



func config() drink {
	var choiceSugar int
	var choiceTemp int
	fmt.Println("1. Less sugar")
	fmt.Println("2. Normal")
	fmt.Println("3. Extra Sugar")
	fmt.Print(">")
	fmt.Scanf("%d", &choiceSugar)
	fmt.Println("1. Hot")
	fmt.Println("2. Cold")
	fmt.Print("> ")
	fmt.Scanf("%d", &choiceTemp)

	var config drink

	
	switch choiceSugar {
	case 1:
	config.SugarLevel = "Less Sugar"
case 2:
	config.SugarLevel = "Normal"
	if inventory.sugar > 0 {
		inventory.sugar--
	} else {
		fmt.Println("Sorry, we're out of Sugar!")
	}
case 3:
	config.SugarLevel = "Extra Sugar"
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

