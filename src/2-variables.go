package main


import "fmt"

func main() {
	var name string = "Yuri"
	var age int = 28
	var version float32
	var isVersion bool = true
	fmt.Println("Hello,", name, "your age is", age)
	fmt.Println("This programs version is", version, "isn't it?", isVersion)

	//Default go values for data types:
	/**
	string -> empty
	int, float -> 0
	bool -> false
	*/
}