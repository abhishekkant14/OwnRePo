package main

import"fmt"

func main(){
arr:=[]int{11,24,12,45,34,55,78,89}

largestNum:=arr[0]

for i:=1;i<len(arr);i++{
if arr[i]>largestNum{
	largestNum=arr[i]
}

}

fmt.Println("LargestNum",largestNum)

}