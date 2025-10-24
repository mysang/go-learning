# Lesson 07 — Maps in Go

A **map** in Go is an unordered collection of key-value pairs.
Maps allow you to associate unique keys with specific values and retrieve data efficiently.

---

## 1️⃣ Declaring a Map

There are several ways to create maps:

### a) Using `make()`

```go
m := make(map[string]int)
m["apple"] = 5
m["banana"] = 7
m["orange"] = 3
```

### b) Using a Map Literal

```go
fruits := map[string]int{
    "apple":  5,
    "banana": 7,
    "orange": 3,
}
```

Both declarations create a map with string keys and integer values.

---

## 2️⃣ Accessing Values

```go
fmt.Println(fruits["apple"])  // 5
fmt.Println(fruits["grape"])  // 0 (key not found)
```

💡 Note:
If a key does not exist, Go returns the **zero value** of the value type (e.g. 0 for int, "" for string, false for bool).

---

## 3️⃣ Checking if a Key Exists

Use the two-value form:

```go
value, exists := fruits["grape"]
if exists {
    fmt.Println("Found:", value)
} else {
    fmt.Println("Not found!")
}
```

Output:
    Not found!

---

## 4️⃣ Updating and Deleting Keys

### Update an existing value

```go
fruits["apple"] = 10
```

### Delete a key

```go
delete(fruits, "banana")
```

💡 If you delete a non-existing key, nothing happens (no panic).

---

## 5️⃣ Iterating Over a Map

Use `for range`:

```go
for key, value := range fruits {
    fmt.Printf("%s -> %d\n", key, value)
}
```

Order is **not guaranteed** — maps in Go are **unordered**.

---

## 6️⃣ Map Length

Use `len()` to count entries:

```go
fmt.Println("Total:", len(fruits))
```

---

## 7️⃣ Nil Maps

A map declared without initialization is `nil`:

```go
var users map[string]int
fmt.Println(users == nil) // true
```

Accessing or writing to a nil map causes a panic:

```go
users["John"] = 5 // ❌ runtime error
```

✅ Always initialize maps before using:

```go
users = make(map[string]int)
```

---

## 8️⃣ Map of Structs

You can store complex data types as values:

```go
type Student struct {
    Name  string
    Score int
}

students := map[string]Student{
    "sang":  {Name: "Sang", Score: 90},
    "long":  {Name: "Long", Score: 80},
}

fmt.Println(students["sang"].Score) // 90
```

---

## 9️⃣ Map of Maps

Maps can also hold nested maps:

```go
data := map[string]map[string]int{
    "2025": {"Jan": 100, "Feb": 200},
    "2026": {"Jan": 150, "Feb": 250},
}

fmt.Println(data["2025"]["Jan"]) // 100
```

---

## 1️⃣0️⃣ Comparing Maps

Maps **cannot** be compared directly (except with `nil`):

```go
a := map[string]int{"x": 1}
b := map[string]int{"x": 1}
fmt.Println(a == b) // ❌ compile error
```

You must compare them manually or using reflection.

---

## 🧠 Summary

- Maps store key-value pairs.
- Access values with `map[key]`.
- Use two-value form `value, exists := map[key]` to check existence.
- Use `delete(map, key)` to remove an entry.
- `len(map)` gives the number of entries.
- Maps are unordered and cannot be compared directly.
- Always initialize maps before use.

---

## 🚀 Try It Yourself

Run:

```bash
go run ./lessons/07_maps
```

Expected output (order may vary):

```plaintext
apple -> 10
orange -> 3
Found: 10
Not found!
Total: 2
Sang's score: 90
Data 2025 Jan: 100
```
