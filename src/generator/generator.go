package generator

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/drummerzzz/cpftools/src/calculator"
)

func generateVerifyDigits(cpfWithoutVerifyDigits []string) string {
	cpf := strings.Join(cpfWithoutVerifyDigits, "")
	firstDigit := calculator.CalcDigit(cpf, true)

	cpfWithFirstVerifyDigit := append(cpfWithoutVerifyDigits, strconv.Itoa(firstDigit))
	cpf = strings.Join(cpfWithFirstVerifyDigit, "")
	lastDigit := calculator.CalcDigit(cpf, false)
	verifiedDigits := fmt.Sprintf("%d%d", firstDigit, lastDigit)
	return verifiedDigits
}

func makeCpf(cpfWithoutVerifyDigits []string) string {
	baseCpf := strings.Join(cpfWithoutVerifyDigits, "")
	verifyDigits := generateVerifyDigits(cpfWithoutVerifyDigits)
	str := fmt.Sprintf("%s%s", baseCpf, verifyDigits)
	return str
}

func Generate(withMask bool) string {
	var cpfWithoutVerifyDigits []string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 9; i++ {
		cpfWithoutVerifyDigits = append(cpfWithoutVerifyDigits, strconv.Itoa(rand.Intn(9)))
		if withMask {
			if i == 2 || i == 5 {
				cpfWithoutVerifyDigits = append(cpfWithoutVerifyDigits, ".")
			} else if i == 8 {
				cpfWithoutVerifyDigits = append(cpfWithoutVerifyDigits, "-")
			}
		}
	}
	return makeCpf(cpfWithoutVerifyDigits)
}
