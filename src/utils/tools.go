package utils

import (
	"regexp"
	"strings"
)

func GetOnlyDigits(cpf string) string {
	re := regexp.MustCompile("\\d+")
	return strings.Join(re.FindAllString(cpf, -1), "")
}
