package servererrors

import "fmt"

type ConflictError struct {
	Login string
}

func (c ConflictError) Error() string {
	return fmt.Sprintf("%s already used", c.Login)
}

type UnauthorizedError struct {
	Message string
}

func (u UnauthorizedError) Error() string {
	return u.Message
}
