package englishtocron

import "fmt"

// Error represents the different kinds of errors that can occur.
type Error struct {
	Kind    ErrorKind
	State   string
	Token   string
	Value   string
	ErrDesc string
}

type ErrorKind int

const (
	ErrInvalidInput ErrorKind = iota
	ErrCapture
	ErrParseToNumber
	ErrIncorrectValue
)

func (e *Error) Error() string {
	switch e.Kind {
	case ErrInvalidInput:
		return "Please enter human readable"
	case ErrCapture:
		return fmt.Sprintf("Could not capture: %s in state: %s ", e.Token, e.State)
	case ErrParseToNumber:
		return fmt.Sprintf("Could not parse: %s to number. state: %s ", e.Value, e.State)
	case ErrIncorrectValue:
		return fmt.Sprintf("value is invalid in state: %s. description: %s ", e.State, e.ErrDesc)
	}
	return "unknown error"
}

func errInvalidInput() *Error {
	return &Error{Kind: ErrInvalidInput}
}

func errCapture(state, token string) *Error {
	return &Error{Kind: ErrCapture, State: state, Token: token}
}

func errParseToNumber(state, value string) *Error {
	return &Error{Kind: ErrParseToNumber, State: state, Value: value}
}

func errIncorrectValue(state, desc string) *Error {
	return &Error{Kind: ErrIncorrectValue, State: state, ErrDesc: desc}
}
