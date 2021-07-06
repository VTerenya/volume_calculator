package test

import (

	"github.com/VTerenya/volume_calculator/internal/calculator"
	"github.com/shopspring/decimal"
	"testing"
)

func Test(t *testing.T) {
	testTable := []struct {
		shape                         string
		radius, length, width, hieght decimal.Decimal
		expected                      decimal.Decimal
	}{
		{
			shape:    "sphere",
			radius:   decimal.NewFromFloat(1.5),
			expected: decimal.NewFromFloat(14.1371669411540681465708264711482875),
		},
		{
			shape:    "cylinder",
			hieght:   decimal.NewFromFloat(10.1),
			radius:   decimal.NewFromFloat(5.5),
			expected: decimal.NewFromFloat(959.835095488021506325),
		},
		{
			shape:    "pyramid",
			length:   decimal.NewFromFloat(5.5),
			width:    decimal.NewFromFloat(4.4),
			hieght:   decimal.NewFromFloat(3.85),
			expected: decimal.NewFromFloat(31.056666666666663561),
		},
	}

	for _, testCase := range testTable {
		result,_ := calculator.Calculate(testCase.shape, testCase.radius, testCase.length,
			testCase.width, testCase.hieght)
		if result.Equals(testCase.expected) {
			t.Errorf("Error!\n Expected : %#v;\nResult: %#v\n", testCase.expected, result)
		}
	}
}
