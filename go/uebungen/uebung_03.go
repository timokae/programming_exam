package main

import (
	"fmt"
	"reflect"
	"time"
)

type Student struct {
	Name	string	`json:"name" xml:"Name"`
	MatrNr	int		`json:"matrNr" xml:"MatrNr"`
}

func (a Student) SayMyName() {
	fmt.Printf("I am the %s!\n", a.Name)
}

type Speaker interface {
	SayMyName()
}

func Perform(a Speaker) {
	a.SayMyName()
}

func main() {
	// 1.2
	fmt.Println(add(1, 2))
	fmt.Println(mul(3, 2))
	fmt.Println(div(6, 2))
	fmt.Println(sub(4, 2))
	fmt.Println(quad(2))

	// 1.3
	stud1 := Student{ "Anton", 123 }
	stud2 := Student{ "Berta", 456 }
	stud3 := Student{ "Ceasar", 789 }

	rStud1 := reflect.TypeOf(stud1)
	rStud2 := reflect.TypeOf(stud2)
	rStud3 := reflect.TypeOf(stud3)
	
	fmt.Println(rStud1.Field(0).Tag.Get("json"))
	fmt.Println(rStud2.Field(1).Tag.Get("xml"))
	fmt.Println(rStud3.Field(0).Tag)

	// 1.4
	Perform(stud1)

	// 1.5
	n := 10

	start := time.Now()
	resultIterative := fibIterative(n)
	durationIterative := time.Since(start)
	
	chanResult := make(chan int)
	start = time.Now()
	go fibRecursive(n, chanResult)
	resultRecursive := <-chanResult
	durationRecursive := time.Since(start)

	fmt.Printf("Fibonacci for n=%d\n", n)
	fmt.Printf("Iteraive: %d in %dns\n", resultIterative, durationIterative)
	fmt.Printf("Recursive: %d in %dns\n", resultRecursive, durationRecursive)
}

// 1.2
func add(a float64, b float64) float64 {
	return a + b
}

func mul(a float64, b float64) float64 {
	return a * b
}

func div(a float64, b float64) float64 {
	if b != 0 {
		return a / b
	} else {
		panic("zero division")
	}
}

func sub(a float64, b float64) float64 {
	return a - b
}

func quad(a float64) float64 {
	return a * a
}

// 1.5
func fibIterative(n int) int {
	x := 0
	y := 1
	z := 1
	for i := 0; i < n; i++ {
		x = y
		y = z
		z = x + y
	}

	return x
}

func fibRecursive(num int, out chan<- int) {
	if num == 1 || num == 0 {
		out <- num
		return
	}

	chanX := make(chan int)
	chanY := make(chan int)

	go fibRecursive(num - 1, chanX)
	go fibRecursive(num - 2, chanY)

	numX := <-chanX
	numY := <-chanY

	out <- (numX + numY)
}