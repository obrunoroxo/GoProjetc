package main

import (
	"fmt"
	"goprojects/FourthCode/methods"
	"goprojects/utils"
	// "github.com/fatih/color"
)

// type Person struct {
// 	Name     string
// 	LastName string
// 	Age      int
// 	Email    string
// }

// func (p *Person) ShowMe() {
// 	utils.Clear()
// 	fmt.Println("Name:", p.Name)
// 	fmt.Println("Last Name:", p.LastName)
// 	fmt.Println("Age:", p.Age)
// 	fmt.Println("E-mail:", p.Email)
// }

// func (p *Person) GetName() string {
// 	return p.Name
// }

// func (p *Person) SetName(name string) string {
// 	green := color.New(color.FgGreen, color.Bold).SprintFunc()
// 	p.Name = name
// 	fmt.Printf("Name chaged with success %s", green("✓"))
// 	fmt.Println("Name:", name)
// 	return p.Name
// }

// func (p *Person) GetAge() int {
// 	return p.Age
// }

// func (p *Person) SetAge(age int) int {
// 	green := color.New(color.FgGreen, color.Bold)
// 	p.Age = age
// 	green.Println("Age chaged with success ", "✓")
// 	fmt.Println("Age:", age)
// 	return p.Age
// }

// func (p *Person) GetEmail() string {
// 	return p.Email
// }

// func (p *Person) SetEmail(email string) string {
// 	green := color.New(color.FgGreen, color.Bold)
// 	p.Email = email
// 	green.Println("Email chaged with success ", "✓")
// 	fmt.Println("Email:", email)
// 	return p.Email
// }

func main() {
	// Person := methods.Person

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

	mydatas := methods.Person{
		Name:     name,
		LastName: lastName,
		Age:      age,
		Sex:      sex,
	}

	methods.ShowMe(mydatas)
	utils.Sleep(1)
	teste := mydatas.SetName(name)
	fmt.Println(teste)
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
