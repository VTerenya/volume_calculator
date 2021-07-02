//Main package.
package main

import (
	"fmt"
	"volume_calculator/internal/calculator"
	"volume_calculator/internal/cli"
)

func main() {
	cliParams := cli.Load()
	volume := calculator.Calculator{}.Calculate(cliParams.Shape, cliParams.Radius, cliParams.Length,
		cliParams.Width, cliParams.Hieght)
	fmt.Println(volume)
}
