package berr

import (
	"encoding/json"
	"fmt"
	"strings"
)

const ErrorSeparator = "%::%"

type ErrorMap map[string][]error

func (errs ErrorMap) Set(k string, v error) {

	if val, ok := errs[k]; ok {
		errs[k] = append(val, v)
	} else {

		errs[k] = []error{v}
	}

}

func (errs ErrorMap) IsNil() bool {
	return len(errs) == 0
}

func (errs ErrorMap) HasErr() bool {
	return len(errs) != 0
}

func (errs ErrorMap) Error() string {

	message := new(strings.Builder)
	first := true

	for k, v := range errs {
		if !first {
			message.WriteString(";")
		}
		message.WriteString(fmt.Sprintf("%v=%v", k, v))
		first = false

	}

	return message.String()
}

func (errs ErrorMap) Delete(key string) {
	delete(errs, key)
}

func (errs ErrorMap) MarshalJSON() ([]byte, error) {
	serialized := make(map[string][]string, len(errs))
	for k, v := range errs {
		values := make([]string, 0, len(v))
		for _, e := range v {
			if e == nil {
				values = append(values, "")
				continue
			}
			values = append(values, e.Error())
		}
		serialized[k] = values
	}

	return json.Marshal(serialized)
}

// Helper function to initialize new Error map.
func NewErrorMap() ErrorMap {
	return make(ErrorMap)
}

// Check if ErrorMap is Empty Error is a empty error
//
// it would return true if the error type is ErrorMap and error is empty.
// it would return false if error type is not ErrorMap or ErrorMap is not empty.
func IsErrorMapEmpty(err error) bool {
	val, ok := err.(ErrorMap)

	if !ok {
		return false
	}
	return val.IsNil()

}
