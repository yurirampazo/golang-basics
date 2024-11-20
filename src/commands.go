package main

import "fmt"

func main() {
		/**Interacting with user*/
		fmt.Println("1- Start monitoring")
		fmt.Println("2- Show logs")
		fmt.Println("0- Quit")
	
		var instruction int
		// fmt.Scanf("%d", &instruction)
		fmt.Scan(&instruction) /*can just use the scan method instead*/
		//& stands for the pointer of variable in machine memory
		fmt.Println("The address in memory is", &instruction)
		fmt.Println("The choosen instruction was", instruction)
}