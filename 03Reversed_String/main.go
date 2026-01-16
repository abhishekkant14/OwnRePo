package main

import"fmt"

func main(){
Name:="Vivek"
ReversedName:=""



for i:=0;i<len(Name);i++{
	ReversedName=string(Name[i])+ReversedName
}
 
fmt.Println("Reversed Name is",ReversedName)



}