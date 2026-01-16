package main

import "fmt"

func main() {
	arrey:= []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	reversed := []int{}

	for i := len(arrey) - 1; i >= 0; i-- {
		reversed = append(reversed, arrey[i])
	}
	fmt.Println("Orignal Number ",arrey)
	fmt.Println("Reversed Number ",reversed)
}
