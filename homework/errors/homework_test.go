package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	sb := strings.Builder{}
	for _, v := range e.errors {
		sb.WriteString("\t* ")
		sb.WriteString(v.Error())
	}
	return fmt.Sprintf("%d errors occured:\n%s\n", len(e.errors), sb.String())
}

func Append(err error, errs ...error) *MultiError {
	res := &MultiError{}

	me, ok := err.(*MultiError)
	if ok {
		res.errors = append([]error(nil), me.errors...)
		res.errors = append(res.errors, errs...)
		return res
	}

	res.errors = append([]error(nil), errs...)
	return res
}

func (e *MultiError) Is(err error) bool {
	for _, v := range e.errors {
		if err == v {
			return true
		}
	}
	return false
}

func (e *MultiError) As(target any) bool {
	val := reflect.ValueOf(target)
	typ := val.Type().Elem()
	for _, v := range e.errors {
		if reflect.TypeOf(v).AssignableTo(typ) {
			val.Elem().Set(reflect.ValueOf(v))
			return true
		}
	}
	return false
}

func (e *MultiError) Unwrap() error {
	if len(e.errors) == 0 {
		return nil
	}
	return &MultiError{errors: append([]error(nil), e.errors[:len(e.errors)-1]...)}
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
