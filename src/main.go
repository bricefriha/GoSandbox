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
	num, ptr = PushInt(10);

	ptr = &num;
	fmt.Print(" variable:", num);
	fmt.Print("\n pointer:", ptr);
}