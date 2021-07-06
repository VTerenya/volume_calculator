//Main package.
package main

import (
	"fmt"
	"github.com/VTerenya/volume_calculator/internal/calculator"
	"github.com/VTerenya/volume_calculator/internal/cli"
)

func main() {
	cliParams := cli.Load()
	volume, err := calculator.Calculate(cliParams.Shape, cliParams.Radius, cliParams.Length,
		cliParams.Width, cliParams.Hieght)
	if err == nil {
		fmt.Println(volume)
	} else {
		fmt.Printf("Error: %v", err)
	}
}
