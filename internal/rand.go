package internal

import "math/rand"

//GetRand - get random number
func GetRand(min, max int) int {
	return rand.Intn(max-min) + min
}
