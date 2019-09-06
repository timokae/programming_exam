# Swift Projekt

Swift Projekt besteht aus

## Swift Kommandozeilen-Compiler

## Standard Library

- Datentypen, Protoklle und Funktionen
- Int, Double, ...
- Arrays, Dictionaries, zusammen mit Protokollen und Algorithmen
- Character, Strings und Low Level Primitives
- Selber in Swift geschrieben, unterscheidet sich aber in einigen Punkten von normalen Swift Code

## Core Libraries für High-Levele Funktionen

- Dienen zur Bereitstellung von plattformunabhängigen Funktionen, die der Standard Library übergeordnet sind
- Allgemein benötigte Typen wie URLs, Character Sets und spezialisierte Collections
- Modultests
- Grundlegende Netzwerkfunktionen
- Scheduling und Prozessausführung, einschließlich Threads, Queues und Notifications
- Persistenz, einschließlich Prperty Lists, Archives, JSON Parsing und XML Parsing
- Unterstützung vin Datums, Zeit und Kalenderfunktionen sowie Berechnung
- Abstraktion von bestriebssystemspezifischen Funktionen
- Interaktion mit dem Dateisystem
- Internationalisierung, einschließlich Datum- und Zahlenformatierung sowie sprachenspezifische Ressourcen
- Benutzerpräferenzen
- Corelibs-Libdispatch für die Handhabung von Nebenläufigkeit auf Multicore Hardware
- Corelibx-xctest für grundlegende Test-Infrastruktur

## Dem LLDB-Debugger mit Swift REPL (Read Eval Print Loop)

## Swift Package Manager

- Werkzeug zur Verwaltung und Verteilung von Swift-Code
- Ist in das Swift Build-System integriert
- Ist zuständig für Download, Übersetzen und Linken von Abhängigkeiten
- Ist ab Swift 3.0 und höher verfügbar
- Unter Abhängigkeiten versteht man Module die für die Übersetzung eines Quellcodes erforderlich sind
- Eine Abhängigkeit besteht aus einer relativen oder absoluten URL und einem Satz von Anforderungen bezüglich der benøotigten Paketversion
- Diese Abhängigkeiten sind rekursiv, d.h. ein Paket kann weitere Pakete als Abhängigkeit erfordern

## XCode Playground Support

# Sicherheitsaspekte

- Swift ist von Grund aufentworfen um höhere Sicherheitsanforderungen zu entsprechen als C-basierte Sprachen
- Das Sprachkonzept vermeidet verschiedene Kategorien von unsicherem Code
- Variablen sind vor der ersten Verwendung immer initialisiert
- Arrays und Integer werden auf Bereichsüberläufe geprüft
- Es gibt ein automatisches Speichermanagement
- Swift Objekte können niemals undefiniert (Null) sein
- Der Versuch ein Null-Objekt zu benutzen führt zu einem Compilerfehler
- Dadurch werden viele typische Laufzeitfehler vermieden
- Für Fälle in denen Null-Werte benötigt werden kennt Swift "Optionals"
- Durch dieses Verfahren wird eine sichere Behandlung von Null-Werten erzwungen

# Sprachelemente

## Kommentare

```swift
// Zeilenkommentar
printf("Hallo, Swift") // Zeilenkommentar

/*
  Blockkommentar
*/

/**
  Kommentare mit Markdown Syntax
*/
```

## Semikolons und main()

- keine Semikolon benönigt um Statement Ende zu markieren
- Wird nur benönigt um mehrere Statements innerhalb eienr Zeile zu trennen

```swift
let text = "Hallo"
let text = "Hallo"; print(text)
```

- Code im globalen Gültigkeitsbereich dient als Einstiegspunkt für die Programausführung
- Dadurch ist eine `main()`-Funktion wie in Java oder C unnötig
  - `print("Hallo, Swift")` ist ein vollständiges Programm

## Variablen und Konstanten

- Int
- Double
- Float
- Bool
- Character
- String
- Array
- Set
- Dictionary
- Tupel
- Optional
- seit Swift 4.0 String (wieder) als Collection implementiert

* Swift ist eine streng typisierte Sprache
* Eine spätere Typänderung ist nicht möglich

```swift
var varName : Int     // explizite Typzuweisung
var number = 123.4    // Typ durch Typinferenz
var wert : Double = 3 // explizit mit Initialisierung
var textVar = "Hallo
```

