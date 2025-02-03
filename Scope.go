package main

import (
	"archive/zip"
	"fmt"
)

var a = 20
var b = 30

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}
func main() {
	p := 30
	q := 40
	add(p, q)

	add(a, b)

	add(a, p)
	add(a, q) add(a, b)

}
