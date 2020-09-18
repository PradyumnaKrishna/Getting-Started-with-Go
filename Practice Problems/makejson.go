//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, add string
	fmt.Printf("Enter First Name : ")
	fmt.Scan(&name)
	fmt.Printf("Enter First Address : ")
	fmt.Scan(&add)

	p := map[string]string{
		"Name":    name,
		"Address": add,
	}

	j_obj, err := json.Marshal(p)
	if err == nil {
		fmt.Printf("\nCreated JSON object : %s\n", j_obj)
	}
}
