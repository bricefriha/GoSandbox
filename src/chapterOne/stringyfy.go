package chapterOne;
import (
	"fmt"
)

func swap(x,y string) (string, string){
	return y, x;
}

func PrintSwap(x,y string){
	a, b := swap(x,y);

	fmt.Println("\n");
	fmt.Println(a, b);

}