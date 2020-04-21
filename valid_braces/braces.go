package main

import "fmt"

// "(){}[]"   =>  True
// "([{}])"   =>  True
// "(}"       =>  False
// "[(])"     =>  False
// "[({})](]" =>  False
func ValidBraces(str string) bool {
	var stack Stack
	var lastValue string
	s := convrt(str)

	for _, v := range s {
		//fmt.Println(v)
		if (v == "(") || (v == "[") || (v == "{") {
			stack.push(v)
			fmt.Printf("stack.push(\"%s\")\n", v)
			// lastValue = v
		} else {
			if stack.isEmpty() {
				return false
			}
			lastValue = stack.values[stack.siz()-1]
			if (v == ")" && lastValue == "(") || (v == "]" && lastValue == "[") || (v == "}" && lastValue == "{") {
				topEl := stack.pop()
				fmt.Printf("stack.pop() -> %s\n", topEl)
				fmt.Printf("stack size -> %d\n", stack.siz())
			} else {
				break
			}
		}
	}

	return stack.isEmpty()
}

func convrt(str string) []string {
	result := make([]string, len(str))

	for pos, char := range str {
		result[pos] = string(char)
	}

	return result
}

type Stack struct {
	values []string
}

func (s *Stack) push(elem string) {
	s.values = append(s.values, elem)
}

func (s *Stack) pop() string {
	if s.isEmpty() {
		return ""
	}

	topIndex := len(s.values) - 1
	topElement := s.values[topIndex]
	s.values = s.values[:topIndex]
	return topElement
}

func (s Stack) popWithoutSaving() string {
	if s.isEmpty() {
		return ""
	}

	topIndex := len(s.values) - 1
	topElement := s.values[topIndex]
	s.values = s.values[:topIndex]
	return topElement
}

func (s *Stack) siz() int {
	return len((*s).values)
}

func (s *Stack) isEmpty() bool {
	if len(s.values) == 0 {
		return true
	}

	return false
}

func (s Stack) printStack() {
	for !(s.isEmpty()) {
		fmt.Println(s.pop())
	}
}

func main() {

	cases := [...]string{
		"(){}[]",
		"([{}])",
		"(}",
		"[(])",
		"[({})](]",
	}

	for _, v := range cases {
		fmt.Printf("RESULT of %s -> %t\n", v, ValidBraces(v))
	}
}
