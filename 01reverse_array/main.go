package main 


import"fmt"

func main(){
Arr:=[]int{1,2,3,4,5,6,7,8,9}

for i,a:=0,len(Arr)-1;i<a;i,a=1+i,a-1{
	Arr[i],Arr[a]=Arr[a],Arr[i]
}
fmt.Println("Reversed Arrar Number ",Arr)

}