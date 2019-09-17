# Grundlagen

## Features

- UTF-8
- Keine Default Parameter
- Kein implizites Casting
- Kein Operator Overloading
- Keine Exceptions
- Keine Makros
- keine Funktionsankündigung

## Eigenschaften

- Stabil / Rückwärtskompatibel
- Einfach typisiert aber nicht typenlos
- wenig Memory Allocation
- Structs und Arrays
- Concurrency
- Sprache zur Systemprogrammierung
- Praktische Sprache
- Relativ nah an der Hardware
- Kontrolle wichtiger als Bequemlichkeit
- Schnell zu übersetzen
- beinhaltet Software-Werkzeuge
- Gegenentwurf zu C++

## APIs und Blöcke

- Grafik
- Kryptographie
- Netzwerk, verteilte Systeme
- Unterstützung vieler File-Formate
- Unterstützung vieler Protokolle

## Open Source

- Open-Source
- Quellcode einsehbar unter [golang.org](https://golang.org/pkg)
- Läuft auf Unix-ähnlichen Systemen, Linux, FreeBSD, OpenBSD, MacOSX, Windows
- In der Regel läuft der Code plattformunabhängig

# Structs

## "main"-Funktion

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}
```

- Schlüsselwort func definiert Funktion
- Funktion main in allen main-Paketen kann von außen aufgerufen werden (Einstiegspunkt)
- Aufrufen der `Println` Funktion aus dem `fmt` Paket

```go
package main
import "fmt"
func main() {
    var i int = 23
    fmt.Println("i=", i)
}
```

```bash
$ go run i.go
i= 23
```

- Variablen in go sind statisch getypt
- Gleichzeitige Deklaration mehrerer Variablen
- `var vorname, nachname string = "Willi", "Hans"`

## Deklaration

- Lokale Typinferenz (Übersetzer erkennt Type automatisch)
  - `var pi = 3.141`
  - Kurzform lokaler Typinferenz `x, y := 42, 23`
- Deklaratiion ohne Initialisierung (Auf Nullwert gesetzt)
  - `var x string`
- Jede Variable hat einen statischen Typ
- Muss zur Übersetzung bekannt sein

| Vorteile                 | Nachteile                             |
| ------------------------ | ------------------------------------- |
| Effiezientere Ausführung | Frühe Fehlererkennung                 |
| Schnellere Übersetzung   | In einigen Fällen mehr Notation nötig |
| Frühere Fehlererkennung  |                                       |

## Basistypen

- bool
- string
- int, int8, ..., int64
- uint uint8, ..., uint64
- uintprt (kann beliebig viele Zeiger enthalten)
- byte
- rune (alias für int32, Unicode Code Point)
- float32, float64
- complex64, complex128

## Operatoren

- Arithmetische Operatoren
  - `-, *, +, /`
  - `% (Modulo, Divisionsrest)`
  - `+ (String Konkatenation)`
- Vergleichsoperatoren
  - `a == b`
  - `a != b`
  - `a < b`
  - `a > b`

## Formatierung

| General |                                                   |
| ------- | ------------------------------------------------- |
| %v      | the value in default format                       |
| %#v     | Go syntax representation of the value             |
| %T      | Go syntax representation of the type of the value |
| %%      | a literal percent sign; consumes no value         |

| Integer |                                                              |
| ------- | ------------------------------------------------------------ |
| %d      | dezimal                                                      |
| %b      | base 2                                                       |
| %o      | base 8                                                       |
| %x      | base 16                                                      |
| %c      | the char represented by the corresponding Unicode code point |

| Floating Point |                                             |
| -------------- | ------------------------------------------- |
| %e             | scientific notation, e.g. -1.234456e+78     |
| %f             | decimal point but no exponent, e.g. 123.456 |

| String and slice |                                                                                 |
| ---------------- | ------------------------------------------------------------------------------- |
| %s               | the uninterpreted bytes of the string or slice                                  |
| %q               | the double quoted string safely escaped with go syntax                          |
| %p               | (Slice only) the address of the 0th element i nbase 16 notation with leading 0x |

| Pointer |                                  |
| ------- | -------------------------------- |
| %p      | base 16 notation with leading 0x |

[Documentation](https://golang.org/pkg/fmt/)

## For-Schleife

- Go hat nur ein Schleifenkonstrukt, die for-Schleife
- Die normale for-Schleife hat drei, durch Semikolon
  - das **Init-Statement** wird vor der ersten Schleifeniteration ausgeführt
  - Der **Condition-Ausdruck** wird vor jeder Iteration ausgewertet
  - Das **Post-Statement** wird am Ende jeder Iteration ausgeführt
- Die drei Statements werden **nicht** von runden Klammern eingefasst, dafür müssen {} zwingend gesetzt werden

```go
for a := 0; a < 10; a++ {
    fmt.Printf("value of a: %d\n", a)
}
```

- Alternativform: Wiederhole, solange Bedingung wahr ist
- Entspricht while-schleife in anderen Sprachen

```go
for a < b {
    a += 1
}
```

- Alternativform: Mit range

```go
a := []int{1,2,3}
for index, value := range a {
    fmt.Printf("index: %d, value: %d\n", index, value)
}
/* Ausgabe
index: 0, value: 1
index: 1, value: 2
index: 2, value: 3
*/
```

## Funktionen

```go
package main
import "fmt"

