package assert

type Error struct {
	Message string
}

func NewError(message string) *Error {
	return &Error{message}
}

func (as *Error) Error() string {
	return "Assertion error: "+as.Message
}

func Assert(expr bool, message string) {
	if !expr {
		panic(NewError(message))
	}
}
