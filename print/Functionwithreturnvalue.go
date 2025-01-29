package main 

import "fmt"

func add(num1 int ,num2 int ) int { //num=10,num2=20

	sum := num1+num2

	return sum 
}
func getnumbers(num1 int ,num2 int)(int , int){
	sum := num1+num2
}
func main(){

	a:=10
	b:=20

	sum := add(a,b)

	fmt.Println(sum)

}