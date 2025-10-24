package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Simple if
	isRaining := true
	if isRaining {
		fmt.Println("Bring an umbrella!")
	}

	// 2. if-else
	temperature := 30
	if temperature > 35 {
		fmt.Println("It's too hot!")
	} else {
		fmt.Println("Weather is fine.")
	}

	// 3. if-else if
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

	// 4. short variable declaration
	if v := math.Pow(2, 3); v < 10 {
		fmt.Println("Small number")
	} else {
		fmt.Println("Big number")
	}

	// 5. switch
	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend is coming!")
	default:
		fmt.Println("Midweek day")
	}

	// 6. switch true
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

	// 7. fallthrough
	num := 1
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
}
