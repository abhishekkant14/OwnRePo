package main

import"fmt"

func main (){
	arr:=[6]int{1,2,3,43,45,46}
	max:=arr[0]

	for i:=1;i<len(arr);i++{
		if arr[i]>max {
			max=arr[i]
			
		}

	}
	fmt.Println("LargestNum",max)
}