//usr/bin/env go run "$0" "$@"; exit

package main

import "fmt"

func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	sf := func(t float64) float64 {
		return (1 / 2 * (a * t * t)) + (v * t) + s
	}
	return sf
}

func main() {
	var a, v, s, t float64
	fmt.Printf("Enter Acceleration : ")
	fmt.Scan(&a)
	fmt.Printf("Enter Velocity : ")
	fmt.Scan(&v)
	fmt.Printf("Enter Initial Displacement : ")
	fmt.Scan(&s)

	fn := GenDisplaceFn(a, v, s)

	fmt.Printf("Enter Time : ")
	fmt.Scan(&t)

	fmt.Printf("Final Displacement : %f\n", fn(t))
}
