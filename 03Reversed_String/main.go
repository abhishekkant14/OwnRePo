package main

import"fmt"

func main(){
Name:="Abhishek"
Reversed:=""

for i:=0;i<len(Name);i++{
Reversed=string(Name[i])+Reversed
}
fmt.Println("Reverse Name ",Reversed)

}