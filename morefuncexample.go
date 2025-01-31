package main 

import "fmt"

func add(num1 int ,num2 int ) int { //num=10,num2=20

	sum := num1+num2

	return sum 
}

func getnumbers(num1 int ,num2 int)(int , int){
	
	sum := num1+num2//ans30
	
	mul:=num1*num2//ans200
	
	return sum,mul 
}

func printSomething(){
	fmt.Println("Education must be free!")
	
}
func sayHello(name string){
	fmt.Println("Welcome",name)
}
func main(){

	printSomething()
	sayHello("Anitya")
	
}