package main

import (
	"fmt"
	"math"
)

//tanpa parameter
func sayHello() {
	fmt.Println("hello")
}

//dg parameter
func greeting(hour int) {
	if hour < 12 {
		fmt.Println("pagi")
	} else if hour < 18 {
		fmt.Println("sore")
	} else {
		fmt.Println("malam")
	}
}

//single return value
func calculateSquare(side int) int {
	return side * side
}

//multiple return value
func calculateCircle(diameter float64) (float64, float64) {
	var keliling = math.Pi * math.Pow(diameter/2, 2)
	var luas = math.Pi * diameter
	return keliling, luas
}

func main() {
	hour := 15
	sayHello()
	greeting(hour)

	var side = 5
	wide := calculateSquare(side)
	fmt.Printf("luas persegi : %d\n\n", wide)

	var diameter float64 = 15
	keliling, luas := calculateCircle(diameter)

	fmt.Printf("luas Lingkaran : %.2f \n", keliling)
	fmt.Printf("keliling : @.2f\n", luas)
}
