package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVowel bool = true
)

func randBool() bool {
	return rand.Intn(2) == 0
}
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		if randBool() {
			var vI int = -1
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						vI = i
					}
				}
			}
		}
	}
}
