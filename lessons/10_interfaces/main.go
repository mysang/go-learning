package main

import "fmt"

// 1. Interface declaration
type Speaker interface {
	Speak()
}

// 2. Types implementing the interface
type Person struct {
	Name string
	Age  int
}

func (p Person) Speak() {
	fmt.Println("Hello, my name is", p.Name)
}

type Dog struct{}

func (Dog) Speak() { fmt.Println("Woof!") }

type Cat struct{}

func (Cat) Speak() { fmt.Println("Meow!") }

// 4. Empty interface
func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

// 6. Type switch
func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

// 7. Interface with multiple methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	// 1. Basic interface
	var s Speaker = Person{Name: "Sang", Age: 33}
	s.Speak()

	// 2. Multiple implementations
	animals := []Speaker{Person{"Sang", 33}, Dog{}, Cat{}}
	for _, a := range animals {
		a.Speak()
	}

	// 3. Type information
	fmt.Printf("Type: %T\n", s)

	// 4. Empty interface
	describe(42)
	describe("Go")
	describe(Person{"Sang", 33})

	// 5. Type assertion
	var i interface{} = "hello"
	str, ok := i.(string)
	if ok {
		fmt.Println("String:", str)
	}

	// 6. Type switch
	checkType(10)
	checkType("hello")

	// 7. Multiple methods
	var shape Shape = Rectangle{10, 5}
	fmt.Println("Area:", shape.Area())
	fmt.Println("Perimeter:", shape.Perimeter())
}
