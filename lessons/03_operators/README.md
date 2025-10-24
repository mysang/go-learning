# Lesson 03 — Operators in Go

Operators are symbols that perform operations on variables and values.
Go supports arithmetic, relational, logical, assignment, and bitwise operators.

---

## 1️⃣ Arithmetic Operators

Arithmetic operators are used for mathematical calculations.

| Operator | Description | Example |
|-----------|-------------|----------|
| + | Addition | a + b |
| - | Subtraction | a - b |
| * | Multiplication | a * b |
| / | Division | a / b |
| % | Modulus (remainder) | a % b |

Example:

```go
a := 10
b := 3
fmt.Println(a + b) // 13
fmt.Println(a - b) // 7
fmt.Println(a * b) // 30
fmt.Println(a / b) // 3 (integer division)
fmt.Println(a % b) // 1
```

💡 Note:
If both operands are integers, division results in an integer.
Use float64 if you want a decimal result:

```go
x := 10.0
y := 3.0
fmt.Println(x / y) // 3.3333
```

---

## 2️⃣ Relational (Comparison) Operators

Used to compare two values — result is always `true` or `false`.

| Operator | Description | Example |
|-----------|-------------|----------|
| == | Equal to | a == b |
| != | Not equal to | a != b |
| > | Greater than | a > b |
| < | Less than | a < b |
| >= | Greater or equal | a >= b |
| <= | Less or equal | a <= b |

Example:

```go
fmt.Println(a == b) // false
fmt.Println(a > b)  // true
```

---

## 3️⃣ Logical Operators

Combine or invert boolean expressions.

| Operator | Description | Example |
|-----------|-------------|----------|
| && | Logical AND | a && b |
| \|\| | Logical OR | a \|\| b |
| ! | Logical NOT | !a |

Example:

```go
x := true
y := false
fmt.Println(x && y) // false
fmt.Println(x || y) // true
fmt.Println(!x)     // false
```

---

## 4️⃣ Assignment Operators

Used to assign values to variables.

| Operator | Description | Example |
|-----------|-------------|----------|
| = | Assign | a = 10 |
| += | Add and assign | a += 5 |
| -= | Subtract and assign | a -= 2 |
| *= | Multiply and assign | a *= 3 |
| /= | Divide and assign | a /= 4 |
| %= | Modulus and assign | a %= 2 |

Example:

```go
count := 10
count += 5 // 15
count *= 2 // 30
```

---

## 5️⃣ Increment and Decrement

Go supports `++` and `--`, but only as statements, not expressions.

✅ Valid:

```go
i := 5
i++ // i = 6
i-- // i = 5
```

❌ Invalid:

```go
j := i++ // compile error
```

---

## 6️⃣ Bitwise Operators (Advanced)

Operate on bits directly.

| Operator | Description | Example |
|-----------|-------------|----------|
| & | AND | a & b |
| \| | OR | a \| b |
| ^ | XOR | a ^ b |
| << | Left shift | a << 1 |
| >> | Right shift | a >> 1 |

Example:

```go
a := 10  // 1010 (binary)
b := 3   // 0011
fmt.Println(a & b)  // 2  (0010)
fmt.Println(a | b)  // 11 (1011)
fmt.Println(a ^ b)  // 9  (1001)
fmt.Println(a << 1) // 20 (10100)
fmt.Println(a >> 1) // 5  (0101)
```

---

## 7️⃣ Operator Precedence

Go evaluates operators based on precedence (highest to lowest):

1. `* / % << >> & &^`
2. `+ - | ^`
3. `== != < <= > >=`
4. `&&`
5. `||`

You can use parentheses `()` to control evaluation order.

Example:

```go
result := 10 + 2*3    // 16
result = (10 + 2) * 3 // 36
```

---

## 🧠 Summary

- Arithmetic → + - * / %
- Relational → == != > < >= <=
- Logical → && || !
- Assignment → = += -= *= /= %=
- Bitwise → & | ^ << >>
- Use parentheses to control operation order.
- `++` and `--` are statements only.

---

## 🚀 Try It Yourself

Run:

```bash
go run ./lessons/03_operators
```

Expected output:

```plaintext
Arithmetic:
13 7 30 3 1

Relational:
false true

Logical:
false true false

Assignment:
count: 30

Bitwise:
2 11 9 20 5
```
