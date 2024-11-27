package main

import "fmt"

func main() {

	num := 1

	switch num {
	case 1:
		fmt.Println("Is 1")
	case 2, 3, 4, 5, 6, 7, 8, 9:
		fmt.Println("Is between 2 and 9")
	case 0:
		fmt.Println("Is 0")
	default:
		fmt.Println("Don't be so annoying, please!")
		break
		// fmt.Println("test") unreadeble code
	}

	//IN go we dont need to use break, it automatically breaks when matches any condition

}
