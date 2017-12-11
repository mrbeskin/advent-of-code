package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	trashOpen  = '<'
	trashClose = '>'
	cancel     = '!'
	groupOpen  = '{'
	groupClose = '}'
)

func main() {
	inBytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input := string(inBytes)
	depth := 0
	inTrash := false
	total := 0
	trashCount := 0
	for pos, char := range input {
		if !inTrash {
			if char == groupOpen {
				depth += 1
			}
			if char == groupClose {
				total += depth
				depth--
			}
			if char == trashOpen {
				inTrash = true
			}
		} else {
			if char == trashClose {
				inTrash = false
				if input[pos-1] == cancel {
					if (countCancels(pos-1, input) % 2) > 0 {
						inTrash = true
					}
				}
			} else {
				if (!((countCancels(pos-1, input) % 2) > 0)) && (char != cancel) {
					trashCount++
				}
			}
		}
	}
	fmt.Println(total)
	fmt.Println(trashCount)
}

func countCancels(pos int, input string) int {
	done := false
	count := 0
	for done == false {
		if (pos <= 0) || (input[pos] != cancel) {
			done = true
		}
		if input[pos] == cancel {
			count += 1
		}
		pos--
	}
	return count
}
