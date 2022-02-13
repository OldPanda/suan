package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRPN(t *testing.T) {
	exp := "(3 + 4) * 5 - 2 * (3 + 9)"
	expected := "3.000000 4.000000 + 5.000000 * 2.000000 3.000000 9.000000 + * -"
	rpn, err := generateRPN(exp)
	rpnTokens := make([]string, 0)
	for _, item := range rpn {
		switch v := item.(type) {
		case string:
			rpnTokens = append(rpnTokens, v)
		case float64:
			rpnTokens = append(rpnTokens, fmt.Sprintf("%f", v))
		}
	}
	rpnStr := strings.Join(rpnTokens, " ")
	assert.Nil(t, err, fmt.Sprintf("Expected no error: %v", err))
	assert.Equal(t, expected, rpnStr, fmt.Sprintf("Expected: %s, got: %s", expected, rpnStr))

	invalidExp := "(1 + 2) * abc"
	_, err = generateRPN(invalidExp)
	assert.Error(t, err, "Expected error, got nil")
}

func TestCalculate(t *testing.T) {
	exp := "(1 + 2)^3 * 4 - 5/(6 + 7) * 8^9"
	expected := -5.162209507692308e+07
	result, err := calculate(exp)
	assert.Nil(t, err, fmt.Sprintf("Expected no error: %v", err))
	assert.Equal(t, expected, result, fmt.Sprintf("Expected: %f, got: %f", expected, result))
}
