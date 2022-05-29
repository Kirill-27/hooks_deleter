package main

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const MaxBytes = 1000000

func main() {
	inFile, err := os.Open("test.txt")
	defer inFile.Close()
	check(err)
	b1 := make([]byte, MaxBytes)
	n1, err := inFile.Read(b1)
	check(err)
	input := string(b1[:n1])
	var res []string
	isGood := true
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			isGood = false
		}
		if isGood {
			res = append(res, string(input[i]))
		}
		if input[i] == ')' {
			isGood = true
		}
	}
	outFile, err := os.Create("result.txt")
	defer outFile.Close()
	check(err)
	outFile.WriteString(strings.Join(res, ""))
}
