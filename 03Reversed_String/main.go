package main

import"fmt"
func main(){
	str:="Vivek"
	Reversed:=""
fmt.Println("Name is", str)
	for i:=len(str)-1;i>=0;i--{
		
		Reversed+=string(str[i])

		
	}
	fmt.Println("Name is",Reversed)
	
}