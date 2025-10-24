package main

import (
	"fmt"
)

// 1. Simple function
func greet(name string) {
	fmt.Println("Hello,", name)
}

// 2. Function with return value
func add(a int, b int) int {
	return a + b
}

// 3. Multiple return values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// 4. Named return values
func rectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

// 5. Variadic function
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// 6. Pass by reference
func increase(x *int) {
	*x++
}

// 9. Defer example
func process() {
	defer fmt.Println("End of process")
	fmt.Println("Start of process")
}

func main() {
	greet("Sang")

	fmt.Println("Sum:", add(5, 7))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	a, p := rectangle(5, 3)
	fmt.Println("Area:", a, "Perimeter:", p)

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20, 30, 40))

	n := 10
	increase(&n)
	fmt.Println(n)

	greetFunc := func(name string) {
		fmt.Println("Hello,", name)
	}
	greetFunc("Go Learner")

	sumOp := func(x, y int) int { return x + y }
	productOp := func(x, y int) int { return x * y }

	fmt.Println(sumOp(10, 5), productOp(10, 5))

	process()
}
