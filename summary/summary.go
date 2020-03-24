package main

import (
	"fmt"
)

func main() {
	var varInteger int
	var varString string
	var varBoolean bool
	var varFloat64 float64

	fmt.Println("Printing default values from GO lenguage:")
	fmt.Println(varInteger)
	fmt.Println(varString)
	fmt.Println(varBoolean)
	fmt.Println(varFloat64)

	var x int = 2
	var y int = 3
	var sum int = x + y
	fmt.Println(sum)

	// Also you can inferit the values without adding their type
	anotherXasInt := 5
	anotherYasInt := 10
	anotherSum := anotherXasInt + anotherYasInt
	fmt.Println(anotherSum)

	// on the simple conditions you can avoid parenthesis
	if anotherSum > 10 {
		fmt.Println("is bigger than 10")
	} else if anotherSum < 10 {
		fmt.Println("is less than 10")
	} else {
		fmt.Println("is 10")
	}

	var arrayInts [4]int
	arrayInts[2] = 3
	fmt.Println("Arrays")
	fmt.Println(arrayInts)
	arrayInitialized := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayInitialized)
	// arrayInts[5] = 3 -> this will fail.
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	fmt.Println(slice)

	fmt.Println("Maps")
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 4
	vertices["random"] = 88888
	fmt.Println(vertices)
	fmt.Println(vertices["random"])
	delete(vertices, "random")
	fmt.Println(vertices)

	fmt.Println("loops")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// for working as a while
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	for index, value := range slice {
		fmt.Println("index", index, "value", value)
	}
	for key, value := range vertices {
		fmt.Println("key", key, "value", value)
	}
}
