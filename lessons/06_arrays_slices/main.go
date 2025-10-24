package main

import "fmt"

func main() {
	// Arrays
	nums := [3]int{10, 20, 30}
	fmt.Println("Array:", nums)

	// Copy array
	arrCopy := nums
	fmt.Println("Array copy:", arrCopy)

	// Slice from array
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	fmt.Println("Slice from array:", slice)
	fmt.Println("len=", len(slice), "cap=", cap(slice))

	// Append to slice
	numsSlice := []int{1, 2, 3}
	numsSlice = append(numsSlice, 4, 5)
	fmt.Println("Append:", numsSlice)

	// Copy slice
	src := []int{1, 2, 3}
	dest := make([]int, len(src))
	copy(dest, src)
	fmt.Println("Copy:", dest)

	// Make slice
	s := make([]int, 3, 5)
	fmt.Println("Make:", s, "len=", len(s), "cap=", cap(s))

	// Append another slice
	a := []int{1, 2}
	b := []int{3, 4, 5}
	a = append(a, b...)
	fmt.Println("Append slice:", a)

	// Remove an element
	nums2 := []int{10, 20, 30, 40, 50}
	index := 2
	nums2 = append(nums2[:index], nums2[index+1:]...)
	fmt.Println("Remove index 2:", nums2)
}
