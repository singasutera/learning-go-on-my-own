/*
	Write a Bubble Sort program in Go. The program
	should prompt the user to type in a sequence of up to 10 integers. The program
	should print the integers out on one line, in sorted order, from least to
	greatest. Use your favorite search tool to find a description of how the bubble
	sort algorithm works.

	As part of this program, you should write a
	function called BubbleSort() which
	takes a slice of integers as an argument and returns nothing. The BubbleSort() function should
	modify the slice so that the elements are in sorted
	order.

	A recurring operation in the bubble sort algorithm is
	the Swap operation which swaps the position of two adjacent elements in the
	slice. You should write a Swap() function which performs this operation. Your Swap()
	function should take two arguments, a slice of integers and an index value i which
	indicates a position in the slice. The Swap() function should return nothing, but it should swap
	the contents of the slice in position i with the contents in position i+1.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputStr string
	var intSlice []int
	var outputSlice []string

	fmt.Println("Please type in up to 10 integer numbers, each one separated by a coma:")
	fmt.Scan(&inputStr)

	strSlice := strings.Split(inputStr, ",")
	for i := range strSlice {
		intValue, _ := strconv.Atoi(strSlice[i])
		intSlice = append(intSlice, intValue)
	}
	
	BubbleSort(intSlice)

	for i := range intSlice {
		strValue := strconv.Itoa(intSlice[i])
		outputSlice = append(outputSlice, strValue)
	}
	outputStr := strings.Join(outputSlice, ",")
	
	fmt.Println("Here are those numbers in ascending order:")
	fmt.Println(outputStr)
}

//Swaps the contents of the slice at index i with index i+1
func Swap(sliceToModify []int, i int) {
	temp := sliceToModify[i]
	sliceToModify[i] = sliceToModify[i+1]
	sliceToModify[i+1] = temp
}

//Sorts a given slice of integers in an ascending order
func BubbleSort(sliceToSort []int) {
	lengthOfSlice := len(sliceToSort)
	for i := 0; i < lengthOfSlice; i++ {
		for j := 0; j < lengthOfSlice-i-1; j++ {
			if sliceToSort[j] > sliceToSort[j+1] {
				Swap(sliceToSort, j)
			}
		}
	}
}
