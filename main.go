//Main package.
package main

import (
	"fmt"
	"volume_calculator/cli"
)

func main() {
	calc := cli.Handle()
	fmt.Println(calc.CalcVolume())
}
