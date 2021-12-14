package main

import "fmt"


func main() {
	var inputFloat float64
	
	/*
		The commented lines below will print an error
		if the user inputs letters instead of numbers
	*/
	// input, err := fmt.Scan(&inputFloat)
	// if err != nil {
	// 	panic(err.Error())
	// }
	
	fmt.Scan(&inputFloat)
	inputInt := int(inputFloat)

	// fmt.Print("Number of inputs: ", input, ", with the input: ", inputFloat, ", truncated into: ", inputInt)
	fmt.Print(inputInt)
}