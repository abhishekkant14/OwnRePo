package main

import"fmt"
func main(){
	array:=[]int{10,23,24,26,25,78,56,89}
	Largest:=array[0]
	for i:=1;i<len(array);i++{
		if array[i]>Largest{
			Largest=array[i]
		}
	}
	fmt.Println(Largest)
}