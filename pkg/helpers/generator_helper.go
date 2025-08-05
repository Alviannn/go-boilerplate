package helpers

import (
	"math/rand/v2"
	"time"
)

type (
	GeneratorRandomStringParam struct {
		Length        int
		WithUpperCase bool
		IncludeNumber bool
		IncludeSymbol bool
		Seed          uint64
	}

	GeneratorRandomBoolParam struct {
		Seed uint64
	}

	GeneratorRandomIntParam struct {
		Min  int
		Max  int
		Seed uint64
	}
)

var (
	charsetLower  = "abcdefghijklmnopqrstuvwxyz"
	charsetUpper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumber = "0123456789"
	charsetSymbol = "!@#$%^&*()_+"
)

func GeneratorRandomString(param GeneratorRandomStringParam) string {
	if param.Seed == 0 {
		param.Seed = uint64(time.Now().UnixNano())
	}

	var (
		randomer         = rand.New(rand.NewPCG(param.Seed, param.Seed))
		generateFuncList = []func() string{
			func() string {
				return string(charsetLower[randomer.IntN(len(charsetLower))-1])
			},
		}
	)

	if param.WithUpperCase {
		generateFuncList = append(generateFuncList, func() string {
			return string(charsetUpper[randomer.IntN(len(charsetUpper))-1])
		})
	}
	if param.IncludeNumber {
		generateFuncList = append(generateFuncList, func() string {
			return string(charsetNumber[randomer.IntN(len(charsetNumber))-1])
		})
	}
	if param.IncludeSymbol {
		generateFuncList = append(generateFuncList, func() string {
			return string(charsetSymbol[randomer.IntN(len(charsetSymbol))-1])
		})
	}

	var result string
	for i := 0; i < param.Length; i++ {
		result += generateFuncList[randomer.IntN(len(generateFuncList))]()
	}
	return result
}

func GeneratorRandomBool(param GeneratorRandomBoolParam) bool {
	if param.Seed == 0 {
		param.Seed = uint64(time.Now().UnixNano())
	}

	randomer := rand.New(rand.NewPCG(param.Seed, param.Seed))
	return randomer.IntN(2) == 1
}

func GeneratorRandomInt(param GeneratorRandomIntParam) int {
	if param.Seed == 0 {
		param.Seed = uint64(time.Now().UnixNano())
	}

	randomer := rand.New(rand.NewPCG(param.Seed, param.Seed))
	return randomer.IntN(param.Max-param.Min+1) + param.Min
}
