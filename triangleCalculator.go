/*package main

import (
	"fmt"
)

func calculateSide(length1, length2, angle float64) float64 {

	//Insert the code here

	//Do not remove this line
	fmt.Println("The 3rd length of the triange is", length3)
	return length3
}

func main() {
	var length1 float64
	var length2 float64
	var angle float64

	fmt.Println("Enter the first length of the triangle: ")
	fmt.Scanln(&length1)
	fmt.Println("Enter the second length of the triangle: ")
	fmt.Scanln(&length2)
	fmt.Println("Enter the angle between the two lengths: ")
	fmt.Scanln(&angle)

	calculateSide(length1, length2, angle)
}
*/
package main

import (
	"fmt"
	"math"
)

func calculateSideC(sideA, sideB, angle float64) float64 {
	sideC := math.Sqrt(math.Pow(sideA, 2) + math.Pow(sideB, 2) - (2 * sideA * sideB * math.Cos(angle*math.Pi/180)))
	fmt.Println("The length of side C is", sideC)
	return sideC
}

func main() {
	var sideA float64
	var sideB float64
	var angle float64
	fmt.Println("Input side a")
	fmt.Scanln(&sideA)
	fmt.Println("Input side b")
	fmt.Scanln(&sideB)
	fmt.Println("Input an angle")
	fmt.Scan(&angle)
	calculateSideC(sideA, sideB, angle)

}
