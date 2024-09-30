package helpers

import "math/rand"

type randomHelper struct{}

func (randomHelper) RandomRange(min int, max int) int {
	return rand.Intn(max+1-min) + min
}
