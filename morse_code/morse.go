package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	morseDict = map[string]string{
		".-":        "A",
		"-...":      "B",
		"-.-.":      "C",
		"-..":       "D",
		".":         "E",
		"..-.":      "F",
		"--.":       "G",
		"....":      "H",
		"..":        "I",
		".---":      "J",
		"-.-":       "K",
		".-..":      "L",
		"--":        "M",
		"-.":        "N",
		"---":       "O",
		".--.":      "P",
		"--.-":      "Q",
		".-.":       "R",
		"...":       "S",
		"-":         "T",
		"..-":       "U",
		"...-":      "V",
		".--":       "W",
		"-..-":      "X",
		"-.--":      "Y",
		"--..":      "Z",
		"   ":       " ",
		"...---...": "SOS",
	}
)

func iamregexpfun(source string) string {
	reTriple, _ := regexp.Compile("[\t\n\f\r ][\t\n\f\r ][\t\n\f\r ]")
	reSingle, _ := regexp.Compile("[\t\n\f\r ]")

	temp := reTriple.ReplaceAllString(source, "t")
	fmt.Println("re_temp:", temp)
	result := reSingle.ReplaceAllString(temp, "s")
	fmt.Println("re_result:", result)

	return result
}

func converterMorse(morseString string) []string {
	result := make([]string, len(morseString))

	for pos, char := range morseString {
		result[pos] = string(char)
	}

	return result
}

func tokenaizerMorse(morseCode string) []string {
	reString := iamregexpfun(morseCode)
	arrMorse := converterMorse(reString)
	fmt.Println("SOURCE ARR: ", arrMorse)
	tokens := make([]string, 0)

	var token string
	for i, v := range arrMorse {
		if (v == ".") || (v == "-") {
			token += v
		}

		if v == "t" {
			tokens = append(tokens, token)

			tokens = append(tokens, "   ")

			token = ""
		}

		if v == "s" {
			tokens = append(tokens, token)
			token = ""
		}

		if i == (len(arrMorse) - 1) {
			tokens = append(tokens, token)
		}

		fmt.Printf("TOKENS %s iterating #%d\n", tokens, i)

	}

	return tokens
}

func DecodeMorse(morseCode string) string {
	var result string
	tokens := tokenaizerMorse(morseCode)

	for _, v := range tokens {
		result += morseDict[v]
	}

	result = strings.TrimSpace(result)

	return result
}

func getFullDict(dict map[string]string) {
	for k, v := range dict {
		fmt.Printf("MORSE_CODE[%s] = %s\n", k, v)
	}
}

func getDifferenceFrom(m1, m2 map[string]string) map[string]string {
	result := make(map[string]string)

	for k, v1 := range m1 {
		_, ok := m2[k]
		if !ok {
			result[k] = v1
		}
	}

	return result
}

func main() {
	// src := ".... . -.--   .--- ..- -.. ."
	// src := "...---..."
	src := "   .   ."
	fmt.Printf("Source: %s\nResult: %s\n", src, DecodeMorse(src))

	srcz1 := map[string]string{
		".-":   "A",
		"-...": "B",
		"-.-.": "C",
		"-..":  "D",
	}

	srcz2 := map[string]string{
		".-":   "A",
		"-...": "B",
		"-.-.": "C",
	}

	fmt.Printf("Difference srcz1 and srcz2: %#v\n", getDifferenceFrom(srcz1, srcz2))

}
