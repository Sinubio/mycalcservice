package calculate

import (
	"errors"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	return parseExpression(expression)
}

func parseExpression(expr string) (float64, error) {
	var stack []float64
	var opStack []rune
	num := ""

	for _, ch := range expr {
		switch ch {
		case '+', '-', '*', '/':
			if num != "" {
				val, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return 0, err
				}
				stack = append(stack, val)
				num = ""
			}
			opStack = append(opStack, ch)

		default:
			if (ch >= '0' && ch <= '9') || ch == '.' {
				num += string(ch)
			} else {
				return 0, errors.New("unsupported character")
			}
		}
	}

	if num != "" {
		val, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return 0, err
		}
		stack = append(stack, val)
	}

	for len(opStack) > 0 {
		if len(stack) < 2 {
			return 0, errors.New("invalid expression")
		}
		b, a := stack[len(stack)-1], stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		var result float64
		switch opStack[len(opStack)-1] {
		case '+':
			result = a + b
		case '-':
			result = a - b
		case '*':
			result = a * b
		case '/':
			if b == 0 {
				return 0, errors.New("division by zero")
			}
			result = a / b
		}
		opStack = opStack[:len(opStack)-1]
		stack = append(stack, result)
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}
