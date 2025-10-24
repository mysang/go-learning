# Lesson 02 — Data Types in Go

Data types define the kind of values a variable can hold and what operations can be performed on them.
Go is a **statically typed** language, so the compiler checks types at compile time.

---

## 1️⃣ Basic Data Types

Go provides several built-in data types categorized as follows:

### a) Boolean

```go
var isActive bool = true
var isAdmin bool = false
```

The zero value of a bool is `false`.

---

### b) Numeric Types

#### Integers

- Signed integers: int, int8, int16, int32, int64
- Unsigned integers: uint, uint8, uint16, uint32, uint64

Example:

```go
var age int = 33
var distance uint = 120
```

💡 Tip:
`int` and `uint` are **platform-dependent** —
on a 64-bit system, they are usually 64 bits wide.

---

#### Floating Point Numbers

```go
var temperature float32 = 36.6
var pressure float64 = 1013.25
```

`float64` is preferred for precision.

---

#### Complex Numbers

Go has built-in support for complex numbers:

```go
var c1 complex64 = 1 + 2i
var c2 complex128 = complex(5, 6)
fmt.Println(c1 + c2)
```

---

### c) Strings

Strings are sequences of bytes (UTF-8 encoded).

```go
var message string = "Hello, Go!"
greeting := "Xin chào"
```

💡 Strings in Go are **immutable** — once created, their content cannot be changed.

To access individual characters:

```go
fmt.Println(message[0])        // prints byte value
fmt.Println(string(message[0])) // prints "H"
```

---

### d) Byte and Rune

- `byte` → alias for `uint8`, used for raw data or ASCII.
- `rune` → alias for `int32`, used for Unicode characters.

Example:

```go
var b byte = 'A'     // 65
var r rune = '語'    // Unicode character
```

---

## 2️⃣ Type Conversion

Go does **not** perform implicit type conversion.
You must explicitly cast types:

```go
var a int = 10
var b float64 = float64(a)
var c int = int(b)
```

Example of incorrect (compile error):

```go
var x int = 10
var y float64 = x // ❌ cannot use int as float64
```

---

## 3️⃣ Type Inference and the `reflect` Package

When using `:=`, Go automatically infers the type:

```go
name := "Sang"    // string
age := 33         // int
isActive := true  // bool
```

You can check the type at runtime using `reflect.TypeOf()`:

```go
import "reflect"
fmt.Println(reflect.TypeOf(name)) // string
```

---

## 4️⃣ Zero Values for Each Type

| Type      | Zero Value |
|------------|------------|
| bool       | false      |
| string     | ""         |
| numeric    | 0          |
| pointer    | nil        |
| slice, map | nil        |

---

## 5️⃣ Best Practices

- Prefer `int` and `float64` unless you have specific size constraints.
- Always convert explicitly when mixing types.
- Use `rune` when handling non-ASCII characters.
- Use `const` for values that don’t change.

---

## 🚀 Try It Yourself

Run the following command:

```bash
go run ./lessons/02_data_types
```

Expected output:

```plaintext
isActive: true
age: 33
temperature: 36.6
message: Hello, Go!
b (byte): 65
r (rune): 語
Type of age: int
Type of temperature: float64
Type of message: string
```
