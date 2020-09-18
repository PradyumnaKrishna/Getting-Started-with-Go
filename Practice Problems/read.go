//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Enter File Name: ")
	names := ""
	fmt.Scan(&names)

	file, _ := os.Open(names)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())

	}

	file.Close()

	nameSlice := make([]name, 0, len(txtlines))

	for _, v := range txtlines {
		someString := v
		words := strings.Fields(someString)
		nameSlice = append(nameSlice, name{words[0], words[1]})
	}

	for _, j := range nameSlice {
		fmt.Println(j.fname, j.lname)
	}
}
