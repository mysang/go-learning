# Lesson 10 — Interfaces in Go

An **interface** defines a set of method signatures that a type must implement.
Any type that implements those methods **automatically satisfies** the interface — no explicit declaration is needed.

This is called **structural typing** or **duck typing**:
“If it walks like a duck and quacks like a duck, it’s a duck.”

---

## 1️⃣ Declaring an Interface

Example:

```go
type Speaker interface {
    Speak()
}
```

Any type that has a `Speak()` method automatically implements this interface.

Example struct:

```go
type Person struct {
    Name string
}

func (p Person) Speak() {
    fmt.Println("Hello, my name is", p.Name)
}
```

Usage:

```go
var s Speaker
s = Person{Name: "Sang"}
s.Speak()
```

Output:

```plaintext
Hello, my name is Sang
```

💡 Go does **not require explicit "implements" keyword**.
If the methods match → it just works.

---

## 2️⃣ Multiple Types Implementing the Same Interface

```go
type Dog struct{}
func (Dog) Speak() { fmt.Println("Woof!") }

type Cat struct{}
func (Cat) Speak() { fmt.Println("Meow!") }
```

Usage:

```go
animals := []Speaker{Person{"Sang"}, Dog{}, Cat{}}
for _, a := range animals {
    a.Speak()
}
```

Output:

```plaintext
Hello, my name is Sang
Woof!
Meow!
```

---

## 3️⃣ Interface Values and Dynamic Types

When you assign a value to an interface variable, Go stores:

- The **value itself**, and
- The **type information**.

You can inspect it with `fmt.Printf("%T", var)`:

```go
var s Speaker = Person{"Sang"}
fmt.Printf("Type: %T\n", s)
```

Output:

```plaintext
Type: main.Person
```

---

## 4️⃣ Empty Interface — `interface{}`

An empty interface can hold **any type**.

```go
var x interface{}
x = 10
x = "Hello"
x = true
```

This is similar to `any` in TypeScript or `object` in C#.

You can use it for generic containers, JSON parsing, etc.

Example:

```go
func describe(i interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", i, i)
}

describe(42)
describe("Go")
describe(Person{"Sang"})
```

---

## 5️⃣ Type Assertion

If you know the concrete type inside an interface, use **type assertion**:

```go
var i interface{} = "hello"
s := i.(string)
fmt.Println(s) // "hello"
```

If type does not match → panic.
To prevent panic, use the safe form:

```go
s, ok := i.(string)
if ok {
    fmt.Println("String:", s)
} else {
    fmt.Println("Not a string")
}
```

---

## 6️⃣ Type Switch

Type switches allow checking the type stored in an interface variable.

```go
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
```

Usage:

```go
checkType(10)
checkType("hello")
```

---

## 7️⃣ Interfaces with Multiple Methods

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

Any type implementing both methods satisfies the interface.

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
```

Usage:

```go
var s Shape = Rectangle{10, 5}
fmt.Println("Area:", s.Area())
fmt.Println("Perimeter:", s.Perimeter())
```

---

## 8️⃣ Embedding Interfaces

Interfaces can be composed together:

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string)
}

type ReadWriter interface {
    Reader
    Writer
}
```

Now any type implementing both `Read()` and `Write()` will satisfy `ReadWriter`.

---

## 9️⃣ Nil Interface Pitfall

If an interface value is **nil**, or holds a nil concrete value,
you must understand the difference:

```go
var s Speaker
fmt.Println(s == nil) // true
```

But if an interface holds a *typed nil*:

```go
var p *Person = nil
var sp Speaker = p
fmt.Println(sp == nil) // false
```

Because the interface holds type info even though the value is nil.

---

## 🧠 Summary

- Interfaces define **behavior**, not data.
- Implementation is **implicit** — no keyword needed.
- The **empty interface** (`interface{}`) can store any value.
- Use **type assertion** or **type switch** to extract concrete types.
- Interfaces can be **embedded**.
- A nil interface ≠ an interface holding a typed nil.

---

## 🚀 Try It Yourself

Run:

```bash
go run ./lessons/10_interfaces
```

Expected output (order may vary):

```plaintext
Hello, my name is Sang
Woof!
Meow!
Type: main.Person
Type: int, Value: 42
Type: string, Value: Go
Type: main.Person, Value: {Sang 33}
String: hello
Integer: 10
String: hello
Area: 50
Perimeter: 30
```
