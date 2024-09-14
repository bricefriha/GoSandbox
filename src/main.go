// Including the main package 
package main;

// Importing fmt 
import ( 
	"GoLangSanbox/src/chapterOne"
    "fmt"
) 

func main() {
	var i, j int = 1,2
	k :=3 ;
	c, python, java := true, false, "no!"
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
	// Step 10
	chapterOne.Print4Arguments(i, j, k, c, python, java)
	fmt.Print("\n---------");
	fmt.Print("\n Finished");
}