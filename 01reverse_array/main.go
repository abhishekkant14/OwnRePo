package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
fmt.Println("Number",arr)
	for i, a := 0, len(arr)-1; i < a; i, a = i+1, a-1 {
		arr[i], arr[a] = arr[a], arr[i]
	}
	
	fmt.Println("Reverse arr",arr)
}
