# Lesson 08 — Functions in Go

A **function** is a reusable block of code that performs a specific task.
Functions help break complex programs into smaller, manageable, and testable parts.

---

## 1️⃣ Declaring a Function

General syntax:

```go
func functionName(parameters) returnType {
    // function body
}
```

Example:

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

Calling the function:

```go
greet("Sang")
```

Output:

```plaintext
Hello, Sang
```

---

## 2️⃣ Function with Return Value

Functions can return values using the `return` statement.

```go
func add(a int, b int) int {
    return a + b
}
```

Call:

```go
sum := add(5, 7)
fmt.Println("Sum:", sum)
```

Output:

```plaintext
Sum: 12
```

---

## 3️⃣ Multiple Return Values

Go supports returning multiple values — commonly used for error handling.

Example:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

Usage:

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

Output:

```plaintext
Result: 5
```

---

## 4️⃣ Named Return Values

You can name the return variables in the function signature.

```go
func rectangle(width, height int) (area int, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return
}
```

Usage:

```go
a, p := rectangle(5, 3)
fmt.Println("Area:", a, "Perimeter:", p)
```

---

## 5️⃣ Variadic Functions (Flexible Arguments)

Variadic functions accept a variable number of arguments.

```go
func sum(nums ...int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    return total
}
```

Usage:

```go
fmt.Println(sum(1, 2, 3))       // 6
fmt.Println(sum(10, 20, 30, 40)) // 100
```

💡 Note: The `...` allows passing any number of arguments of the same type.

---

## 6️⃣ Passing by Value vs Reference

In Go, function parameters are **passed by value** — a copy is made.
To modify the original variable, use **pointers**.

```go
func increase(x *int) {
    *x++
}
```

Usage:

```go
n := 10
increase(&n)
fmt.Println(n) // 11
```

---

## 7️⃣ Anonymous Functions (Function Literals)

Functions can be defined without a name and used immediately.

```go
func() {
    fmt.Println("This runs immediately!")
}()
```

Or assigned to a variable:

```go
greet := func(name string) {
    fmt.Println("Hello,", name)
}
greet("Go Learner")
```

---

## 8️⃣ Higher-Order Functions

Functions can accept other functions as parameters or return them.

```go
func operate(a, b int, fn func(int, int) int) int {
    return fn(a, b)
}
```

Usage:

```go
sum := operate(10, 5, func(x, y int) int { return x + y })
product := operate(10, 5, func(x, y int) int { return x * y })
fmt.Println(sum, product)
```

Output:

```plaintext
15 50
```

---

## 9️⃣ Defer Statement

`defer` delays the execution of a function until the surrounding function returns.
Useful for cleanup (closing files, releasing resources).

```go
func process() {
    defer fmt.Println("End of process")
    fmt.Println("Start of process")
}
```

Output:

```plaintext
Start of process
End of process
```

💡 Deferred calls are executed in **LIFO (Last In, First Out)** order.

---

## 🧠 Summary

- Use `func` to define reusable code blocks.
- Go supports multiple return values and named returns.
- Variadic functions accept a flexible number of arguments.
- Functions are first-class citizens — can be passed or returned.
- Use pointers when you need to modify variables.
- Use `defer` for cleanup logic.

---

## 🚀 Try It Yourself

Run:

```bash
go run ./lessons/08_functions
```

Expected output:

```plaintext
Hello, Sang
Sum: 12
Result: 5
Area: 15 Perimeter: 16
6
100
11
Hello, Go Learner
15 50
Start of process
End of process
```
