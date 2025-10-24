package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Boolean
	isActive := true

	// Integer and float
	age := 33
	temperature := 36.6

	// String
	message := "Hello, Go!"

	// Byte and Rune
	b := byte('A')
	r := rune('語')

	// Output values
	fmt.Println("isActive:", isActive)
	fmt.Println("age:", age)
	fmt.Println("temperature:", temperature)
	fmt.Println("message:", message)
	fmt.Println("b (byte):", b)
	fmt.Println("r (rune):", string(r))

	// Reflect type
	fmt.Println("Type of age:", reflect.TypeOf(age))
	fmt.Println("Type of temperature:", reflect.TypeOf(temperature))
	fmt.Println("Type of message:", reflect.TypeOf(message))
}
