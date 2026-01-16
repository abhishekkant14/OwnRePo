package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	search := 5	
	found := false

	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			fmt.Println("Index Number Is", i)
			fmt.Println("Number Is", arr[i])
			found = true
			break
		}
	}
	if ! found {
		fmt.Println("o")

	}
}