func f(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println(f(1, 2), f(2, 3))
}
```

Beispiel mit zwei Rückgabewerten

```go
func swap1(x int, y int) (int, int) {
    return y, x
}
x, y := swap1(1,2)
// alternativ

func swap2(x int, y int) (rx int, ry int) {
    rx = y
    ry = x
    return
}
```

- Funktionen können Rückgabetyp haben
- Mehrere Rückgabewerte möglich
- `return` beendet Funktion

## Kommentare

```go
// Bis zum Ende der Zeile

var x /* eingebetteter Kommentar */ int
```

## Sonstiges

```go
f(); g(); h()

// ist gleichwertig zu
f()
g()
h()
```

# Details

| Vorzüge                                                                      | Schwächen                                                                     |
| ---------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| plattformübergreifen                                                         | keine generischen Typen                                                       |
| start vereinfachte Syntax                                                    | nur teilweise objektorientiert                                                |
| automatisches Speichermanagement (Garbage Collection)                        | ausbaufähige Unterstützung von IDEs                                           |
| einheitliche Code-Formatierung                                               | vergleichsweise spärliches Angebot an Bibliotheken und packages               |
| einfacher Importprozess                                                      | mühsamer Umstieg von klassichen, objektorientierten Sprachen wie Java und C++ |
| mehrere Rückgabewerte für Funktionen und Methoden möglich                    | (erst) wenige Tutorials, Experten etc.                                        |
| Nebenläufigkeit                                                              |                                                                               |
| umfangreiche Standardbibliothek (insbesondere für HTTP und Netzwerkaufgaben) |                                                                               |

## Grundlagen

- Go enstand 2007 ursprünglich bei Google, ist aber seit 2009 - weit vor Version 1.0 - ein Open-Source-Projekt auf GitHub
- Gründungsväter sind unter anderem die Bell-Labs/Unix-/C-Berühmtheiten Rob Pike und Ken Thompson

## Kompilieren und Ausführen

- `$ go run hello.go`
- Dieses Kommando kompiliert und lässt das Programm laufen
- Für die Binary wird ein temporäres Verzeichnis erstellt und direkt wieder gelöscht
- das Verzeichnis wird sichtbar wenn man folgenden Befehl eingibt
  - `go run --work hello.go`
- Importierte aber nicht genutzte Pakete werfen Fehler

## Pointer

ohne Pointer

```go
func setX(x int, value int) {
    x = value
}

func main() {
    x := 1
    setX(x, 5)
    fmt.Printf("Value of x is %d", x)
}
// Value of x is 1
```

mit Pointer

```go
func setX(x *int, value int) {
    *x = value
}

func main() {
    x := 1
    setX(&x, 5)
    fmt.Printf("Value if x is %d", x)
}
// Value of x is 5
```

```go
var a int
var b *int

a = 15
b = &a
*b = 20
fmt.Println(a) // 20
```

## struct

- Go ist nicht objektorientiert
- also auch keine Polymorphie oder Overloading
- aber Go hat structs

```go
type Auto struct {
    Name string
    Power int
}
```

### struct initialisieren

- Wie bei Variablen gibt es hier auch verschiedene Varianten

```go
audi := Auto {
    Name: "Audi Quattro",
    Power: 900,
}

// möglich auch
audi := Auto{}

