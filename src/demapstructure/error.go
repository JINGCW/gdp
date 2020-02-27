package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

//Error implements the error interface and can represents multiple
//errors that occur in the course of a single decode.
type Error struct {
	Errors []string
}

func (e *Error) Error() string {
	points := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		points[i] = fmt.Sprintf("* %s", err)
	}
	sort.Strings(points)
	return fmt.Sprintf(
		"%d error(s) decoding:\n\n%s",
		len(e.Errors), strings.Join(points, "\n"))
}

func (e *Error) WrappedErrors() []error {
	if e == nil {
		return nil
	}
	out := make([]error, len(e.Errors))
	for i, e := range e.Errors {
		out[i] = errors.New(e)
	}
	return out
}

func appendErrors(errors []string, err error) []string {
	switch e := err.(type) {
	case *Error:
		return append(errors, e.Errors...)
	default:
		return append(errors, e.Error())
	}
}
