package main

import "fmt"

func main() {

	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	vegetablesSlice := []string{"Carrot", "Potato", "Tomato", "Onion"}
	slice1 := vegetablesSlice[1:3]
	fmt.Println(slice1)
	slice1[0] = "Potato2"
	fmt.Println(vegetablesSlice)
}
