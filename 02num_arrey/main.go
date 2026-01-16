package main

import"fmt"
func main(){
	arrey:=[5]int{1,2,3,4,5}
	search:=5
	found:=false

	for i:=0;i<len(arrey);i++{
		if arrey[i]==search {
			defer fmt.Println("Found Index Number:",i)
			fmt.Println("Number Found",arrey[i])
found=true
			break
		}
	}
	if ! found{
		fmt.Println("O")
	}
}