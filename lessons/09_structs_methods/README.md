# Lesson 09 — Structs & Methods in Go

Go does not have classes like other object-oriented languages,
but it provides **structs** (for data) and **methods** (for behavior).
Together, they form Go’s simple but powerful model for structuring data and logic.

---

## 1️⃣ Struct Basics

A **struct** is a collection of fields (variables) grouped together.

Example:

```go
type Person struct {
    Name string
    Age  int
}
```

Creating a struct:

```go
var p1 Person
p1.Name = "Sang"
p1.Age = 33
```

Or using a struct literal:

```go
p2 := Person{"Long", 28}
p3 := Person{Name: "Lam", Age: 30}
```

---

## 2️⃣ Accessing and Modifying Fields

Use the dot `.` operator:

```go
fmt.Println(p1.Name)
p1.Age = 34
```

Structs are **value types** → assigning a struct copies all its fields.

Example:

```go
p2 := p1
p2.Name = "Another"
fmt.Println(p1.Name) // Sang (unchanged)
```

---

## 3️⃣ Anonymous Structs

You can create a struct without defining a named type.

```go
student := struct {
    Name  string
    Score int
}{"Minh", 90}

fmt.Println(student)
```

Useful for temporary or inline data.

---

## 4️⃣ Pointers to Structs

Structs can be large, so passing them by pointer avoids copying.

```go
p := Person{"Sang", 33}
ptr := &p
fmt.Println(ptr.Name) // Go automatically dereferences the pointer
```

Modifying through pointer affects the original:

```go
ptr.Age = 34
fmt.Println(p.Age) // 34
```

---

## 5️⃣ Structs in Functions

Functions can accept structs or pointers to structs.

```go
func printPerson(p Person) {
    fmt.Println(p.Name, p.Age)
}

func increaseAge(p *Person) {
    p.Age++
}
```

---

## 6️⃣ Methods in Go

A **method** is a function with a *receiver* — the variable it operates on.

Syntax:

```go
func (receiver ReceiverType) methodName(params) returnType {
    // code
}
```

Example:

```go
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}
```

Usage:

```go
p := Person{"Sang", 33}
p.Greet() // Hello, my name is Sang
```

---

## 7️⃣ Pointer Receiver vs Value Receiver

You can define methods on both **value** and **pointer** receivers.

### Value Receiver (copy)

```go
func (p Person) ShowAge() {
    fmt.Println("Age:", p.Age)
}
```

Modifying inside this method does **not** change the original struct.

### Pointer Receiver (reference)

```go
func (p *Person) IncreaseAge() {
    p.Age++
}
```

This method **does modify** the original instance.

---

## 8️⃣ Combining Structs and Methods

Example:

```go
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
```

Usage:

```go
rect := Rectangle{10, 5}
fmt.Println("Area:", rect.Area())
rect.Scale(2)
fmt.Println("Scaled area:", rect.Area())
```

Output:

```plaintext
Area: 50
Scaled area: 200
```

---

## 9️⃣ Embedding (Composition)

Go uses **embedding** instead of inheritance.

Example:

```go
type Contact struct {
    Email string
    Phone string
}

type Employee struct {
    Name string
    Contact // embedded struct
}

e := Employee{Name: "Sang", Contact: Contact{Email: "sang@mail.com", Phone: "123"}}
fmt.Println(e.Email) // direct access due to embedding
```

---

## 1️⃣0️⃣ Method Promotion via Embedding

Embedded structs also promote their methods to the outer struct.

Example:

```go
type Logger struct{}

func (Logger) Log(msg string) {
    fmt.Println("[LOG]", msg)
}

type Service struct {
    Logger
}

s := Service{}
s.Log("Service started") // works! method promoted automatically
```

---

## 🧠 Summary

- Structs = custom data types.
- Methods = functions bound to structs.
- Value receivers → work on copies.
- Pointer receivers → modify the original.
- Embedding → composition instead of inheritance.
- Methods are promoted through embedded structs.

---

## 🚀 Try It Yourself

Run:

```bash
go run ./lessons/09_structs_methods
```

Expected output:

```plaintext
Hello, my name is Sang
Age: 33
Age: 34
Area: 50
Scaled area: 200
sang@mail.com
[LOG] Service started
```
