package berr

import (
	"fmt"
	"strings"
)

// RtnPrint prints concatenated messages to stdout and returns the provided error.
// It joins the variadic msg arguments using ";" as a separator before printing.
func RtnPrint(e error, msg ...string) error {
	finalMsg := strings.Join(msg, ";")
	fmt.Println(finalMsg)
	return e
}

// RtnStr executes a function to generate a message, prints it to stdout, and returns the provided error.
// The function parameter allows custom message formatting logic to be applied before printing.
func RtnStr(e error, fn func(msg ...string) string) error {
	finalMsg := fn()
	fmt.Println(finalMsg)
	return e
}

// RtnExec executes a function for side effects, then returns the provided error.
// The function parameter is called but no return value is used, making it useful for custom error handling logic.
func RtnExec(e error, fn func(msg ...string)) error {
	fn()
	return e
}
