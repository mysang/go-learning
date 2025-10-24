# Lesson 06 — Arrays & Slices in Go

Go provides two powerful data structures for storing collections of elements:

- **Arrays** — fixed-size sequences of elements.
- **Slices** — flexible, dynamically sized views of arrays.

---

## 1️⃣ Arrays

An **array** is a fixed-length sequence of elements of the same type.

Declaration:

```go
var nums [3]int
nums[0] = 10
nums[1] = 20
nums[2] = 30
```

Or declare and initialize in one line:

```go
nums := [3]int{10, 20, 30}
```

If you omit the size, Go infers it automatically:

```go
nums := [...]int{10, 20, 30, 40}
```

Access elements using index:

```go
fmt.Println(nums[0]) // 10
fmt.Println(nums[2]) // 30
```

---

### Properties of Arrays

- **Length is fixed** and part of its type.
  Example: `[3]int` ≠ `[4]int`.
- **Copied by value** — assigning one array to another copies all elements.
- **Zero values** are automatically assigned.

Example:

```go
var arr [3]string
fmt.Println(arr) // ["", "", ""]
```

---

### Iterating Arrays

Use a `for` loop or `range`:

```go
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}

for index, value := range nums {
    fmt.Println(index, value)
}
```

---

## 2️⃣ Slices

A **slice** is a dynamically-sized, flexible view into an array.

Declaration:

```go
var s []int                // nil slice
s = []int{1, 2, 3, 4, 5}   // initialized slice
```

You can also create slices from arrays:

```go
arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4] // elements from index 1 to 3
fmt.Println(slice) // [20 30 40]
```

---

### Length and Capacity

Use built-in functions:

```go
fmt.Println(len(slice)) // number of elements
fmt.Println(cap(slice)) // capacity (from start index to end of array)
```

Example:

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:3]
fmt.Println(len(s)) // 2
fmt.Println(cap(s)) // 4 (from index 1 → end of array)
```

---

### Appending to a Slice

Use the built-in `append()` function:

```go
nums := []int{1, 2, 3}
nums = append(nums, 4, 5)
fmt.Println(nums) // [1 2 3 4 5]
```

When a slice exceeds its capacity, Go automatically allocates a new underlying array.

---

### Copying Slices

Use the `copy()` function to copy elements between slices:

```go
src := []int{1, 2, 3}
dest := make([]int, len(src))
copy(dest, src)
fmt.Println(dest) // [1 2 3]
```

---

### Creating a Slice with `make()`

`make()` creates slices with specified length and capacity.

```go
s := make([]int, 3, 5)
fmt.Println(s)        // [0 0 0]
fmt.Println(len(s))   // 3
fmt.Println(cap(s))   // 5
```

---

## 3️⃣ Common Operations on Slices

### Append Another Slice

```go
a := []int{1, 2}
b := []int{3, 4, 5}
a = append(a, b...) // spread operator
fmt.Println(a)      // [1 2 3 4 5]
```

---

### Remove an Element

```go
nums := []int{10, 20, 30, 40, 50}
index := 2
nums = append(nums[:index], nums[index+1:]...)
fmt.Println(nums) // [10 20 40 50]
```

---

### Slicing Tricks

```go
s := []int{1, 2, 3, 4, 5}
fmt.Println(s[:3]) // [1 2 3]
fmt.Println(s[2:]) // [3 4 5]
fmt.Println(s[:])  // [1 2 3 4 5]
```

---

## 4️⃣ Difference Between Array and Slice

| Feature | Array | Slice |
|----------|--------|--------|
| Size | Fixed | Dynamic |
| Memory | Value copy | Reference to array |
| Flexibility | Limited | Very flexible |
| Typical Use | Known small set of values | Variable-length data |

---

## 🧠 Summary

- Arrays have fixed size, slices are dynamic.
- Slices are views into arrays.
- Use `len()` and `cap()` to check size and capacity.
- Use `append()` to grow slices.
- Use `copy()` to clone data.
- Use `make()` for preallocated slices.

---

## 🚀 Try It Yourself

Run:

```bash
go run ./lessons/06_arrays_slices
```

Expected output (shortened):

```plaintext
    Array: [10 20 30]
    Array copy: [10 20 30]
    Slice from array: [20 30 40]
    len=2 cap=4
    Append: [1 2 3 4 5]
    Copy: [1 2 3]
    Make: [0 0 0] len=3 cap=5
    Append slice: [1 2 3 4 5]
    Remove index 2: [10 20 40 50]
```
