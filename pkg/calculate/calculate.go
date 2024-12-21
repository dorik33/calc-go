package calculate

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ErrInvalidInput   = "Invalid input"
	ErrDivisionByZero = "Division by zero"
)

func Calculate(expression string) (float64, error) {
	tokens := tokenizeExpr(expression)
	postfix, err := toPostfix(tokens)
	if err != nil {
		return 0, err
	}
	return evalPostfix(postfix)
}

func isNum(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOp(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func prec(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func tokenizeExpr(expr string) []string {
	var tokens []string
	var curToken strings.Builder

	for _, char := range expr {
		switch char {
		case ' ':
			continue
		case '+', '-', '*', '/', '(', ')':
			if curToken.Len() > 0 {
				tokens = append(tokens, curToken.String())
				curToken.Reset()
			}
			tokens = append(tokens, string(char))
		default:
			curToken.WriteRune(char)
		}
	}

	if curToken.Len() > 0 {
		tokens = append(tokens, curToken.String())
	}

	return tokens
}

func toPostfix(tokens []string) ([]string, error) {
	var output []string
	var ops []string

	for _, token := range tokens {
		if isNum(token) {
			output = append(output, token)
		} else if token == "(" {
			ops = append(ops, token)
		} else if token == ")" {
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				output = append(output, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			if len(ops) == 0 {
				return nil, fmt.Errorf(ErrInvalidInput)
			}
			ops = ops[:len(ops)-1]
		} else if isOp(token) {
			for len(ops) > 0 && prec(ops[len(ops)-1]) >= prec(token) {
				output = append(output, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, token)
		} else {
			return nil, fmt.Errorf(ErrInvalidInput)
		}
	}

	for len(ops) > 0 {
		if ops[len(ops)-1] == "(" {
			return nil, fmt.Errorf(ErrInvalidInput)
		}
		output = append(output, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}

	return output, nil
}

func evalPostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if isNum(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if isOp(token) {
			if len(stack) < 2 {
				return 0, fmt.Errorf(ErrInvalidInput)
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, fmt.Errorf(ErrDivisionByZero)
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf(ErrInvalidInput)
			}
		} else {
			return 0, fmt.Errorf(ErrInvalidInput)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf(ErrInvalidInput)
	}

	return stack[0], nil
}
