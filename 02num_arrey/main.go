package main 
import"fmt"

func main(){
arrey:=[5]int{1,2,3,4,5}
search:=5
found:=false

for i:=0; i<len(arrey);i++{
	if arrey[i]==search{
		found=true
		defer fmt.Println("index Number ",i)
		fmt.Println("Number",arrey[i])
		break
	}
}
if ! found{
	fmt.Println("Number not found")}

}