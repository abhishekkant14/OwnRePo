package main

import "fmt"
func main(){
arrey:=[5]int{1,2,3,4,5}

fmt.Println("Orignal Number ",arrey)

for i,a:=0,len(arrey)-1;i<a;i,a=i+1,a-1{
	arrey[i],arrey[a]=arrey[a],arrey[i]


}
    fmt.Println("Reversed Number ",arrey)
	//fmt.Println("Capacity of",cap(arrey))
	





}