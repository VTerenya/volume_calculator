//Main package.
package main

import (
	"fmt"
	"volume_calculator/handle"
)

func main() {
	calc := handle.Handle()
	fmt.Println(calc.Volume())
}
