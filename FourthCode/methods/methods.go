package methods

import (
	"fmt"
	"goprojects/utils"

	"github.com/fatih/color"
)

type Person struct {
	Name     string
	LastName string
	Age      int
	Sex      string
}

func ShowMe(p Person) {
	utils.Clear()
	fmt.Println("Name:", p.Name)
	fmt.Println("Last Name:", p.LastName)
	fmt.Println("Age:", p.Age)
	fmt.Println("Sexo:", p.Sex)
}

func GetName(p Person) string {
	return p.Name
}

func (p *Person) SetName(name string) string {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	p.Name = name
	fmt.Printf("Name changed with success %s", green("✓\n"))
	fmt.Println("Name:", name)
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) SetAge(age int) int {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	p.Age = age
	fmt.Printf("Age changed with success %s", green("✓"))
	fmt.Println("Age:", age)
	return p.Age
}

func (p *Person) GetSex() string {
	return p.Sex
}

func (p *Person) SetSex(sex string) string {
	green := color.New(color.FgGreen, color.Bold)
	p.Sex = sex
	green.Println("Sex chaged with success ", "✓")
	fmt.Println("Sex:", sex)
	return p.Sex
}
