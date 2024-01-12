package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			l := len(stack)
			n1 := stack[l-1]
			n2 := stack[l-2]
			stack = stack[:l-2]
			switch v {
			case "+":
				stack = append(stack, n2+n1)
			case "-":
				stack = append(stack, n2-n1)
			case "*":
				stack = append(stack, n2*n1)
			case "/":
				stack = append(stack, n2/n1)
			}
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

type Example struct {
	token []string
	ans   int
}

func main() {
	examples := []Example{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for i, e := range examples {
		ans := evalRPN(e.token)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