// oder
audi := Auto{name: "Audi Quattro"}
audi.Power = 900
```

### struct mit Funktionen

```go
package main

import "fmt"

type Auto struct {
    Name string
    Power int
}

func (s *Auto) Super() {
    s.Power += 10000
}

func main() {
    audi := &Auto{ "Audi Quattro", 900 }
    audi.Super()
    fmt.Println(audi.Power)
}
```

### new

- `audi := &Auto{}`
- ist gleich zu
- `audi := new(Audi)`

### struct Komposition

```go
package main
import "fmt"

type Person struct {
    Name string
}

type Auto struct {
    *Person
    Power int
}

func (p *Person) Introduce() {
    fmt.Printf("Hi, ich bin %s\n", p.Name)
}

func main() {
    audi := &Auto {
        Person: &Person{ "Willi" },
        Power: 9001,
    }
    audi.Introduce()
}
```

## Arrays

```go
package main
import "fmt"

func main() {
    var a [5]int
    fmt.Println("Init:", a)         // Init: [0 0 0 0 0]
    a[4] = 100
    fmt.Println("Geändert:", a)     // Geändert: [0 0 0 0 100]
    fmt.Println("Hole:", a[4])      // Hole: 100
    fmt.Println("Länge:", len(a))   // Länge: 5

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Explizit: ", b)    // Explizit:  [1 2 3 4 5]

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)       // 2d:  [[0 1 2] [1 2 3]]
}
```

## Slices

- Slices sind Arrays ähnlich
- Ein Slice hat drei Kompenenten
  - Pointer
  - Länge
  - Kapazität

```go
package main
import "fmt"

func main() {
    s := make([]string, 3)  // Schlüsselwort make
    fmt.Println("emp:", s)  // sonst wie ein Array

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

    fmt.Println("set:", s)      // set: [a b c]
    fmt.Println("get:", s[2])   // get: c
    fmt.Println("len:", len(s)) // len: 3

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("append:", s)   // append: [a b c d e f]

    c := make([]string, len(s))

    copy(c, s)
    fmt.Println("copy:", c)         // copy: [a b c d e f]

    l := s[2:5]
    fmt.Println("slice1:", l)       // slice1: [c d e]

    l := s[:5]
    fmt.Println("slice2:", l)       // slice2: [a b c d e]

    l := s[2:]
    fmt.Println("slice3:", l)       // slice3: [c d e f]

    t := []string{ "g", "h", "i" }
    fmt.Println("dcl:", t)          // dcl: [g h i]

    twoD := make([][]int, 3)        // [[] [] []]
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d:", twoDs)       // 2d: [[0] [1 2] [2 3 4]]
}
```

# Coole Dinge

## Go-Workspace

- Der Go-Workspace ist ein zentrales Verzeichnis, welches über die Umgebungsvariable `GOPATH` referenziert wird
- Unterhalb des `GOPATH` finden sich drei Unterverzeichnise
  - `bin/`
  - `pkg/`
  - `src/`
- Go-Programme werden compiliert und zu einem statischen Binary verlinkt
- Die Idee ist eine zentrale ablage aller Go-Projekte
- Gute Idee, aber es gibt Probleme bei unterschiedlichen versionierten Packages
- Ab 1.5 gibt es eine so genannte "Dependency-Management-Technik"
  - und ab 1.11 gibt es die Go-Module

## Test der Installation

- Erstellen Sie den Go-Workspace mit den drei genannten Unterverzeichnissen
- Exportieren Sie das Verzeichnis in der Umgebungsvariable `GOPATH`
- `$ cd $HOME`
- `$ mkdir -p go/bin go/pkg go/src`
- `$ export GOPATH=$HOME/go`
- Danach können Sie in dem `src` Verzeichnis z.B. ein "Hello World" ablegen

## struktur eines Go-Programms

- Der Suffix der Quelldatei ist .go
- Sechs Schlüsselworte charakterisieren ein Go-Programm
  - package
  - import
  - var
  - const
  - type
  - func
- packages sind wichtig, da sie die Funktion von Namesräumen übernehmen

## Variadische Funktionen

```go
func summe(s ...float64) float64 {
    sum := 0.0
    for _, i := range s {
        sum += i
    }
    return sum
}

