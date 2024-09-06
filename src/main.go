// Including the main package 
package main;

// Importing fmt 
import ( 
	"GoLangSanbox/src/pointer"
    "fmt"
) 

func main() {
	var num int;
	var ptr *int;
	num, ptr = pointer.PushInt(10);
	
	fmt.Print(" variable:", num);
	fmt.Print("\n pointer:", ptr);
}