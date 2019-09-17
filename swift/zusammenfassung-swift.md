# Swift Projekt

Swift Projekt besteht aus

## Swift Kommandozeilen-Compiler

## Standard Library

- Datentypen, Protokolle und Funktionen
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
- Einige Sonderzeichen sind (\$) nur als erstes Zeichen verboten
- Leerzeichen, Rechenzeichen und Satzzeichen sind verboten
- Variablen sollen mit einem kleinbuchstaben beginnen
- zusammengesetzte Variablennamen mit CamelCase-Notation

```swift
var wert = 123      // zulässig
var ein Wert = 123  // unzulässiges Zeichen
var ein_Wert = 123  // zulässig
var Zähler = 123    // zulässig
var μ = 1.23        // zulässig
var zähler1 = 123   // zulässig
var 1wert = 123     // unzulässig (Ziffer am Anfang)
var _wert = 123     // zulässig
var wert$ = 123     // zulässig
var $wert = 123     // $ als erstes Zeichen verboten
var h=0.0, w=0.0    // zulässige Mehrfachzuweisung
```

### Konstanten

- Für Konstanten gelten weitestgehend die gleichen Regeln wie für Variablen
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
- String speichert veränderliche Zeichenketten bestehend aus einzelnen Zeichen
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

- Jede Funktion wird durch den Funktionsnamen und die Typen der Übergabe- und Rückgabeparameter eindeutig identifiziert
- Funktionen sind ein "first class type", d.h. dass Funktionen auch als Funktionsparameter und Rückgabewerte verwendet werden
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
- Mehrere Rückgabewerte sind durch die Verwendung von Tupeln möglich
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
var iPrim = [2, 3, 5, 7]
let ampelFarben = ["Rot", "Gelb", "Grün"]
```

Erzeugen eines Arrays mit Vorgabewerten

```swift
var pi = [Double](repeating: 3.14, count: 1000)
var pi = Array(repeating: 3.14, count: 1000)
```

## Verwendung von Arrays

- Zugriff auf Elemente
  - `let drittePrimZahl = iPrim[2]`
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
  - `for i in 10...100` Werte von 10 - 100
- Halboffener Bereich
  - `for i in 5..<20` Werte von 5 - 19
- Laufvariable i muss nicht ausdrücklich mit `var` initialisiert werden
- Wird die Laufvariable nicht benötigt kann stattdessen ein Unterstrich angegeben werden
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
- Verschachtelte Funktionen dürfen auf die in der umgebenen Funktion definierten Variablen zugreifen
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

# Closures

- Auch anonyme Funktion oder Lambda-Ausdruck genannt
- Wird nicht über ihren Namen, sondern nur über Verweise angesprochen
- Kann auf Informationen zugreifen, die sich in dem Kontext befinden in dem die Closure erstellt wurde
  - `{ (Parameter) -> Rückgabetyp in Code`
  - `{ (x: Int) -> Int in return x + 1}`
- Bei Funktionen die nur aus einer Anweisung bestehen darf das Schlüsselwort `return` weggelassen werden
- Der (oder die) Typen der Übergabeparameter dürfen, wenn die Typen aus dem Kontext erkennbar sind, ebenfalls entfallen

```swift
var cosfunc: (Double)->Double
cosfunc = { (c: Double)->Double in return cos(c/10.0) }
cosfunc = { (c)->Double in cos(c/10.0) }
var resultat = cosfunc(25)
```

- Wenn auch der Rückgabetyp aus dem Kontext ersichtlich ist kann die gesamte Closure-Definition auf die eigentliche Funktion reduziert werden
- Die dann nicht mehr explizit angegebenen Übergabeparameter bekommen automatisch die Namen $0, $1, u.s.w.

```swift
cosfunc = { (c: Double)->Double in return cos(c/10.0) }
cosfunc = { cos($0/10.0) }

var vergleich: (Int, Int)->Bool
vergleich = { (a: Int, b: Int)->Bool in return a > b }
vergleich = { (a, b)->Bool in a > b }
vergleich = { $0 > $1 }
```

- Bei Funktionen die eine Closure als letzten (oder einzigen) Parameter erwarten darf die Closure auch außerhalb der Parameter-Klammern, also hinter der Funktion stehen (trailing closure)

```swift
var zahlen = Array(1...6)
zahlen.map({ print($0) })
zahlen.map(){ print($0) }
```

# Collections

- Dictionaries speichern Key-Value-Paare
  - Die Wertetypen von Key und Value sind unabhängig voneinander
  - Der Schlüsseltyp muss das Protokoll Hashable implementieren
- Sets sind ungeordnete Datensammlungen ohne Doppelgänger
  - Der Datentyp muss das Protokol Hashable implementieren
- Die Reihenfolge der Elemente in Dictionaries und Sets sind undefiniert

```
Array             Set                 Dictionary
--------------    ---------------     -------
| 0 | Eggs   |    |  Rock       |     | YYZ | ----------> Toronto
|---+--------|    |             |     -------
| 1 | Milk   |    |       Jazz  |     | DUB | \   /-----> London
|---+--------|    |             |     -------  \ /
| 2 | Flour  |    |    Classical|     | LHR |  / \------> Dublin
|--+---------|    |             |     -------
| 3 | Apples |    |             |
|---+--------|    | Hip Hop     |
| 4 | Milk   |    ---------------
--------------
```

## Dictionaries

```swift
var elemente = ["AL": "Aluminium", "AS": "Arsen", "B": "Bor"]

elemente["PB"] = "Blau"   // Element hinzufügen
elemente["PB"] = "Blei"   // Element verändern
elemente.updateValue("Blei", forKey: "PB") // Element verändern
elemente["B"] = nil       // Element löschen
elemente.removeValue(forKey: "B") // Element löschen
elemente.removeAll()      // Dictionarie löschen

print(elemente["AS"]!)    // Rückgabewert ist ein Optional!

var sammlung = [Int: String]()    // leeres Dictionary anlegen
var sammlung: [Int:String] = [:]  // leeres Dictionary anlegen

for (key, value) in elemente {
  print("Das Symbol \(value) ist \(key)")
}
for key in elemente.keys{ print(key) }
for value in elemente.values{ print(name) }
let chemie = [String](elemente.values)  // Umwandlung in Array
```

## Sets

```swift
var buchstaben = Set<Character>       // leeres Set anlegen
let ampel: Set = ["Rot", "Grün", "Gelb"]

buchstaben.insert("A")                // element zufügen
buchstaben = []                       // Set leeren
let entfernt = buchstaben.remove("c") // Element entfernen
buchstaben.removeAll()                // Set leeren

if buchstaben,contains("A") { ... }   // Element prüfen
for b in buchstaben { print(b) }      // Elemente ausgeben
let s = buchstaben.sorted()           // Sortiertes Array aus Set

let schnitt = a.intersection(b)       // Schnitt von a und b
let vereinigung = a.union(b)          // Vereinigung a und b
let symDifferenz = a.symmetricDifference(b) // (Vereinigung a und b) - (Schnitt a und b)
let differenz = a.substracting(b)     // a - b
```

# Klassen und Strukturen

| Gemeinsamkeiten                                    | Unterschiede                                                   |
| -------------------------------------------------- | -------------------------------------------------------------- |
| Gekapselte Eigenschaften zur Speicherung von Daten | Strukturen sind Wertetypen                                     |
| Methoden zur Bereitstellung von Funktionalitäten   | Klassen sind Referenztypen                                     |
| Initialisierbar                                    | Klassen können vererbt werden                                  |
| Berechnete Eigenschaften                           | Klassen können deinitialisiert werden                          |
| Erweiterbarkeit                                    | Referenzzählung erlaubt mehr als eine Referent auf eine Klasse |

```swift
struct EineStruktur {
  // Definition der Struktur
}
let struktur = EineStruktur()

class EineKlasse {
  // Definition der Klasse
}
```

Klassen und Strukturen können Properties enthalten

- Alle Eigenschaften einer Klasse/Struktur müssen initialisiert werden
- Die Initialisierung kann sofort bei der Definition oder später in einer Init-Funktion erfolgen
- Swift unterscheidet zwischen _stored properties_ und _computed properties_
- Auf die Eigenschaften einer Klasse kann in der Punkt-Schreibweise zugegriffen werden
- `open` erlaubt Zugriff und Vererbung aus anderen Modulen
- `public` erlaubt Zugriff aus fremden Modulen, Vererbung aber nur im definierten Modul
- `internal` erlaubt Zugriff und Vererbung nur aus dem definierenden Modul
- `fileprivate` erlaubt Zugriff und Vererbung nur aus der definierrenden Datei
- `prvate` beschränkt den Zugriff auf die definierende Datei
- Public ist die am höchsten (am wenigsten einschränkende) Schutzstufe, private ist die niedrigste (restriktivste) Schutzstufe
- Zugriffsbeschränkungen in Swift unterscheiden sich von ähnlichen Konzepten in anderen Sprachen dadurch, dass hier auf Dateiebene gearbeitet wird
- Ohne Angabe wird automatisch die Stufe internal vergeben

```swift
struct Fahrzeug {
  var anzahlPassagiere = 0
  var anzahlRäder = 0
}
let einFahrzeug = Fahrzeug()
einFahrzeug.anzahlPassagiere = 2
print("Fahrzeug hat \(einFahrzeug.anzahlRäder) Räder")
```

- _computed_ properties können einen getter und einen (otionalen) setter zum Zugriff definieren
- _computed properties_ ihne setter haben die Eigenschaft read-only

```swift
var name: Typ {
  get { return ... }
  set { ... } // neuer Wert als newValue
}

struct Geschwindigkeit {
    var speed: Double = 0
    init(kmh: Double) { speed = kmh/3.6 }
    init(mph: Double) { speed = mph/2.237 }
    var kmh: Double {
        get { return speed * 3.6 }
        set { speed = newValue / 3.6 }
    }
    var mph: Double {
        get { return speed * 2.237 }
        set { speed = newValue / 2.237 }
    }
}

var v0 = Geschwindigkeit()
v0.mph = 100.0
print("Geschwindigkeit in Km/h: \(v0.kmh)")
print("Geschwindigkeit in MPH: \(v0.mph)")
```

- Die init-Funktion entspricht dem Konstruktor in C/Java
- Swift erzeugt als Default eine init-Methode ohne Parameterliste
- Bei Strukturen wird zusätzlich ein default-Initialisierer mit allen Eigenschaften der Klasse in der Reihenfolge des Auftretens erzeugt
- Wenn eine eigene init-Methode programmiert wird entfallen die default Methoden
- Es können beliebig viele init-Methoden programmiert werden, die sich durch eine eindeutige Parameterliste unterscheiden müssen
- Eine init-Methode beginnt mit dem Schlüsselwort `init` und hat keinen Rückgabewert
- `deinit` entspricht dem Destruktor aus C und Java

  - wird unmittelbar vor dem Entfernen der Instanz aus dem Speicher ausgeführt
  - Strukturen haben keine deinit-Methode

- observers sind spezielle Methoden um auf die Veränderung von eigenschaften einer Klasse zu reagieren
  - willSet: unmittelbar vor dem Verändern einer Eigenschaft, neuer Wert als `newValue`
  - didSet: unmittelbar nach dem Verändern einer Eigenschaft, letzter Wert als `oldValue`

```swift
struct Geschwindigkeit {
  var speed: Double = 0 {
    willSet { "new Speed = \(newValue)" }
    didSet { "old Speed = \(oldValue)" }
  }
}
```

- Für die Definition eines Subscript wird das Schlüsselwort `subscript` verwendet
- Werte werden mit get- und set-Methoden gelesen und verändert
- Eine Klasse oder Struktur darf auch mehrere Subscripts mit verschiedenen, auch "Mehrdimensionalen" Parametern haben

```swift
struct EineStruktur {
  var a = 1
  subscript(index: Int) -> Int {
    get { return index + a}
    set(newValue) { a = newValue }
  }
}
```

```swift
struct Fibonacci {
    subscript(index: Int)->Int {
        get { return fib(i: index) }
    }
    func fib(i: Int)->Int {
        if i < 2 {
            return i
        } else {
            return fib(i: i-1) + fib(i: i-2)
        }
    }
}


var f = Fibonacci()
print("Fibonacci(5) = \(f[5])")
```

## Subclass

- Neuimplementierung einer Methode muss mit `override` gekennzeichnet werden
- DIe entsprechende Methode der Elternklasse kann mit dem Prefix `super` aufgerufen werden
- Ein überschriebener subscript kann die Implementierung der Elternklasse mit der Schreibweise `super[index]` aufrufen

```swift
class Kind: Vater {
  var istPleite = false

  override func ausgeben(betrag: Double) -> Double {
    super.ausgeben(betrag)
    if geld <= 0 {
      istPleite = true
    }
    return geld
  }
}
```

# Protokolle

- auch Interface genannt
- Ein Protokoll definiert eine Vorlage (oder Muster) von Eigenschaften, Methoden und Anforderungen
  Eine Klasse oder eine Struktur kann die Anforderungen eines Protokolls erfüllen
  In diesem Fall spricht man davon, dass der Typ das Protokoll einhält oder befolgt
- Ein Potokoll ist ein Datentyp ohne konkrete Implementierung

```swift
protocol MeinProtokoll {
  ...
}

class MeineKlasse: MeinProtokoll, EinProtokoll {
  ...
}

// Soll von Klasse Kind erben
// Soll Protokoll Verwander befolgen
class Bruder: Kind, Verwander {

}
```

- Ein Protokoll kann erfodern, dass ein Typ der dieses Potokoll befolgt eine EIgenschaft mit einem vorgegebenen Namen und Datentyp enthält
  - egal ob Variable oder Computed Property
- Bei der Festlegung wird zusätzlich bestimmt, ob die Eigenschaft nur lesbar oder auch veränderbar ist
- Ein Protokoll kann erfordern, dass ein Typ der dieses Protokoll befolgt eine EMthode mit einem vorgegebenen Namen, Übergabe und Rückgabeparametern enthält
- Die Definition im Protokoll entspricht der Methodendeklaration ohne den Methodenrumpf mit seinen geschweiften Klammern

```swift
protocol MeinProtokoll {
  var nurLesbar: Int { get }
  var kannGeändertWerden: Int { get set }
  func beschleinigen(wert: Int) -> Int
}
```

- Beispiele für Protokolle
  - Numeric, SignedNumeric
  - Comparable, Hashable
  - Collection, OptionSet
  - ...

```swift
struct Auto: Vehicle {
    init(name: String, kwh: Int) {
        self.name = name
        self.kwh = kwh
    }

    var name: String
    var kwh: Int

    var horsepower : Int {
        get { return Int((Double(self.kwh) * 1.35962).rounded()) }
    }
}

protocol Vehicle {
    var horsepower: Int { get }
}

func printHorsepower(v: Vehicle) {
    print(v.horsepower)
}

var audi = Auto(name: "Audi", kwh: 10)
print(audi.horsepower)
printHorsepower(v: audi)
```

# Generics

- Mit Generics können unnötige Wiederholugen im Code verhindert werden
- An der Stelle des Datentyps wird bei Generics ein Platzhalter eingesetzt

```swift
// ohne Generic
func printToConsole(i: Int) {
    print(i)
}
printToConsole(i: 10)

// mit Generics
func printToConsole<T>(i: T) {
    print(i)
}
printToConsole(i: 10)
printToConsole(i: "Hallo")

```

# iOS Apps

## Autplayout - Size Classes

- Size Classes reagieren auf die Kriterien
  - Display Höhe (Height)
  - Display Breite (Width)
- Für jedes dieser Kriterien können die Gruppen
  - Regular
  - Compact
  - und Any unterschieden werden
- Ein Gerät kann mit seiner Ausrichtung über diese Kriterien eingeordnet werden

| Handy         | Portait | Landscape |
| ------------- | ------- | --------- |
| iPhone 4s     | wC, hR  | wC, hC    |
| iPhone Xs     | wC, hR  | wC, hC    |
| iPhone Xs Max | wC, hR  | wR, hC    |
| iPad Pro      | wR, hR  | wR, hR    |

- In Abhängigkeit von der vorliegenden SizeClass können die Layouteigenschaften einer App angepasst werden
- Anpassung von Abständen, Ausrichtung und Abmessung von Objekten
- Anpassung von Farbe, Schriften und Sichtbarkeit
- Über eine beim Wechseln der SizeClass ausgeführten Methode (`traitCollectionDidChange`) können beliebige Anpassugnen vorgenommen werden
- Anpassungen im Storyboard über den Button _Vary for Traits_ oder Einzelanpassugen über `+` in den Eigenschaftsdialogen
- **Empfohlene Vorgensweise**
  - Zuert werden Layoutregeln für den Normalfall Any/Any festgelegt
  - Danach werden Regeln für Sonderfälle ergänzt

## Multi-View Apps

- Die Transition zwischen den Views erfolgt unter Kontrolle der App
- Für den Datenaustausch zwischen verschiedenen Views sind spezielle Methoden vorgegeben
- Für die Verwendung dieser Daten ist ein grundlegendes Verständnis des Lebenszyklus einer View erforderlich

### Hinweg

- Ziehen Sie mit CTRL-Drag eine Verbindung von dem jeweiligen Vor-Button zum Ziel View, wählen Sie dabei die Action-Option "Show"
- Diese Verbindung bezeichnet Apple als `Segue`

### Rückweg

`@IBAction func zurückZu2(segque: UIStoryBoardSegue{ ... }`

- Der Name der Methode ist beliebig, sollte aber eindeutig sein
- Die Parameterliste ist fest vorgegeben
- Um die Verbindung aufzubauen ziehen Sie mit CTRL+Drag eine Verbindung vom Button zum exit-Button in der gleichen View
- In dem daraufhi auftauchenden Dialog wählen Sie die Methode der Ziel-View
- Der Weg zurück ist auch über mehrere Views möglich
- Der Übergang kann auch aus dem Programm ausgelöst werden durch den Aufruf der Methode
  - `performSegue(withIdentifier: "SegueName", sender: self)`
- Zum Übertragen von Informationen auf dem Rückweg kann die bereits programmierte `unwind`-Methode verwendet werden
- Der Übergabeparameter `segue` dieser Methode erlaubt den Zugriff auf das Zielobjekt

### Daten übertragen

```swift
@IBAction func zurückZu2(segue: UIStoryBoardSegue) {
  if let quelle = segue.source as? ViewController2 {
    print(quelle.label2.text)
  }
}
```

- Um Informationen vorwärts zu Übertragen muss im Quellobjekt die Methode prepare überschrieben werden
- Die Übergabeparameter `segue` dieser Methode erlaubt den Zugriff auf das Zielobjekt
- Über `segue.destination` das **nur auf Variablen** der Zielklasse zugegriffen werden
- Die Oberflächenelemente der View sind zu diesem Zeitpunkt noch nicht erzeugt
- Die Übertragung aus Variablen der GUI-Elemente muss in der `viewDidLoad()` Methode erfolgen

```swift
override func prepare(for segue: UIStoryBoardSegue, sender: Any?) {
  if let ziel = segue.destination as? ViewController2 {
    ziel.variable = 123
  }
}

override func viewDidLoad() {
  super.viewDidLoad()

  guiElement.text = variable
}
```

## Lebenszyklus

- viewDidLoad() (einmalig)
  - Geometrie des VC ist hier noch unbekannt
- viewWillAppear()
  - Kann mehrfach auftreten (z.B. durch Rotation)
  - Geometrie ist hier bekannt
- viewWillLayoutSubview() (mehrfach)
  - AutoLayout
- viewDidLayoutSubviews()
- viewDidAppear()
- viewWillDisappear()
- viewDidDisappear()

- viewWillTransitionToSize()
  - Wird hauptsächlich bei Rotation aufgerufen
- didReceivememoryWarning()
  - Hier sollte nicht benötigter Speicher freigegeben werden
- awakeFromNib
  - Outlets sind noch nicht gesetzt
  - Wird ausgeführt wenn der VC aus einem InterfaceBuilder Archive File (NIB File) geladen wird
  - Outlets/Actions sind hier bereits gesetzt

```
  --------------------------
  |     viewDidLoad()      |
  --------------------------
              v
  --------------------------
  |viewWillLayoutSubviews()|<----------
  --------------------------          |
              v                       |
  --------------------------          |
  |viewDidLayoutSubviews() |<----     |
  --------------------------    |     |
              v                 |     |
  --------------------------    |     |
  |    viewDidAppear()     |-----     |
  --------------------------          |
              v                       |
  --------------------------          |
  |   viewWillDisappear()  |-----------
  --------------------------
              v
  --------------------------
  |    viewDidDisappear()  |
  --------------------------
```

## Delegation

- Delegation ist ein Entwurfsmuster (Design Pattern) bei dem eine Struktur oder eine Klasse einen teil ihrer Funktionalität in eine Instanz eines anderen Typs ausgelagert (delegiert)
- Die Umsetzung erfolgt durch die Definition eines Protokolls das die ausgelagerten Funktionalitäten beschreibt
- Die als Delegate konfigurierte Instanz wird über auftretende Ereignisse informiert
