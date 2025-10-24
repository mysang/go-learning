package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	// 1. Create a map
	fruits := map[string]int{
		"apple":  5,
		"banana": 7,
		"orange": 3,
	}

	// 2. Access and update
	fruits["apple"] = 10
	fmt.Println("apple:", fruits["apple"])
	fmt.Println("grape:", fruits["grape"]) // key not found -> 0

	// 3. Check existence
	value, exists := fruits["grape"]
	if exists {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found!")
	}

	// 4. Delete
	delete(fruits, "banana")

	// 5. Iterate
	for key, value := range fruits {
		fmt.Printf("%s -> %d\n", key, value)
	}

	// 6. Length
	fmt.Println("Total:", len(fruits))

	// 7. Map of structs
	students := map[string]Student{
		"sang": {Name: "Sang", Score: 90},
		"long": {Name: "Long", Score: 80},
	}
	fmt.Println("Sang's score:", students["sang"].Score)

	// 8. Map of maps
	data := map[string]map[string]int{
		"2025": {"Jan": 100, "Feb": 200},
		"2026": {"Jan": 150, "Feb": 250},
	}
	fmt.Println("Data 2025 Jan:", data["2025"]["Jan"])
}
