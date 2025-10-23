package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s Student) PrintInfo() {
	fmt.Printf("Name: %s\nAge: %d\nScore: %d\n", s.Name, s.Age, s.Score)
}

func (s *Student) UpgradeScore(delta int) {
	s.Score += delta
}

func main() {
	student := Student{Name: "Sang", Age: 35, Score: 75}
	student.PrintInfo()
	fmt.Println("Increasing score by 10...")
	student.UpgradeScore(10)
	student.PrintInfo()
}
