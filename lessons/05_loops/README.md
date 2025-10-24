# Lesson 05 — Loops in Go

Go provides only one looping construct — the `for` loop — but it’s flexible enough to handle all loop patterns:

- Traditional `for` loops
- While-style loops
- Infinite loops
- Iterating over collections (`range`)

---

## 1️⃣ Basic For Loop

Standard `for` loop syntax:

```go
for initialization; condition; post {
    // code
}
```

Example:

```go
for i := 1; i <= 5; i++ {
    fmt.Println("Count:", i)
}
```

Explanation:

- **Initialization:** executed once before the loop starts.
- **Condition:** checked before every iteration; if false → loop ends.
- **Post:** executed after every iteration.

Output:

```plaintext
Count: 1
Count: 2
Count: 3
Count: 4
Count: 5
```

---

## 2️⃣ While-Style Loop

Go doesn’t have a `while` keyword, but you can mimic it easily:

```go
i := 1
for i <= 5 {
    fmt.Println("While-style:", i)
    i++
}
```

Equivalent to:

```plaintext
    while (i <= 5) { ... } in other languages.
```

---

## 3️⃣ Infinite Loop

If you omit all conditions, you get an infinite loop.

```go
for {
    fmt.Println("Running forever...")
}
```

⚠️ Warning: Always include a `break` or `return` inside, otherwise it never stops.

Example:

```go
count := 0
for {
    count++
    if count > 3 {
        break
    }
    fmt.Println("Loop", count)
}
```

---

## 4️⃣ Using `break` and `continue`

- `break` → exits the loop immediately.
- `continue` → skips the current iteration and jumps to the next one.

Example:

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue // skip 3
    }
    if i == 5 {
        break // stop loop at 5
    }
    fmt.Println(i)
}
```

Output:

```plaintext
1
2
4
```

---

## 5️⃣ Nested Loops

You can place one loop inside another.

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 2; j++ {
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

Output:

```plaintext
i=1, j=1
i=1, j=2
i=2, j=1
i=2, j=2
i=3, j=1
i=3, j=2
```

---

## 6️⃣ Range Loop (Iterating Collections)

The `for range` form iterates over elements in an array, slice, map, or string.

### a) Slice or Array

```go
nums := []int{10, 20, 30}
for index, value := range nums {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

If you only need values:

```go
for _, v := range nums {
    fmt.Println(v)
}
```

---

### b) Map

```go
scores := map[string]int{"Alice": 90, "Bob": 80, "Charlie": 95}
for name, score := range scores {
    fmt.Printf("%s: %d\n", name, score)
}
```

---

### c) String

```go
for index, char := range "GoLang" {
    fmt.Printf("%d -> %c\n", index, char)
}
```

💡 Note: Go treats strings as sequences of runes (UTF-8 characters).

---

## 7️⃣ Loop Labels (Advanced)

You can label loops and control which one to break/continue.

Example:

```go
outer:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if i*j > 4 {
            break outer
        }
        fmt.Printf("%d * %d = %d\n", i, j, i*j)
    }
}
```

Output stops when `i*j > 4`.

---

## 🧠 Summary

- Go has only one loop construct: `for`
- Supports all loop types (for, while, infinite)
- Use `break` to exit early, `continue` to skip
- Use `range` to iterate slices, maps, and strings
- Use loop labels for advanced control flow

---

## 🚀 Try It Yourself

Run:

```bash
go run ./lessons/05_loops
```

Expected output (shortened):

```plaintext
    Count: 1
    Count: 2
    Count: 3
    Count: 4
    Count: 5
    While-style: 1
    While-style: 2
    While-style: 3
    While-style: 4
    While-style: 5
    Loop 1
    Loop 2
    Loop 3
    1
    2
    4
    i=1, j=1
    i=1, j=2
    ...
    Index: 0, Value: 10
    Index: 1, Value: 20
    Index: 2, Value: 30
    Alice: 90
    Bob: 80
    Charlie: 95
    0 -> G
    1 -> o
    2 -> L
    ...
    1 * 1 = 1
    1 * 2 = 2
    1 * 3 = 3
    2 * 1 = 2
    2 * 2 = 4
```
