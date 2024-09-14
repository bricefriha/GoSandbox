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
func add(a, b int) int{
	return a + b
}
func PrintAddition(x, y int){
	xstr := strconv.Itoa(x);
	ystr := strconv.Itoa(y);
	fmt.Printf("%v+%v=%v",xstr,ystr, strconv.Itoa(add(x,y)));
}
func split(sum int) (x,y int){
	x = sum *4/9;
	y=sum-x

	return
}
func PrintSplit(sum int){
	fmt.Println(split(sum));
}
func PrintArguments(arg1 bool, arg2 bool, arg3 string){
	var i int;
	fmt.Println(i, arg1, arg2, arg3);
}
func Print4Arguments(arg1 int, arg2 int, arg3 int, arg4 bool, arg5 bool, arg6 string){
	fmt.Println(arg1, arg2, arg3, arg4, arg5, arg6);
}