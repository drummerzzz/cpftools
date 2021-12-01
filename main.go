package main

import (
	"cpftools/src/generator"
	"cpftools/src/validator"
	"fmt"
)

// Generate a cpf with mask
func GenerateWithMask() string {
	return generator.Generate(true)
}

// Generate a cpf without mask
func GenerateWithoutMask() string {
	return generator.Generate(false)
}

// Check if cpf is valid
func IsValid(cpf string) bool {
	return validator.IsValid(cpf)
}

func main() {
	fmt.Println(GenerateWithMask())
	fmt.Println("--------")
	fmt.Println(GenerateWithoutMask())
}
