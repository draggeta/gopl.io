// Ex: add conversion to and from Kelvin

package main

import (
	"fmt"

	"ex1/tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(100))
	fmt.Println(tempconv.CToK(0))
	fmt.Println(tempconv.FToC(212))
	fmt.Println(tempconv.FToK(0))
	fmt.Println(tempconv.KToC(0))
	fmt.Println(tempconv.KToC(100))
}
