package main

import (
	"fmt"

	"goprojects/utils"
)

func sum_operation(a, b int) int {
	return a + b
}

func subtract_operation(a, b int) int {
	return a - b
}

func multiply_operation(a, b int) int {
	return a * b
}

func divide_operation(a, b int) int {
	return a / b
}

func main() {

	sum := "[1] - Sum"
	subtract := "[2] - Subtract"
	multiply := "[3] - Multiply"
	divide := "[4] - Divide"
	quit := "[5] - Exit"
	a_value := "Choose your first value:"
	b_value := "Choose your second value:"

	fmt.Println("Hello, World!")

	fmt.Println("\nWelcome to your calculator!")

	for {
		var first_input string

		fmt.Println("\nChoose your operation: ")
		fmt.Println(sum)
		fmt.Println(subtract)
		fmt.Println(multiply)
		fmt.Println(divide)
		fmt.Println(quit)

		fmt.Print("\nAnswer: ")
		_, err := fmt.Scan(&first_input)

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		if first_input == "5" {
			fmt.Println("Exiting calculator...")
			break // Exit the loop when user chooses option 5
		}

		var a, b int

		if first_input == "1" {
			fmt.Println("\nYou chose:", sum)
			fmt.Println(a_value)
			fmt.Scan(&a)
			fmt.Println(b_value)
			fmt.Scan(&b)
			result := sum_operation(a, b)
			fmt.Println("\nResult of your operation:", result)
			utils.Sleep(2)
			utils.Clear()

		} else if first_input == "2" {
			fmt.Println("\nYou chose:", subtract)
			fmt.Println(a_value)
			fmt.Scan(&a)
			fmt.Println(b_value)
			fmt.Scan(&b)
			result := subtract_operation(a, b)
			fmt.Println("\nResult of your operation:", result)
			utils.Sleep(2)
			utils.Clear()

		} else if first_input == "3" {
			fmt.Println("\nYou chose:", multiply)
			fmt.Println(a_value)
			fmt.Scan(&a)
			fmt.Println(b_value)
			fmt.Scan(&b)
			result := multiply_operation(a, b)
			fmt.Println("\nResult of your operation:", result)
			utils.Sleep(2)
			utils.Clear()

		} else if first_input == "4" {
			fmt.Println("\nYou chose:", divide)
			fmt.Println(a_value)
			fmt.Scan(&a)
			fmt.Println(b_value)
			fmt.Scan(&b)
			result := divide_operation(a, b)
			fmt.Println("\nResult of your operation:", result)
			utils.Sleep(2)
			utils.Clear()

		} else {
			fmt.Println("\nOps! Invalid input:", first_input, "\nPlease, try again!")
			utils.Sleep(2)
			utils.Clear()
		}
	}
}