- Variablennamen können aus (fast) allen Unicode-Zeichen bestehen
- Ziffern sind erlaubt wenn die Ziffer nicht das erste Zeichen des Names ist
- Einige SOnderzeichen sind (\$) sind nur als erstes Zeichen verboten
- Leerzeichen, Rechenzeichen und Satzzeichen sind verboten
- Variablen sollen mit einem kleinbuchstaben beginnen
- zusammengesetzte Variablennamen mit CamelCase-Notation

```swift
var wert = 123      // zulässig
var ein Wert = 123  // unzulässiges Zeichen
var ein_Wert = 123  // zulässig
var Zähler = 123    // zulässig
var μ = 1.23        // zulässig
var zähler1 = 123   // zulässig
var 1wert = 123     // unzulässig (Ziffer am Anfang)
var _wert = 123     // zulässig
var wert$ = 123     // zulässig
var $wert = 123     // $ als erstes Zeichen verboten
var h=0.0, w=0.0    // zulässige Mehrfachzuweisung
```

### Konstanten

- Für Kosntanten gelten weitestgehend ide gleichen Regeln wie für Variablen
- Konstanten werden mit dem Schlüsselwort `let` definiert
- Konstantent müssen **nicht** bei der Definition initialisiert werden
- Eine Wertänderung nach der Initialisierung ist naturgemäß nicht möglich

```swift
let dieAntwortAufAlles = 42
let pi : Double = 3.14159
let firmenName = "ACME Corp."
let w,h: Double
```

### Integer-Typen und Grenzwerte von Variablen

- Int, Int8, Int16, Int32, Int64
- UInt, UInt8, UInt16, UInt32, UInt64
- Grenzwerte Int64.min -> Int64.max
- Apple empfiehlt auch für rein positive Werte ausdrücklich die Verwendung von Int

### Zahlendarstellung

- Dezimale Werte ohne Prefix
  - `let dezimalZahl = 789`
- Binäre Werte mit Prefix `0b`
  - `let bin = 0b11`
- Hexadezimale Werte mit Prefix `0x`
  - `let hex = 0x1FE`
- Exponentialdarstellung möglich
  - `0.37e4` -> `0.34*10^4` -> `3700`

### String und Character

- Character speichert einzelne Unicode-Zeichen
- String speichert veränderliche Zeichenketten bestehnt aus einzelnen Zeichen
- Character und String werden durch doppelte Anführungszeichen gekennzeichnet
- Swift verwendet als Default, auch für einzelne Zeichen, den Typ String
- Character und String sind "Value Types" und werden systemintern als struct realisiert
- String sind nicht immutabe, beim hinzufügen wird kein neuer String angelegt

```swift
var s : String    // uninitialisierter String
var s = ""        // leerer String
var s = String()  // auch ein leerer String
var s = "ABC"     // mit ABC initialiserter String
var s = String("ABC") // auch mit ABC initialiserter String
var s = String(255, radix: 16) // Hex-String
var s = "Preis in \u{20ac}" // Preis in €
var s = "Text\nüber\nmehrere\nZeilen"
var s = """
Text
über
mehrere
Zeilen
"""
var s = "AbcDef"
s.uppercased() // ABCDEF
s.lowercased() // abcdef
```

## Optionals

- Swift kennt zur Darstellung von undefinierten Werten für jeden Wert den speziellen Typ `optional`
- Optionals werden mit einem `?` oder `!` als Kennung definiert
- Zur Druchführung von Operationen müssen Optionals in den entsprechenden "einfachen" Datentyp umgewandelt werden
- Diese Umwandlung wird mit dem `!` Operator durchgeführt (unwrapping)
- Undefinierte Optionals haben den Wert `nil`
- Nil markiert die Abwesenheit eines bestimten Typs
- mit einem `?` definerte Optonals müssen explizit in einfache Variablen umgewandelt werden
- Mit einem `!` definierte Optionals werden implizit umgewandelt

## If

```swift
if wert == 1 {
  print("Wert ist 1")
} else if wert == 2{
  print("Wert gleich 2")
} else {
  print("Wert ungleich 1")
}

// Mehrer Bedingungen werden druch Kommata (und-)verknüpft
if x > 0, y > 0 {
  ...
}

// Logische Verknüpfungen mit && und || möglich
```

Das Unwrapping von Optionals kann mit Hilfe einer Verzweigung abgesichert werden

