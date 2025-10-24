# Lesson 04 — Conditional Statements in Go

Conditional statements allow your program to make decisions and execute code selectively based on conditions.

Go provides several ways to control logic flow:

- `if`, `if-else`, and `if-else if`
- `switch`
- short variable declarations inside conditions

---

## 1️⃣ The `if` Statement

Basic structure:

```go
if condition {
    // code executes if condition is true
}
```

Example:

```go
isRaining := true

if isRaining {
    fmt.Println("Bring an umbrella!")
}
```

If the condition is `false`, the block is skipped.

---

## 2️⃣ `if-else` Statement

When you need alternative actions:

```go
temperature := 30

if temperature > 35 {
    fmt.Println("It's too hot!")
} else {
    fmt.Println("Weather is fine.")
}
```

---

## 3️⃣ `if-else if` Ladder

Multiple conditions can be chained:

```go
score := 85

if score >= 90 {
    fmt.Println("Excellent")
} else if score >= 75 {
    fmt.Println("Good")
} else if score >= 50 {
    fmt.Println("Average")
} else {
    fmt.Println("Fail")
}
```

💡 Go evaluates conditions top to bottom and stops at the first match.

---

## 4️⃣ Short Variable Declaration in `if`

You can declare and use a variable *only within the if statement*.

Syntax:

```go
if v := math.Pow(2, 3); v < 10 {
    fmt.Println("Small number")
} else {
    fmt.Println("Big number")
}
```

The variable `v` exists **only inside** this if-else scope.

---

## 5️⃣ The `switch` Statement

`switch` simplifies complex if-else chains.

### Basic usage

```go
day := "Tuesday"

switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("Weekend is coming!")
default:
    fmt.Println("Midweek day")
}
```

💡 Unlike many languages, Go’s `switch` **automatically breaks** after each case.
You don’t need `break`.

---

## 6️⃣ Multiple Values in a Case

You can match multiple values in one case:

```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Weekday")
}
```

---

## 7️⃣ Expressionless Switch (Switch True)

You can use `switch` without an expression — similar to if-else chains:

```go
age := 25

switch {
case age < 13:
    fmt.Println("Child")
case age < 20:
    fmt.Println("Teenager")
case age < 60:
    fmt.Println("Adult")
default:
    fmt.Println("Senior")
}
```

---

## 8️⃣ Fallthrough Keyword

To explicitly continue to the next case, use `fallthrough`:

```go
num := 2

switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
}
```

Output:

```plaintext
    One
    Two
    Three
```

---

## 9️⃣ Nested Conditions

You can nest if statements, but deep nesting is discouraged for readability.

```go
if isMember {
    if age >= 18 {
        fmt.Println("Adult member")
    } else {
        fmt.Println("Minor member")
    }
}
```

Use `switch` or separate functions for cleaner code when logic grows complex.

---

## 🧠 Summary

- Use `if`, `else`, `else if` for simple branching.
- Use short declaration (`if v := ...;`) for temporary values.
- Use `switch` for multiple discrete conditions.
- No need for `break` in Go `switch`.
- Use `fallthrough` only when necessary.

---

## 🚀 Try It Yourself

Run:

```bash
go run ./lessons/04_conditionals
```

Expected output:

```plaintext
Bring an umbrella!
Weather is fine.
Good
Small number
Weekend is coming!
Adult
One
Two
Three
```
