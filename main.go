package main

import "math/rand"

const (
	duplicateVowel bool = true
)

func randBool() bool {
	return rand.Intn(2) == 0
}
