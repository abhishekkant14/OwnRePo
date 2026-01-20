package main


import"fmt"
func main(){
arr:=[]int{1,22,34,54,57,65,80}

max:=arr[0]

for i:=1;i<len(arr);i++{
	if arr[i]>max{
		max=arr[i]
	}
}
fmt.Println("Largest number is thr Array Is",max)

}