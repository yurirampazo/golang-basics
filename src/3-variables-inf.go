package main


import "fmt"
import "reflect"

func main() {
	// := defines a variable without having to write var
	var name = "Yuri"
	age := 28
	// version := float32 = 1.1
	version := 1.1  /*will be float64*/
	var isVersion = true
	fmt.Println("Hello,", name, "your age is", age)
	fmt.Println("This programs version is", version, "isn't it?", isVersion)

	fmt.Println(name, "is", reflect.TypeOf(name))
	fmt.Println(age, "type is", reflect.TypeOf(age))
	fmt.Println(version, "is", reflect.TypeOf(version))
	fmt.Println(isVersion, "type is", reflect.TypeOf(isVersion))
}