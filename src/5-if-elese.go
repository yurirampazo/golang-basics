package main

import "fmt"

func main() {

	num := 1

	if num > 0 {
		fmt.Printf("Number %d is greater than 0", num)
	} else if num < 0 {
		fmt.Printf("Number %d is lower than 0", num)
	} else {
		fmt.Printf("Number %d is equals 0", num)
	}

}
