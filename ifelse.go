package main

import "fmt"

func main() {
	var x = 40
	if x < 15 {
		fmt.Println("x is less than 15")
	} else if x == 15 {
		fmt.Println("x equals 15")
	} else {
		fmt.Println("x is more than 15")
	}
}