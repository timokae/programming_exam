import Cocoa

func zufall(count: Int, min: Int, max: Int) -> [Int] {
    var numbers : [Int] = Array()
    
    for _ in 0..<count {
        numbers.append(Int.random(in: min..<max))
    }
    
    return numbers
}

func compare(a: Int, b: Int) -> Bool {
    return a > b
}

func bubble(numbers: inout [Int], compare: (Int, Int)->Bool) {
    for n in (2...numbers.count).reversed() {
        for i in 0..<(n-1) {
            if compare(numbers[i], numbers[i+1]) {
                let tmp = numbers[i]
                numbers[i] = numbers[i+1]
                numbers[i+1] = tmp
            }
        }
    }
}

func output(numbers: inout [Int]) {
    for i in 0..<numbers.count {
        var s : String = String(numbers[i])
        if i < numbers.count - 1 {
            s += ", "
        }
        
        print(s, terminator: "")
    }
}

var numbers = zufall(count: 10, min: 0, max: 10)
bubble(numbers: &numbers, compare: compare(a:b:))
output(numbers: &numbers)
