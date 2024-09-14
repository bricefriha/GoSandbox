// Including the main package 
package main;

// Importing fmt 
import ( 
	"GoLangSanbox/src/chapterOne"
    "fmt"
) 
var c, python, java bool = true, false, "no!"
func main() {
	// step 1
	chapterOne.RandomNum();
	// step 2
	chapterOne.SquareSeven();
	// step 3
	chapterOne.PrintPi();
	// step 4 and 5
	chapterOne.PrintAddition(3,5);
	// step 6
	chapterOne.PrintSwap("Sky","Blue");
	// step 7
	chapterOne.PrintSplit(17);
	// step 8 and 9
	chapterOne.PrintArguments(c, python, java);
	// Step 9
	fmt.Print("\n---------");
	fmt.Print("\n Finished");
}