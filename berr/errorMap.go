package berr

import (
	"errors"
	"fmt"
)

type ErrorMap map[string]error

func (errs ErrorMap) Set(k string, v error) {
	if val, ok := errs[k]; ok {
		errs[k] = errors.Join(val, v)
	} else {

		errs[k] = v
	}

}

func (errs ErrorMap) IsNil() bool {
	return len(errs) == 0
}

func (errs ErrorMap) Error() string {

	message := ""

	for k, v := range errs {
		if message == "" {
			message += fmt.Sprintf("%v=%v", k, v.Error())
		} else {
			message += fmt.Sprintf(";%v=%v", k, v.Error())
		}
	}

	return message
}

func (errs ErrorMap) Delete(key string) {
	delete(errs, key)
}

func (errs ErrorMap) ToStringSlice() [][2]string {
	_arr := [][2]string{}

	for k, v := range errs {
		_arr = append(_arr, [2]string{k, v.Error()})
	}

	return _arr
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
