package main

import (
	"errors"
	"fmt"
	"math"
)

type room struct {
	door  string
	floor int
}

type person struct {
	name string
	age  int
	room room
}

func main() {
	firstSum := sum(1, 2)

	fmt.Println(firstSum)
	validSqrt, errorValidSqrt := sqrt(4)
	fmt.Println(validSqrt, errorValidSqrt)
	noValidSqrt, errornoValidSqrt := sqrt(-1)
	fmt.Println(noValidSqrt, errornoValidSqrt)

	executeSqrt(16)
	executeSqrt(-16)

	p := person{name: "Jhon", age: 5}
	fmt.Println(p)

	p2 := person{name: "Jane", age: 10, room: room{door: "D", floor: 2}}
	fmt.Println(p2)

	x := 0
	incBadCoded(x)
	fmt.Println(x)
	incPointer(&x)
	fmt.Println(x)
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func executeSqrt(x float64) {
	result, error := sqrt(x)
	if error == nil {
		fmt.Println(result)
	} else {
		fmt.Println(error)
	}
}

func incBadCoded(x int) {
	x++
}

func incPointer(x *int) {
	*x++
}
