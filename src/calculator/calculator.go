package calculator

import (
	"cpftools/src/utils"
	"strconv"
)

func CalcDigit(cpf string, isFirstVerifyDigit bool) int {
	cpf = utils.GetOnlyDigits(cpf)
	maxFactory := 10
	if !isFirstVerifyDigit {
		maxFactory = 11
	}
	sum := 0
	cpfDigits := []rune(cpf)
	for currentFactory := maxFactory; currentFactory > 1; currentFactory-- {
		index := maxFactory - currentFactory
		number, _ := strconv.Atoi(string(cpfDigits[index]))
		sum += number * currentFactory
	}
	rest := sum % 11
	if rest < 2 || rest > 10 {
		return 0
	}
	return 11 - rest
}
