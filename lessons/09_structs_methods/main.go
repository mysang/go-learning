package main

import "fmt"

// 1. Struct definition
type Person struct {
	Name string
	Age  int
}

// 2. Methods
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p Person) ShowAge() {
	fmt.Println("Age:", p.Age)
}

func (p *Person) IncreaseAge() {
	p.Age++
}

// 3. Rectangle struct with methods
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 4. Embedded struct example
type Contact struct {
	Email string
	Phone string
}

type Employee struct {
	Name string
	Contact
}

// 5. Method promotion via embedding
type Logger struct{}

func (Logger) Log(msg string) {
	fmt.Println("[LOG]", msg)
}

type Service struct {
	Logger
}

func main() {
	p := Person{"Sang", 33}
	p.Greet()
	p.ShowAge()
	p.IncreaseAge()
	p.ShowAge()

	rect := Rectangle{10, 5}
	fmt.Println("Area:", rect.Area())
	rect.Scale(2)
	fmt.Println("Scaled area:", rect.Area())

	e := Employee{Name: "Sang", Contact: Contact{Email: "sang@mail.com", Phone: "123"}}
	fmt.Println(e.Email)

	s := Service{}
	s.Log("Service started")
}
