//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"fmt"
	"strconv"
)

func Swap(x []int, i int) {
	x[i], x[i+1] = x[i+1], x[i]
}

func BubbleSort(x []int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i-1; j++ {

			if x[j] > x[j+1] {
				Swap(x, j)
			}
		}
	}
}

func main() {
	var x []int
	strnum := ""

	for {
		fmt.Printf("Enter a no or type x to exit : ")
		fmt.Scan(&strnum)

		if strnum == "X" || strnum == "x" {
			fmt.Printf("Current Slice : %v\n", x)
			break
		}

		intnum, err := strconv.Atoi(strnum)
		if err != nil {
			fmt.Print("Please Enter a valid Number\n")
			continue
		}

		x = append(x, intnum)
	}

	BubbleSort(x)
	fmt.Printf("\nSorted Slice : %v\n", x)
}
