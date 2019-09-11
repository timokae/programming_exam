//
//  ViewController.swift
//  helloSwift
//
//  Created by Timo on 08.09.19.
//  Copyright © 2019 Timo. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    @IBOutlet weak var temperatureLabel: UILabel!
    @IBOutlet weak var temperatureInput: UITextField!
    
    @IBAction func celciusClicked(_ sender: Any) {
        let input = temperatureInput.text
        
        if let f = Float(input ?? "") {
            temperatureLabel.text = String((5.0/9.0) * (f - 32.0))
        }
    }
    
    @IBAction func fahrenheitClicked(_ sender: Any) {
        let input = temperatureInput.text
        
        if let c = Float(input ?? "") {
            temperatureLabel.text = String((9.0/5.0) * c + 32.0)
            
        }
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
    }


}

