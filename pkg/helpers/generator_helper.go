package helpers

import (
	"math/rand/v2"
)

type (
	GeneratorRandomStringParam struct {
		Length        int
		WithUpperCase bool
		IncludeNumber bool
		IncludeSymbol bool
	}

	GeneratorRandomIntParam struct {
		Min int
		Max int
	}
)

var (
	charsetLower  = "abcdefghijklmnopqrstuvwxyz"
	charsetUpper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumber = "0123456789"
	charsetSymbol = "!@#$%^&*()_+"
)

func GeneratorRandomString(param GeneratorRandomStringParam) string {
	var (
		generateFuncList = []func() string{
			func() string {
				return string(charsetLower[rand.IntN(len(charsetLower))-1])
			},
		}
	)

	if param.WithUpperCase {
		generateFuncList = append(generateFuncList, func() string {
			return string(charsetUpper[rand.IntN(len(charsetUpper))-1])
		})
	}
	if param.IncludeNumber {
		generateFuncList = append(generateFuncList, func() string {
			return string(charsetNumber[rand.IntN(len(charsetNumber))-1])
		})
	}
	if param.IncludeSymbol {
		generateFuncList = append(generateFuncList, func() string {
			return string(charsetSymbol[rand.IntN(len(charsetSymbol))-1])
		})
	}

	var result string
	for i := 0; i < param.Length; i++ {
		result += generateFuncList[rand.IntN(len(generateFuncList))]()
	}
	return result
}

func GeneratorRandomBool() bool {
	return rand.IntN(2) == 1
}

func GeneratorRandomInt(param GeneratorRandomIntParam) int {
	return rand.IntN(param.Max-param.Min+1) + param.Min
}
