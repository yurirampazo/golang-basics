package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	// fmt.Println("Hello World")
	// fmt.Println("go build hello.go + ./hello = go run hello.go")
	// showVariables()
	// showVariablesInfo()
	// readComands()
	// firstIfElse()
	// firstSwitch()
	// _, age := returnNameAndAge(123) // _ for ignoring undesirable return variables
	// println(age)
	startMonitoring()
	showNames()
}

func showVariables() {
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

func showVariablesInfo() {
	// := defines a variable without having to write var
	var name = "Yuri"
	age := 28
	// version := float32 = 1.1
	version := 1.1 /*will be float64*/
	var isVersion = true
	fmt.Println("Hello,", name, "your age is", age)
	fmt.Println("This programs version is", version, "isn't it?", isVersion)

	fmt.Println(name, "is", reflect.TypeOf(name))
	fmt.Println(age, "type is", reflect.TypeOf(age))
	fmt.Println(version, "is", reflect.TypeOf(version))
	fmt.Println(isVersion, "type is", reflect.TypeOf(isVersion))
}

func readComands() {
	/**Interacting with user*/
	for {
		fmt.Println("1- Start monitoring")

		fmt.Println("2- Show logs")
		fmt.Println("0- Quit")

		var instruction int
		// fmt.Scanf("%d", &instruction)
		fmt.Scan(&instruction) /*can just use the scan method instead*/
		//& stands for the pointer of variable in machine memory
		fmt.Println("The address in memory is", &instruction)
		fmt.Println("The choosen instruction was", instruction)
		if instruction == 0 {
			break
		}
	}
}

func firstIfElse() {
	num := 1

	if num > 0 {
		fmt.Printf("Number %d is greater than 0", num)
	} else if num < 0 {
		fmt.Printf("Number %d is lower than 0", num)
	} else {
		fmt.Printf("Number %d is equals 0", num)
	}

}

func firstSwitch() {
	num := 1

	switch num {
	case 1:
		fmt.Println("Is 1")
	case 2, 3, 4, 5, 6, 7, 8, 9:
		fmt.Println("Is between 2 and 9")
	case 0:
		fmt.Println("Is 0")
		os.Exit(0)
	default:
		fmt.Println("Don't be so annoying, please!")
		os.Exit(-1)
		// fmt.Println("test") unreadeble code
	}

	//IN go we dont need to use break, it automatically breaks when matches any condition

}

func returnNameAndAge(id int) (string, int) {
	fmt.Println("Calling database with id", id, "...")
	fmt.Println("Response name and age for the given ID!")
	return "name", 21
}

func startMonitoring() {
	sites := []string{
		"https://www.alura.com.br",
		"https://google.com"}
	fmt.Println("Starting monitoring...")
	response, err := http.Get(sites[0])
	fmt.Println("Response: ", response)
	fmt.Println("Error: ", err)

	response2, _ := http.Get(sites[1])
	fmt.Println("MOnitoring sites:", sites)

	if response2.StatusCode == 200 {
		fmt.Println("Site: ", sites[0], "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", sites[1],
			"está com problemas.Status code",
			response2.StatusCode)
	}
	fmt.Println(response2)
	fmt.Println(reflect.TypeOf(sites))


// for i:= 0; i < len(sites); i++ {
// 	fmt.Println(sites[i])
// }
for i, site := range sites {
	fmt.Println("Position", i, "of sites with site: ", site)
} 


for _, site := range sites {
	fmt.Println("Test of sites with site: ", site)
} 
}

func showNames() {
	names := []string{"katia" , "jana", "wilton", "jose", "maria", "dj", "wm"}
	fmt.Println(names) 
	fmt.Println(reflect.TypeOf(names))
	fmt.Println(len(names))
	names = append(names, "gaga")
	fmt.Println(names)
	fmt.Println(cap(names))	
}
