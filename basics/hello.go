package main

import (
	"fmt"
	"strings"
)

var age int //declaration without value have default values

var split_main string = strings.Repeat("-", 10)
var split_func string = strings.Repeat("%", 10)

func main() {
	fmt.Println(split_main, "Program start", split_main)
	var hello string = "Hello"
	world := "World!"
	fmt.Printf("%s, %s\n", hello, world)

	basics()
	fmt.Println(split_main, "Program end", split_main)
}

func basics() {
	fmt.Println(split_func, "Start func: basic", split_func)

	//comment

	var first_name string = "Ira" //normal variable
	last_name := "Ramos"          //inferred type variable, can only be used INSIDE functions

	age = 19

	var day, month, year int = 7, 3, 2024 //declaration of multiple vars

	//honestly there's more basics but im lazy so hehehehe

	fmt.Printf("I am %s %s\nToday's date is %d-%d-%d\n", first_name, last_name, day, month, year) //format where %s is string, %d is int and %f is float
	//additionally %v is value and %T is type

	const THIS_IS_A_CONSTANT string = "I will forever love programming"
	fmt.Println(THIS_IS_A_CONSTANT)

	//arrays
	var list_of_nums = [...]int{1, 2, 3} //you can replace ... with a number for an fixed array or just leave it blank
	fmt.Println("Looping through a list of numbers...")
	for number := range list_of_nums { //loop through list the smart way 😎
		println(number)
	}

	var list_index = []string{"Not it", "Found me!", "Not it again..."}
	println(list_index[1])

	//update list
	list_index[0] = "Value Updated!"
	fmt.Println(list_index)

	//initialize specific parts of a fixed array
	specific_index := [3]string{1: "second element", 2: "third element"}
	fmt.Println(specific_index)

	//len(arr) to find array size

	//switch case - both expressions must be of same type
	days_of_the_week := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	day_int := 0

	switch day_int {
	case 0:
		fmt.Println("Today is:", days_of_the_week[day_int])
	default:
		fmt.Println("Not in range")
	}

	fmt.Println(split_func, "End func: basic", split_func)
}
