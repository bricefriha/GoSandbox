package packages

import (
	"fmt"
	"math/rand"
)

func RandomNum(){
	fmt.Println("My favorite number is", rand.Intn(10))
}