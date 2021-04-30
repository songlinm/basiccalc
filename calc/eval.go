package calc

import (
	"fmt"
)

// eval evaluates "arg1 op arg2"
func eval(op string, arg1, arg2 float64) (float64, error) {
	switch op {
	case "+":
		return arg1 + arg2, nil
	case "-":
		return arg1 - arg2, nil
	case "*":
		return arg1 * arg2, nil
	case "/":
		if arg2 == 0 {
			return 0, fmt.Errorf("Divided by zero: %s %f %f", op, arg1, arg2)
		}
		return float64(arg1) / float64(arg2), nil
	default:
		return 0, fmt.Errorf("Unsupported operator - %s %f %f", op, arg1, arg2)
	}
}