```swift
if optInt == nil {
  print("Nix drin")
} else {
  print(" Der Wert ist \(optInt!)")
}
```

oder in einer kombinierten Variante von `if` und `let`

```swift
if let a = optInt {
  print("Der Wert ist \(a)")
} else {
  print("auch nix")
}
```

Diese Variante wird als Optional Binding bezeichnet

### Nil Coalescing Operator

Die Zuweisung eines Default-Wertes an eine Optional-Variable mit dem Wert `nil` kann über ein `if` realisiert werden

```swift
let anzahlRäder : Int?
let räderAmauto : Int
if anzahlRäder == nil {
  räderAmAuto = 4
} else {
  räderAmAuto = anzahlRäder
}
```

noch einfacher geht es mit dem nil coalescing Operator

```swift
räderAmAuto =  anzahlRäder ?? 4
```

## Tupel

- Tupel sind eine Zusammenfassung mehrerer einzelner werte zu einer zusammengesetzten Einheit
- Im Gegensatz zu anderen zusammengesetzten Datentypen können Tupel nicht erweitert werden
- Die Elemente des Tupels können aus beliebigen, auch unterschiedlichen, Datentypen bestehen

```swift
let kegelParameter = (5, 4.5)
kegelParameter.0 // 5

let kegelParameter = (hoehe: 5, radius: 4.5)
kegelParameter.hoehe // 5
```

- Tupel können auch ohne Initialisierung angelegt werden
  - `var tvProgramm : (Int, String)`
- oder mit Namen
  - `var tvProgramm : (nummer : Int, name : String)`
- die Einzelwerte eines tupels können direkt Variablen zugewiesen werden
  - `let (n1, n2) = tvProgramm`
- nicht benötigte Komponenten werden mit einem Unterstrich ausgelassen
  - `let (_, n2) = tvProgramm`

# Funktionen

- Jede Funktion wird durch den Funktionsnamen und ie Typen der Übergabeparamter und der Rückgabeparameter eindeutig identifiziert
- Funktionen sind ein "first class type, d.h. dass Funktionen auch als Funktionsparameter und Rückgabewerte verwendet werden
- Beim Funktionsaufruf müssen für alle Parameter die entsprechenden Namen angegeben werden, auch in der gleichen Reihenfolge

```swift
func fname(param1: Int, param2: String) -> Int {
  return 0
}

fname(param1: 1, param2: "Hallo")
```

- Es gibt die Möglichkeit interne und externe Parameternamen anzugeben
- Um externe Parameternamen abzuschalten wird dieser mit einem Unterstrich angegeben

```swift
func begruessung(extName name : String, foermlich : Bool) {
  print(name)
}
begruessung(extName: "Name")
```

- Es gibt auch die Möglichkeit Default Werte für Parameter festzulegen
- Wird bei mAufruf der Parameter nicht angegeben wird der Default Wert verwenden

```swift
func begruessung(name: String = "Anon", foermlich : Bool) {
  print(name, foermlich: true)
}
begruessung() // "Anon"
```

## In-Out-Parameter

- Als Standard werden alle Übergabeparameter beim Aufruf einer Funktion "by value" übergeben und sind Konstanten
- Durch Verwendung des Schlüsselworts `inout` können Übergabeparameter "by reference" übergeben werden
- Die Parameter müssen beim Aufruf mit einem `&` gekennzeichnet werden

```swift
func tauschen(a: inout Int, b: inout Int) {
  let temp = a
  a = b
  b = temp
}
tauscheInteger(&a, &b)
```

- Wenn eine Funktion keinen Rückgabewert hat kann die Angabe komplett ausgelassen werden
- Wenn ein Rückgabewert verwendet werden soll, dann muss der Typ der Rückgabe exakt angegeben werden
- Mehrere Rückgabeweerte sind durch die Verwendung von Tupeln möglich
- Auch komplexere Datentypen (Arrays, Objekte, Strukturen usw) sind erlaubt
- Die aufrufende Stelle darf den Rückgabetyp einer Funktion ignorieren **wenn der Aufruf eindeutig ist**
- Funktionen können überladen werden

# Arrays

- Indizes eines Arrays mit n Elementen gehen von 0 bis n-1
- Ein Array kann das gleiche Element beliebig oft enthalten
- Ein Array ist in Swift ein Wertetyp
- Mit dem Schlüsselwert let erzeugte Arrays haben konstante Größe und Inhalt
- Mit dem var erzeugte Arrays können in Größe und Inhalt verändert werden