func main() {
    x := summe(1.5, 3, -4.0)
    fmt.Println(x)
}
```

- Funktionen können eine variable Anzahl von Parametern haben, wobei diese vom gleichen Typ sein müssen
- Innerhalb der Funktion sind die Werte in einem Slice verfügbar
- Auch kann ein Slice genutzt werden, um Parameter an eine Funktion mit variadischer Parameterliste zu übergeben

## Maps

- Eine Map ist ein assoziatives Array
- Bei der Definition einer Map werden Schlüssel und Wertetyp angegeben
- Für den Schlüsseltyp muss es dabei eine Vergleichsfunktion (= bzw. !=) geben, damit entschieden werden kann, ob ein Schlüssel bereits vorhanden ist oder nicht
- Die eingebaute Funktion range() kann genutzt werden, um über alle Elemente einer Map zu iterieren (auch gültig für Arrays, Slices oder String)

```go
type Contact struct {
    name string
}
var contactMap = make(map[int]Contact)

commits := map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
}

fmt.Println(commits["rsc"]) // 3711

for key, value := range commits {
    fmt.Println(key, value)
}
/*
rsc 3711
r 2138
gri 1908
adg 912
*/
```

## Maps mit Funktionen

```go
package main
import "fmt"

func main() {
    mf := make(map[int] func(float64, float64) float64)
    mf[1] = add
    mf[2] = mult
    mf[3] = sub
    mf[4] = div
    for i := 1; i <= 4; i++ {
        fmt.Println(mf[i](2,3))
    }
}

func add(a float64, b float64) float64 {
    return a + b
}

func mult(a float64, b float64) float64 {
    return a * b
}
```

## OR-Mapping

- Es gibt einen Medienbruch zwischen objektorientierter Programmierung und relationalen Datenbanken (object-relational impedance mismatch)
- Objektrelationale Abbildung (engl. object-relational mapping, ORM) kann seine Objekte in einer relationalen Datenbank ablegen
- Dem Programm erscheint die DB dann als objektorientierte DB
- Objektorientierte Programmiersprachen (OOP) kapseln Daten und Verhalten in Objekten, hingegen legen relationale datenbanken Daten in Tabellen ab
- Beide Paradigmen sind grundlegend verschieden

```go
package main
import (
	"fmt"
	"reflect"
)

type programmingLanguages struct {
	name 	string 	`json:"fieldName"`
	rating	int     `xml:"rating"`
	future	bool	`abc:""`
}

func main() {
	a := programmingLanguages {
		name: "Go",
		future: true,
		rating:1,
	}
	ra := reflect.TypeOf(a)

	fmt.Println(ra)                             // main.programmingLanguages
	fmt.Println(ra.Field(0))                    // {name main string json:"fieldName" 0 [0] false}
	fmt.Println(ra.Field(0).Tag.Get("json"))    // fieldName
	fmt.Println(ra.Field(0).Tag.Lookup("json")) // fieldName true
	fmt.Println(ra.Field(2).Tag.Lookup("abc"))  // " " true
}
```

## Interfaces

```
project structure

myproject/
├── hello.go
├── animals/
│   └── animals.go
└── circus/
    └── circus.go

```

```go
package animals

type Dog struct{}

func (a Dog) Speaks() string {
	return "woof"
}

```

```go
package circus

type Speaker interface {
	Speaks() string
}

func Perform(a Speaker) string {
	return a.Speaks()
}
```

```go
package main

import (
	"fmt"
	"./animals"
	"./circus"
)

func main() {
	dog := animals.Dog{}
	fmt.Println(dog.Speaks())           // woof
	fmt.Println(circus.Perform(dog))    // woof
}
```

- Das Interface da definieren wo man es braucht
- Reduziert die Abhängikeit von Komponenten des package `animal`
- So entsteht robuste Software
- [golang-interfaces](https://blog.chewxy.com/2018/03/18/golang-interfaces/)

## Channel

```go
package main

import "fmt"

func multiplyByTo(num int, out chan<- int) {
	result := num * 2
	out <- result
}

func main() {
	n := 3
	var result int
	out := make(chan int)

	go multiplyByTo(n, out)
	result = <-out

	fmt.Println(result) // 6
}
```

IN/OUT

```go
package main

import "fmt"

func multiplyByTo(in <-chan int, out chan<- int) {
	num := <-in
	result := num * 2
	out <- result
}

