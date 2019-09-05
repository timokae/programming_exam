package main

import (
	"fmt"
)

type Student struct {
	Name 	string
	MatrNr 	int
}

func findMax(a *int, b *int) int {
	if *a > *b {
		return *a
	} else {
		return *b
	}
}

func main() {
	// 1.1 Variablendeklaration
	var a int = 1
	var b = 2
	c := 3
	fmt.Println(a, b, c)
	
	// 1.2 Pointer
	maxNum := findMax(&c, &b)
	fmt.Printf("findMax(%d, %d) = %d\n", c, b, maxNum)

	/* 1.3 Structs
		new: erzeut neues Objekt und gibt Adresse zurück
		&: gibt Adresse von vorhandenem Objekt zurück
	*/
	stud1 := Student{ "Anton", 123 }
	stud2 := Student{ "Berta", 456 }
	stud3 := Student{ "Ceasar", 789 }
	studArray := []Student{ stud1, stud2, stud3 }
	for _, s := range studArray {
		fmt.Println(s)
	}

	// 1.4 Arrays
	var numbers [10]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}
	min := numbers[0]
	max := numbers[0]
	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}

		sum += numbers[i]
	}
	avg := float64(sum) / float64(len(numbers))
	fmt.Printf("min: %d | max: %d | avg: %f\n", min, max, avg)

	// 1.5 Slice
	var months = make ([]string, 12, 12)
	months[0]  = "January"
	months[1]  = "February"
	months[2]  = "March"
	months[3]  = "April"
	months[4]  = "May"
	months[5]  = "June"
	months[6]  = "July"
	months[7]  = "August"
	months[8]  = "September"
	months[9]  = "October"
	months[10] = "November"
	months[11] = "December"

	var q1 = months[0:3]
	fmt.Println(q1)

	var first_half = months[:7]
	fmt.Println(first_half)

	var second_half = months[7:]
	fmt.Println(second_half)

	
}