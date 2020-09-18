//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct{ name string }

func (c cow) Eat() {
	fmt.Printf("grass")
}
func (c cow) Move() {
	fmt.Printf("walk")
}
func (c cow) Speak() {
	fmt.Printf("moo")
}

type bird struct{ name string }

func (b bird) Eat() {
	fmt.Printf("worms")
}
func (b bird) Move() {
	fmt.Printf("fly")
}
func (b bird) Speak() {
	fmt.Printf("peep")
}

type snake struct{ name string }

func (s snake) Eat() {
	fmt.Printf("mice")
}
func (s snake) Move() {
	fmt.Printf("slither")
}
func (s snake) Speak() {
	fmt.Printf("hss")
}

func main() {
	var command, arg1, arg2 string
	animalMap := map[string][]string{}

	for {
		fmt.Printf("\n> ")
		fmt.Scan(&command)
		fmt.Scan(&arg1)
		fmt.Scan(&arg2)

		if strings.EqualFold(command, "newanimal") {
			animalMap[arg1] = append(animalMap[arg1], arg2)
			fmt.Printf("Created it!")
		} else if strings.EqualFold(command, "query") {
			var a Animal
			if strings.EqualFold(animalMap[arg1][0], "cow") {
				c := cow{name: arg1}
				a = c
			} else if strings.EqualFold(animalMap[arg1][0], "bird") {
				b := bird{name: arg1}
				a = b
			} else if strings.EqualFold(animalMap[arg1][0], "snake") {
				s := snake{name: arg1}
				a = s
			} else {
				fmt.Print("Wrong Name")
			}

			switch arg2 {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Print("Wrong Option")
			}
		} else {
			fmt.Print("Wrong Command")
		}
	}
}
