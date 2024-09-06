package chapterOne

import (
	"fmt"
	"math"
	"math/rand"
)
// Random number
func RandomNum(){
	fmt.Println("My favorite number is", rand.Intn(10));
}

func SquareSeven(){
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7));
}

func PrintPi(){
	fmt.Println(math.Pi); // the capital "P" allows the import
}