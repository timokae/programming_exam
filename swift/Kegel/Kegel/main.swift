//
//  main.swift
//  Kegel
//
//  Created by Timo on 06.09.19.
//  Copyright © 2019 Timo. All rights reserved.
//

import Foundation

let pi : Float = 3.14159

func calculateVolume(height : Int, radius : Int) -> Double {
    return (1.0/3.0) * Double(pi) * Double(radius * radius * height)
}

func output(height : Int, radius : Int) -> Void {
    let volume : Double = calculateVolume(height: height, radius: radius)
    print("Das Volumen des Kegels beträgt: " + String(volume))
}

func readInput() -> (Int, Int)? {
    var values : (height: Int, radius: Int)
    
    print("Höhe: ")
    var input : String? = readLine(strippingNewline: true)
    if let i = Int(input ?? "") {
        values.height = i
    } else {
        return nil
    }
    
    print("Radius: ")
    input = readLine(strippingNewline: true)
    if let i = Int(input ?? "") {
        values.radius = i
    } else {
        return nil
    }
    
    return values
}


let input = readInput()
if let (height, radius) = input {
    output(height: height, radius: radius)
} else {
    print("Enter valid integer")
    exit(0)
}
