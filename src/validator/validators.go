package validator

import (
	"cpftools/src/calculator"
	"cpftools/src/utils"
	"errors"
	"fmt"
)

func hasCorrectLength(cpf string) bool {
	const cpfLength int8 = 11
	cleanedCpf := utils.GetOnlyDigits(cpf)
	return len([]rune(cleanedCpf)) == int(cpfLength)
}

func isBlocked(cpf string) bool {
	cpf = utils.GetOnlyDigits(cpf)
	firstDigit := string([]rune(cpf)[0:1])
	for _, i := range cpf {
		if firstDigit != string(i) {
			return false
		}
	}
	return true
}

func getVerifyDigits(cpf string) string {
	cpf = utils.GetOnlyDigits(cpf)
	firstVerifiedDigit := calculator.CalcDigit(cpf, true)
	lastVerifiedDigit := calculator.CalcDigit(cpf, false)
	verifiedDigits := fmt.Sprintf("%d%d", firstVerifiedDigit, lastVerifiedDigit)
	return verifiedDigits
}

func IsValid(cpf string) bool {

	if !hasCorrectLength(cpf) {
		fmt.Println(errors.New("Invalid length"))
	}

	if isBlocked(cpf) {
		fmt.Println(errors.New("Invalid CPF"))
	}
	cpf = utils.GetOnlyDigits(cpf)
	currentVerificationDigits := string([]rune(cpf)[9:11])
	verifiedDigits := getVerifyDigits(cpf)
	return currentVerificationDigits == verifiedDigits
}
