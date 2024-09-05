// Including the main package 
package main;

// Importing fmt 
import ( 
	"GoLangSanbox/src/pp"
    "fmt"
) 

func main() {
	var num int;
	var ptr *int;
	num, ptr = pp.PushInt(10);
	
	fmt.Print(" variable:", num);
	fmt.Print("\n pointer:", ptr);
}