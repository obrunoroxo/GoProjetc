package main

import (
	"fmt"
	"goprojects/utils"
	"strconv"
)

func Values() map[string]int {
	// Create a map with string keys and int values
	values := make(map[string]int)

	// Add key-value pairs to the map
	values["Alice"] = 12
	values["Bob"] = 85
	values["Charlie"] = 24

	return values
}

func ShowMe(values map[string]int) {
	fmt.Printf("\n")
	for key, value := range values { // Show me all key, values
		fmt.Printf("%s's age: %d\n", key, value)
	}
}

func LineConnector(name string) {
	var answer string

	fmt.Println("\nIf you wish to go out, please press: \n[ E ] | [ Exit ] - to quit \n[ K ] | [ Keep ] - to continue")
	fmt.Printf("Answer: ")
	fmt.Scan(&answer)

	if answer == "Exit" || answer == "E" {
		utils.Close(true)
	} else if answer == "Keep" || answer == "K" {
		Separator(name)
	} else {
		fmt.Printf("\nOops! Invalid input passed: %s \nPlease, try again!", answer)
	}
}

func Separator(name string) {
	var choice string

	fmt.Println("\nChoose ONE to exchange: NAME or AGE? (Name - N / Age - A)")
	fmt.Printf("Answer: ")
	fmt.Scan(&choice)

	if choice == "Name" || choice == "N" {
		EditName(name)
	} else if choice == "Age" || choice == "A" {
		EditAge(name)
	} else {
		fmt.Printf("\nOops! Invalid input passed: %s\nPlease, try again!", choice)
	}
}

func AddNewValues(name string) { //  ->  ADD method
	var age int

	values := Values()

	fmt.Println("Please, choose an age to append to the map: ")
	fmt.Scan(&age)

	fmt.Printf("Adding [ %s: %d ] to the map...", name, age)
	values[name] = age
	utils.Sleep(1)
	utils.Clear()

	ShowMe(values)
}

func EditAge(name string) {
	values := Values()
	var age int

	fmt.Println("Choose an age to edit the map:", values)
	fmt.Scan(&age)

	if _, ok := values[name]; ok {
		// Create a new entry with the new key and the same value
		values[name] = age

		ShowMe(values)
	}
}

func EditName(name string) {
	values := Values()
	var answer string

	if value, ok := values[name]; ok {
		fmt.Printf("\nChoose a new name to edit the map [ %s: %d ]: ", name, value)
		fmt.Scan(&answer)
		// Create a new entry with the new key and the same value
		values[answer] = value

		// Delete the old key from the map
		delete(values, name)

		ShowMe(values)

	} else {
		fmt.Println("Something didn't work correctly, please try again...")
	}
}

func ConvertValuesToINT(value string) interface{} {
	convert, err := strconv.Atoi(value)
	if err != nil {
		return value
	}

	fmt.Println("The int value is:", convert)

	return convert
}

func VerifyValues(name string) {
	var answer string
	values := Values()
	convert := ConvertValuesToINT(name)

	if convert == 1 {
		fmt.Println("Chose a new name to append in your map:")
		fmt.Scan(&answer)
		AddNewValues(answer)
		utils.Close(true)
	}

	if value, ok := values[name]; ok {
		fmt.Printf("Oops! The name passed already exists [ %s: %d ]...", name, value)

		for {
			fmt.Println("\nDo you wish to exchange your name or age? (Yes - Y / No - N)")
			fmt.Printf("Answer: ")
			fmt.Scan(&answer)

			if answer == "Yes" || answer == "Y" {
				LineConnector(name)
			} else if answer == "No" || answer == "N" {
				utils.Clear()
				main()
			} else {
				fmt.Printf("\nOops! Invalid input passed: %s\nPlease, try again!", answer)
			}

		}

	} else {
		fmt.Printf("Excellent! Last name [ %s ] does not yet exist on the map...\n", name)
		AddNewValues(name)
	}
}

func main() {
	var answer string

	menu_list := Menu()
	for _, value := range menu_list {
		fmt.Println(value)
	}

	fmt.Scan(&answer)
	fmt.Printf("\nYou typed: %s\n", answer)
	VerifyValues(answer)
}

func Menu() []string {
	INIT := "===== WELCOME TO YOUR NAGE Map!! ====="
	INT := "\n=== How can i help you today?? ==="
	FIRST_OPT := "\n[ 1 ] - Add a new value in the map."
	SECOND_OPT := "\n[ 2 ] - Edit a value from the map."
	THIRD_OPT := "\n[ 3 ] - Delete a value from the map."
	FOURTH_OPT := "\n[ 4 ] - List all values from your map."
	FIFTH_OPT := "\n[ 5 ] - Exit."
	SIXTH_OPT := "\nOr you can just pass a name: "

	return []string{INIT, INT, FIRST_OPT, SECOND_OPT, THIRD_OPT, FOURTH_OPT, FIFTH_OPT, SIXTH_OPT}
}
