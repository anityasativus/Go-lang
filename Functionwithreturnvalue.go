package main

import "fmt"

func add(num1 int, num2 int) int { //num=10,num2=20

	sum := num1 + num2

	return sum
}

func getnumbers(num1 int, num2 int) (int, int) {

	sum := num1 + num2 //ans30

	mul := num1 * num2 //ans200

	return mul ,  sum
}

func main() {

	a := 10
	b := 20
	b := 200
	p, q := getnumbers(a, b) //p is sum N q is mul

	fmt.Println( p )

	fmt.Println( q )
}
