//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	strnum := ""
	slice := make([]int, 3)

	for {
		fmt.Printf("Enter a no : ")
		fmt.Scan(&strnum)

		if strnum == "X" || strnum == "x" {
			fmt.Printf("\n\nLast Slice : %v\n", slice)
			break
		}

		intnum, err := strconv.Atoi(strnum)
		if err != nil {
			fmt.Print("Please Enter a valid Number\n")
			continue
		}

		slice = append(slice, intnum)
		sort.Ints(slice)
		fmt.Printf("Sorted Slice : %v\n", slice)
	}
}
