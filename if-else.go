package main

import "fmt"

func main() {
	age :=2
	if age > 18 {
        fmt.Println("You are eligible to be married.")
    } else if age <18 {
        fmt.Println("You are not eligible to be married, but you can love someone.")
    } else {
        fmt.Println("You are 18 years old,which is often considered the age of  eligibility for marriage.")
        
    }
}
