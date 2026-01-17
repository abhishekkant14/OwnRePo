package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	reversed := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	fmt.Println("Orignal Number:", arr)
	fmt.Println("Append Number is", reversed)
}
