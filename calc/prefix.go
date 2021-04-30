package calc

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Prefix evaluates a numerical calculation expression in prefix notation. It either
// returns the evaulation result, or an error (e.g. divided by zero)
func Prefix(r io.Reader) (ret float64, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	if !scanner.Scan() {
		return
	}
	return prefixCalc(scanner.Text(), scanner)
}

func prefixCalc(op string, scanner *bufio.Scanner) (float64, error) {
	if !scanner.Scan() {
		// no more words
		if val, err := strconv.ParseFloat(op, 64); err == nil {
			return val, nil
		} else {
			return 0, fmt.Errorf(`Last token is not a number: %v`, err)
		}
	}

	// Evaluate the first argument
	arg1 := scanner.Text()
	val1, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		// token arg1 is an eval
		if val1, err = prefixCalc(arg1, scanner); err != nil {
			return 0, err
		}
	}

	// Evaluate the second argument
	if !scanner.Scan() {
		// no more words
		return 0, fmt.Errorf(`Missing second parameter for eval %s`, op)
	}
	arg2 := scanner.Text()
	val2, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		if val2, err = prefixCalc(arg2, scanner); err != nil {
			return 0, err
		}
	}
	return eval(op, val1, val2)
}
