package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type QueryError struct {
	Func  string
	Input string
	Err   error
}

var ErrNotFound = errors.New("not found")

func (q *QueryError) Error() string {
	return "main." + q.Func + ": " + "input " + q.Input + " " + q.Err.Error()
}

func SearchSomething(id int) (string, error) {
	// assume that search code is written and we need to return an error
	return "", &QueryError{
		Func:  "SearchSomething",
		Input: strconv.Itoa(id),
		Err:   ErrNotFound,
	}
}

func main() {

	user, err := SearchSomething(11)
	if err != nil {
		var qe *QueryError // nil
		if errors.As(err, &qe) {
			fmt.Println(qe.Func)
			fmt.Println(qe.Err)
			fmt.Println(err)
			return
		}
		log.Println(err)
		return
	}

	log.Println(user)

}