func main() {
	out := make(chan int)
	in := make(chan int)

	go multiplyByTo(in, out)
	go multiplyByTo(in, out)timo kaesbach

	in <- 1
	in <- 2

	fmt.Println(<-out) // 2
	fmt.Println(<-out) // 4
}
```

- Daten fließen in Richtung des Pfeils
- `ch <- v` Send v to channel ch.
- `v := <-ch` Receive from ch, and assign value to v.

# Internet

## Net

```go
package main
import (
    "net/http"
    "strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```

```bash
$ curl http://localhost:8080/test
Hello test
```

## Ping Pong / Files

```go
package main
import (
    "net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```

- Funktion akzeptiert eingehende Verbindung auf TCP-Port 8080
- Für jede Verbindung wird eine Go Routine gestartet, die den Request an einen Request-Handler delegiert
- Ein Request-Handler ist eine Go-Funktion mit fester Signatur
  - `func kontakte(writer http.ResponseWriter, request *http.Request)`
- Wichtige Felder von http.Request sind URL und Method für den Zugriff auf URL und HTTP-Methode

## Routing

- Routing bezeichnet den Prozess der Bestimmung und Weiterleitung eingehender HTTP-Requests an einen zuständigen Request-Handler
- Die einfachste Form des Routings ist das Registrieren von und Binden eines Requests-Handlers auf ein URL-Muster über die Funktion `http.HandleFunc`
  - `http.HandleFunc("/kontakte", kontakte)`
- Der Aufruf bewirkt, dass der Request der Form `http://localhost:8080/kontakte` an die Funktion `kontakte` delegiert wird

## Request/Response

- Der Request-Handler `kontakte` kombiniert die earbeitung und das Schreiben der antwort in einem Schritt
- Die Methode `Write` schreibt die Byte-Slice in die vom Writer repräsentierte HTTP-Verbindung
- Der Default-Statuscode ist `200 OK` und wird implizit beim Aufruf von Write gesetzt

## Unter Linux

- Der Code kann auch so übersetzt werden:
  - `env GOOS=linux go build main.go`
- Das entstandene Binary kann so direkt auf Linux-Maschinen ausgeführt werden
  - `$ ./main`
  - `$ curl http://localhost:8080/kontakte`
  - `Anton, Berta, Hans`

## Routing mit gorilla/mux

- `http.ListenAndServe(":8080", nil)` dient der Beantwortung von HTTP-Requests
- Die Go.Standardbibliothek net/http enthält den Typ http.ServerMux, der das Interface http.Handler implementiert
- Wird nil übergeben, dann wird das Routing von DefaultServerMux übernommen (dort werden alle Handler implizit registriert)
- Besser geeignet ist das (externe) Package `gorilla/mux`, ein Multiplexer (Router)
  - `go get -u github.com/gorilla/mux`

```go
package main
import (
	"github.com/gorilla/mux"
	"net/http"
)

func kontakte(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Anton, Berta, Hans"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/kontakte", kontakte)
	http.ListenAndServe(":8080", r)
}
```

## Die Idee REST

- Die Idee bei REST ist die Identifizierung von Ressourcen über URIs
- Die "Pflege" der Ressourcen geschieht über HTTP-Befehle
  - `r.HandleFunc("/kontakte/{id:[0..9]+}", getContact).Methods("GET")`
- Hier wird ein Endpunkt für einzelne Kontakte mit numerischer ID definiert

```go
package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var contacts = make(map[int]string)

func getContact(writer http.ResponseWriter, request *http.Request) {
	v := mux.Vars(request)
	id, _ := strconv.Atoi(v["id"])
	writer.Write([]byte(contacts[id]))
}

func main() {
	contacts[0] = "Anton"
	contacts[1] = "Berta"
	contacts[2] = "Ceasar"

	r := mux.NewRouter()
	r.HandleFunc("/contacts/{id}", getContact).Methods("GET")
	http.ListenAndServe(":8080", r)
}
```

# REST

## CURL

- `curl -X GET http://localhost:8080/kontakte/1`
- `curl -X POST`
  - `-H "Content-Type: application/json"`
  - `-d '{"Vorname": "Willi666", "Nachname": "Baltimore"}'`
  - `http://localhost:8080/kontakte`
- `curl -X DELETE http://localhost:8080/kontakte/1`
- `curl -X PUT`
  - `-H "Content-Type: application/json"`
  - `-d '{"Vorname": "Peter", "Nachname": "Petersen"}'`
  - `http://localhost:8080/kontakte/`
