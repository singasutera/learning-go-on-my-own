package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a string: ")

	//Scans for user input
	scanner.Scan()
    input := scanner.Text()

	//Creates a new string in which the user input is modified into lowercase
	inputInLowerCase := strings.ToLower(input)

	//Checks for prefix "i", suffix "n", and contains "a" all in one line
	if strings.HasPrefix(inputInLowerCase, "i") && strings.HasSuffix(inputInLowerCase, "n") && strings.Contains(inputInLowerCase, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}