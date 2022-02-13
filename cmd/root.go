package cmd

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

const VERSION = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "suan",
	Short: "Mathematical expression calculation tool.",
	Long: `Suan( 算 ) is a CLI tool to calculate given mathematical expression.
Currently it supports addtion, substraction, multiplication, division, and
exponent operations and any of their combinations with parenthesis.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("suan: %s", VERSION)
			return
		}
		if len(args) < 1 {
			cmd.Help()
			return
		}

		exp := strings.Join(args, "")
		result, err := calculate(exp)
		if err != nil {
			fmt.Printf("invalid mathematical expression: %s. error: %v", exp, err)
			return
		}
		fmt.Printf("%f", result)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Version information")
}

func calculate(expression string) (float64, error) {
	// step1: get reverse polish notation from infix expression
	rpn, err := generateRPN(expression)
	if err != nil {
		return 0.0, err
	}

	// step2: calculate reverse polish notation
	return calculateRPN(rpn)
}

// generateRPN implements Shunting-yard algorithm which converts
// infix mathematical expression to reverse polish notation.
// For example:
//
//    (3 + 4) * 5 - 2 * (3 + 9) --> (3 4 + 5 * 2 3 9 + * -)
//
// See details at https://en.wikipedia.org/wiki/Shunting-yard_algorithm
// or https://old-panda.com/2020/11/29/shunting-yard-algorithm/ (in Chinese)
func generateRPN(expression string) ([]interface{}, error) {
	numbers := make([]interface{}, 0)
	operators := make([]string, 0)
	var num float64
	var foundNum bool = false
	for i, char := range expression {
		if unicode.IsDigit(char) {
			if !foundNum {
				num = 0.0
			}
			foundNum = true
			digit, err := strconv.ParseFloat(string(char), 64)
			if err != nil {
				return nil, err
			}
			num = num*10 + digit
		} else {
			if foundNum {
				numbers = append(numbers, num)
				num = 0.0
				foundNum = false
			}
			if char == ' ' {
				continue
			} else if char == rune('^') {
				for len(operators) > 0 && strings.Contains("^", operators[len(operators)-1]) {
					numbers = append(numbers, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}
				operators = append(operators, string(char))
			} else if char == rune('*') || char == rune('/') {
				for len(operators) > 0 && strings.Contains("^*/", operators[len(operators)-1]) {
					numbers = append(numbers, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}
				operators = append(operators, string(char))
			} else if char == rune('+') || char == rune('-') {
				for len(operators) > 0 && strings.Contains("^*/+-", operators[len(operators)-1]) {
					numbers = append(numbers, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}
				operators = append(operators, string(char))
			} else if char == rune('(') {
				operators = append(operators, string(char))
			} else if char == rune(')') {
				for len(operators) > 0 && operators[len(operators)-1] != "(" {
					numbers = append(numbers, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}
				if len(operators) > 0 && operators[len(operators)-1] == "(" {
					operators = operators[:len(operators)-1]
				}
			} else {
				return nil, fmt.Errorf("unknown char at pos %d: %s", i, string(char))
			}
		}
	}

	if foundNum {
		numbers = append(numbers, num)
	}
	for len(operators) > 0 {
		numbers = append(numbers, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}
	return numbers, nil
}

// calculateRPN calculates the result given math expression
// in reverse polish notation
func calculateRPN(rpn []interface{}) (float64, error) {
	stack := make([]float64, 0)
	for _, item := range rpn {
		switch v := item.(type) {
		case float64:
			stack = append(stack, v)
		case string:
			stackLen := len(stack)
			if stackLen < 2 {
				return 0.0, errors.New("invalid syntax")
			}
			a := stack[stackLen-2]
			b := stack[stackLen-1]
			stack = stack[:stackLen-2]
			if v == "+" {
				stack = append(stack, a+b)
			} else if v == "-" {
				stack = append(stack, a-b)
			} else if v == "*" {
				stack = append(stack, a*b)
			} else if v == "/" {
				stack = append(stack, a/b)
			} else if v == "^" {
				stack = append(stack, math.Pow(a, b))
			} else {
				return 0.0, fmt.Errorf("unknown operator: %s", v)
			}
		}
	}
	if len(stack) > 1 {
		return 0.0, errors.New("invalid syntax")
	}
	return stack[len(stack)-1], nil
}
