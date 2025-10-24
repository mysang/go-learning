# Lesson 01 — Variables in Go

This lesson introduces one of the most fundamental concepts in Go — variables.

---

## 1️⃣ What is a Variable?

A variable is a named storage location used to hold a value that may change during program execution.

In Go, variables must be declared before use. Go is a statically typed language, meaning the type of each variable is known at compile time.

---

## 2️⃣ Declaring Variables

There are multiple ways to declare variables in Go:

### a) Using the `var` keyword

You can declare a variable with an explicit type:

```go
var name string = "Sang"
var age int = 33
```

If you omit the value, Go assigns the zero value automatically:

```go
var city string  // empty string ""
var number int   // 0
var isActive bool // false
```

---

### b) Type Inference

You can let Go automatically determine the type using `:=`:

```go
name := "Sang"
age := 33
```

This syntax can only be used **inside functions**, not at the package level.

---

### c) Declaring Multiple Variables

You can declare several variables at once:

```go
var a, b, c int = 1, 2, 3
```

Or even mix types with inference:

```go
x, y, message := 10, 20, "Hello"
```

---

## 3️⃣ Constants

A constant (`const`) is an immutable value — once defined, it cannot change.

```go
const Pi = 3.14159
const AppName string = "Go Learning"
```

Constants are evaluated at compile time, not runtime.

---

## 4️⃣ Naming Conventions and Best Practices

- Variable names should start with a **letter or underscore**, never a number.
- Go prefers **camelCase** for local variables.
- For exported identifiers (accessible from other packages), the name must start with a **capital letter**.
- Avoid using short names unless in small scopes (like loop variables).

Examples:

```go
var totalPrice int
var userName string
const MaxRetries = 3
```

---

## 5️⃣ Zero Values in Go

If you declare a variable without initializing it, Go automatically assigns a *zero value*:

| Type        | Zero Value |
|--------------|------------|
| string       | ""         |
| int, float   | 0          |
| bool         | false      |
| pointer, slice, map | nil |

Example:

```go
var score int
fmt.Println(score) // prints 0
```

---

## 6️⃣ Summary

- Variables in Go are **statically typed**.
- Use `var` for explicit declarations.
- Use `:=` for short-hand inside functions.
- Use `const` for immutable values.
- Uninitialized variables get a *zero value*.
- Prefer meaningful names for readability.

---

## 🚀 Try It Yourself

Run the following command inside your container:

```go
go run ./lessons/01_variables
```

Expected output:

```plaintext
Name: Sang
Age: 33
City: Ho Chi Minh
Year: 2025
Pi: 3.14
```
