package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, World!")

	// Variables
	var myString string = "Hello, World!"
	myString2 := "Hello, End of the World!"
	myString = "Modificando el string"
	var myInt int = 42
	var myFloat float64 = 42.42
	var myBool bool = true
	myBool = false
	fmt.Println(myString, myString2, myInt, myFloat, myBool)
	fmt.Println(float64(myInt) + myFloat)
	fmt.Println(reflect.TypeOf(myInt))


	// Constants
	const myConst string = "Hello, Constant World!"
	fmt.Println(myConst)

	// Arrays
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray2 := [3]int{1, 2, 3}
	fmt.Println(myArray, myArray2)

	// Slices
	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)

	// Maps
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 2
	myMap2 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(myMap, myMap2)

	// Lists
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	// Flow control
	if myInt > 100 && myInt < 200 {
		fmt.Println("myInt is greater than 100 and less than 200")
	} else {
		fmt.Println("myInt is less than 100")
	}

	// Loops
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray {
		fmt.Println(index, value)
	}

	fmt.Println(myFunction())

	myStruct := MyStruct{"John", 25}
	fmt.Println(myStruct)
}

// Functions
func myFunction() string {
	return "My function was called"
}

// Structs
type MyStruct struct {
	name string
	age int
}