// Including the main package 
package main;

// Importing fmt 
import ( 
	"GoLangSanbox/src/chapterOne"
    "fmt"
) 

func main() {
	// step 1
	chapterOne.RandomNum();
	// step 2
	chapterOne.SquareSeven();
	// step 3
	chapterOne.PrintPi();
	// step 4
	chapterOne.PrintAddition(3,5);
	fmt.Print("\n---------");
	fmt.Print("\n Finished");
}