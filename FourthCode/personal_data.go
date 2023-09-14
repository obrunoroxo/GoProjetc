package main

import (
	"fmt"
	"goprojects/FourthCode/methods"
	"goprojects/utils"
)

func main() {

	var name string
	var lastName string
	var age int
	var sex string

	fmt.Printf("Name: ")
	fmt.Scan(&name)

	fmt.Printf("Last Name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Age: ")
	fmt.Scan(&age)

	fmt.Printf("Sex: ")
	fmt.Scan(&sex)

	myDatas := methods.Person{
		Name:     name,
		LastName: lastName,
		Age:      age,
		Sex:      sex,
	}

	fmt.Printf("\n")
	methods.ShowMe(myDatas)
	utils.Sleep(1)
	utils.Clear()

}

// func Menu() []string {
// 	INIT := "===== WELCOME TO YOUR NAGE Map!! ====="
// 	INT := "\n=== How can i help you today?? ==="
// 	FIRST_OPT := "\n[ 1 ] - Add a new value in the map."
// 	SECOND_OPT := "\n[ 2 ] - Edit a value from the map."
// 	THIRD_OPT := "\n[ 3 ] - Delete a value from the map."
// 	FOURTH_OPT := "\n[ 4 ] - List all values from your map."
// 	FIFTH_OPT := "\n[ 5 ] - Exit."
// 	SIXTH_OPT := "\nOr you can just pass a name: "
// 	return []string{INIT, INT, FIRST_OPT, SECOND_OPT, THIRD_OPT, FOURTH_OPT, FIFTH_OPT, SIXTH_OPT}
// }
