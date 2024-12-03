/* q1. Create a function that converts string to an integer
    if any alphabets are passed, wrap strconv error and ErrStringValue error (create ErrStringValue error)

    ErrStringValue contains a message that 'value is of string type' and return the wrapped errors
    otherwise return the original error

    use the regex to check if value is of string type or not
    Hint: regexp.MatchString(`^[a-zA-Z]`, s)
    fmt.Errorf("%w %w") // to wrap error

    In main function check if ErrStringValue error was wrapped in the chain or not
    If yes, log a message 'value must be of int type not string' and log original error message alongside as well

*/

package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var ErrStringValue = errors.New("value is of string type")

func StringToInt(s string) (int, error) {
	isString, err := regexp.MatchString(`^[a-zA-Z]`, s)
	if err != nil {
		return 0, fmt.Errorf("regex error: %w", err)
	}

	if isString {
		return 0, fmt.Errorf("%w %w", strconv.ErrSyntax, ErrStringValue)
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("conversion error: %w", err)
	}

	return num, nil
}

func main() {
	result, err := StringToInt( "78xyz2")
	if err != nil {
		if errors.Is(err, ErrStringValue) {
			fmt.Printf("string value passed: %w\n", err)
			return
		}  
		
		fmt.Printf("conversion failed: %w\n", err)	
		return
		}
	fmt.Println("Result is", result)

	}
