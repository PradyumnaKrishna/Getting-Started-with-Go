//usr/bin/env go run "$0" "$@"; exit

package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Printf("Food Eaten : %s\n", a.food)
}

func (a Animal) Move() {
	fmt.Printf("Locomotion Mehod : %s\n", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Printf("Spoken sound : %s\n", a.noise)
}

func main() {

	for {
		var animal Animal
		var name, info string
		fmt.Printf("> ")
		fmt.Scan(&name)
		fmt.Scan(&info)

		switch name {
		case "cow":
			animal = Animal{"grass", "walk", "moo"}
		case "bird":
			animal = Animal{"worms", "fly", "peep"}
		case "snake":
			animal = Animal{"mice", "slither", "hsss"}
		default:
			animal = Animal{"nothing", "no motion", "no sound"}
		}

		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("Wrong Details\n")
		}
	}
}
