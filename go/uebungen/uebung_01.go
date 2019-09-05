package main
import (
	"fmt"
	"os"
)


type Audi struct {
	Name string
	Power int
}

/* 1.2
	Alle Parameter werden als Array in os.Args übergeben
	Die Elemente im Array sind Strings
	Element an Position 0 ist aber sbsolute Pfad der ausgeführten Binary
*/
func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	/*
		%v:	value in default format
		%T:	syntax representation
		\n:	line break
		%b:	number in base-2 	(binary)
		%x: number in base-16	(hexadecimal)
	*/

	audi := Audi{"A6", 230}
	fmt.Printf("%%v: %v \n", audi)		// %v: {A6 230}
	fmt.Printf("%%v: %v %v \n", audi)	// %v: {A6 230} %!v(MISSING)
	fmt.Printf("%%T: %T\n", audi)		// %T: main.Audi
	fmt.Printf("%%#v: %#v\n", audi)		// %#v: main.Audi{Name:"A6", Power:230}

	fmt.Printf("%%b: %b \n", 6)			// %b: 110
	fmt.Printf("%%x: %x \n", 16)		// %x: 10

	fmt.Printf("%%s: %s \n", "Hello")	// %s: Hello
	fmt.Printf("%%q: %q \n", "Hello")	// %q: "Hello"
}	