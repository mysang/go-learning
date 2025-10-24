package main

import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println("Arithmetic:")
	fmt.Println(a+b, a-b, a*b, a/b, a%b)

	fmt.Println("\nRelational:")
	fmt.Println(a == b, a > b)

	x := true
	y := false
	fmt.Println("\nLogical:")
	fmt.Println(x && y, x || y, !x)

	count := 10
	count += 5
	count *= 2
	fmt.Println("\nAssignment:")
	fmt.Println("count:", count)

	fmt.Println("\nBitwise:")
	fmt.Println(a&b, a|b, a^b, a<<1, a>>1)
}
