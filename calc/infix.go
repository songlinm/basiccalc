package calc

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Infix evaluates a numerical calculation expression in infix notation. It either
// returns the evaulation result, or an error (e.g. divided by zero)
func Infix(r io.Reader) (ret float64, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	return infix(scanner)
}

func infix(scanner *bufio.Scanner) (float64, error) {
	if !scanner.Scan() {
		// no more words
		return 0, nil
	}

	var (
		val1, val2 float64
		err        error
	)
	// Evaluate the first argument
	arg1 := scanner.Text()
	if arg1 == "(" {
		if val1, err = infix(scanner); err != nil {
			return 0, err
		}

		// consume the ending parenthesis
		if !scanner.Scan() || scanner.Text() != ")" {
			return 0, fmt.Errorf(`Unmatched (`)
		}
	} else if val1, err = strconv.ParseFloat(arg1, 64); err != nil {
		return 0, fmt.Errorf(`Token %s is not a number: %v`, arg1, err)
	}

	// Evaluate the eval
	if !scanner.Scan() {
		// allow having one single number
		return val1, nil
	}

	// The next token should be the eval
	op := scanner.Text()

	// Evaluate the second argument
	if !scanner.Scan() {
		return 0, fmt.Errorf(`Missing second argument for %s %s`, arg1, op)
	}
	arg2 := scanner.Text()
	if arg2 == "(" {
		if val2, err = infix(scanner); err != nil {
			return 0, err
		}

		// consume the ending parenthesis
		if !scanner.Scan() || scanner.Text() != ")" {
			return 0, fmt.Errorf(`Unmatched (`)
		}
	} else if val2, err = strconv.ParseFloat(arg2, 64); err != nil {
		return 0, fmt.Errorf(`Token %s is not a number: %v`, arg2, err)
	}

	return eval(op, val1, val2)
}
