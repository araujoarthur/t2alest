package repl

import (
	"fmt"
)

var (
	ERNoPath          = RErrorNew(1, "path is needed but was not found")
	ERMissingParams   = RErrorNew(2, "there are missing parameters") // generic error for missing parameters
	ERWrongParamCount = RErrorNew(3, "wrong parameter count")
	ERNoResults       = RErrorNew(4, "the current search yielded no results")
)

type ERepl struct {
	Code    int32
	Message string
}

func (e *ERepl) Error() string {
	return fmt.Sprintf("repl error(%d): %s", e.Code, e.Message)
}

// Creates the error type
func RErrorNew(code int32, msg string) *ERepl {
	return &ERepl{Code: code, Message: msg}
}
