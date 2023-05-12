package main

import (
	"ShitCalculator/calculate"
	"ShitCalculator/getFormula"
)

func main() {
	getFormula.GenerateFormula()
	calculate.GetAnswer()
}
