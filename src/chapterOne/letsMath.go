package chapterOne

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)
// Random number
func RandomNum(){
	fmt.Println("My favorite number is", rand.Intn(10));
}

func SquareSeven(){
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7));
}

func PrintPi(){
	fmt.Println("Pi is",math.Pi); // the capital "P" allows the import
}
func add(a int, b int) int{
	return a + b
}
func PrintAddition(x int, y int){
	xstr := strconv.Itoa(x);
	ystr := strconv.Itoa(y);
	fmt.Printf("%v + %v = %v",xstr,ystr, strconv.Itoa(add(x,y)));
}