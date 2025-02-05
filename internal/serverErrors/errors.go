package servererrors

import "fmt"

type ConflictError struct {
	Used string
}

func (c ConflictError) Error() string {
	return fmt.Sprintf("%s already used", c.Used)
}

type UnauthorizedError struct {
	Message string
}

func (u UnauthorizedError) Error() string {
	return u.Message
}

type WrongNumberError struct {
	Number string
}

func (w WrongNumberError) Error() string {
	return w.Number
}

type OkayError struct{}

func (o OkayError) Error() string {
	return "we're feeling ok"
}