Erzeugen eines leeren Arrays

```swift
var iField = [Int]()
var iField = Array<Int>()
var iField : [Int] = Array()
```

Erzeugen eines arrays aus Werten

```swift
var iPrim = [2, 3, 4, 5]
let ampelFarben = ["Rot", "Gelb", "Grün"]
```

Erzeugen eines Arrays mit Vorgabewerten

```swift
var pi = [Double](repeating: 3.14, count: 1000)
var pi = Array(repeating: 3.14, count: 1000)
```

## Verwendung von Arrays

- Zugriff auf Elemente
  - `let drittePrimZahl = iPrim[2]
- Element am Ende anhängen
  - `iPrim.append(29)`
- Array mit 2 Elementen anhängen
  - `iPrim += [1,2, 7.4]`
- Drittes Element entfernen
  - `iPrim.remove(at: 2)`
- Anzahl der Elemente
  - `iCount = iPrim.count`
- Inhalt prüfen
  - `if iPrim.contains(7.4) { ... }`

# Schleifen

## For-Schleife

```swift
for i in 1..<10 {
  print(i)
}
```

- Geschlossener Bereich
  - `for i in 10..100` Werte von 10 - 100
- Halboffener Bereich
  - `for i in 5..<20` Werte von 5 - 19
- Laufvariable i muss nicht ausdrücklich mit var initialisiert werden
- Wird die LAufvariable nicht benötigt kann stattdessen ein Unterstrich angegeben werden
  - `for _ in 1..<10 { ... }`
- Durchlaufen aller Elemente eines Arrays/Dictionaries

```swift
let ampelFarben = ["Rot", "Gelb", "Grün"]
for farbe in ampelFarbe {
  print(farbe)
}

let farben["Rot","Gelb","Blau","Schwarz","Weiß"]
for farbe in farben[...2] { print(farbe) } // Bis Blau
for farbe in farben[2...] { print(farbe) } // ab Blau
for farbe in farben[..,2] { print(farbe) } // bis Gelb

for i in 0..farbe.count { print("\(i) \t \(farben[i])") }
for i in farbe.indices { print("\(i) \t \(farben[i])") }
```

## While-Schleife

```swift
var i = 0
while i < 10 {
  i += 1
  print(i)
}
```

## Repeat

```swift
var i = 0
repeat {
  i += 1
  print(i)
} while i < 10
```

# Zufallszahlen

- `random` erzeugt einen Zufallswert in einem vorgegeben Range
  - `Int.random(in: 1..<50)` Zahl von 1 - 49
  - `Float.random(in: -1...1`) Zahl von -1 - 1
- `arc4random_uniform(n)` erzeugt einen zufälligen UInt32-Wert im Bereich von 0 bis n-1
  - Der Zufallszahlengenerator muss nicht initialisiert werden
  - `var lotto = arc4random_uniform(49) + 1`
- `drand48()` erzeugt einen zufälligen Double Wert zwischen 0 (einschließlich) und 1 (ausschließlich)
  - Der Generator muss mit der Anweisung `srand48()` initialisiert werden
  - `srand48(time(nil))`
  - `var dZahl : Double = drand48()` Zahl von 0 - <1

# Verschachtelte Funktionen

- Verschachtelte Funktionen können zur besseren Strukturierung von langem oder unübersichtlichen Code verwendet werden
- Verschachtelte Funktionen dürfen auf die in de umgebenen Funktion definierten Variablen zugreifen
- Eine verschachtelte Funktion darf direkt nur innerhalb der umgebenen Funktion aufgerufen werden
- Verschachtelte Funktionen können als Rückgabewert der umgebenden Funktion zurückgeben und aufgerufen werden

```swift
func compare_qs(a: Int, b: Int) -> Bool {
  func quersumme(z: Int) -> Int {
    if z < 10 {
      return z
    } else {
      return (z % 10) + quersumme(z / 10)
    }
  }

  return quersumme(a)>quersumme(b)
}
```

- Funktionen sind als Übergabeparameter von Funktionen erlaubt
- Funktionen sind als Rückgabewert von Funktionne erlaubt
- Funktionen können einer Variable zugewiesen werden

```swift
var vergleich: (Int, Int)->Bool
vergleich = compare_qs
var resultat = vergleich(123, 456)
```
