package main

import "fmt"

func main() {
	// 1. Basic for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Count:", i)
	}

	// 2. While-style
	i := 1
	for i <= 5 {
		fmt.Println("While-style:", i)
		i++
	}

	// 3. Infinite loop (with break)
	count := 0
	for {
		count++
		if count > 3 {
			break
		}
		fmt.Println("Loop", count)
	}

	// 4. Break & Continue
	for j := 1; j <= 5; j++ {
		if j == 3 {
			continue
		}
		if j == 5 {
			break
		}
		fmt.Println(j)
	}

	// 5. Nested loops
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 2; b++ {
			fmt.Printf("i=%d, j=%d\n", a, b)
		}
	}

	// 6. Range over slice
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Range over map
	scores := map[string]int{"Alice": 90, "Bob": 80, "Charlie": 95}
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// Range over string
	for index, char := range "GoLang" {
		fmt.Printf("%d -> %c\n", index, char)
	}

	// 7. Loop labels
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break outer
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}
