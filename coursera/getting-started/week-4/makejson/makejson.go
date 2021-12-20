/*
	Write a program which prompts the user to first enter a name, and then enter an address.
	Your program should create a map and add the name and address to the map
	using the keys “name” and “address”, respectively.
	Your program should use Marshal() to create a JSON object from the map,
	and then your program should print the JSON object.
*/
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	person := make(map[string]string)

	var name string
	var address string

	fmt.Print("Please enter a name: ")
	fmt.Scan(&name)
	person["name"] = name

	fmt.Print("Please enter an address: ")
	fmt.Scan(&address)
	person["address"] = address

	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))

	//If there is a better way to assign the scanned name and address into the map
	//other than storing it into a variable first like I did, please let me know.
	//Thank you.
}