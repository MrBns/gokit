package berr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type ErrorMap map[string]error

func (errs ErrorMap) Set(k string, v error) {
	if val, ok := errs[k]; ok {
		errs[k] = errors.New(val.Error() + "," + v.Error())
	} else {

		errs[k] = v
	}

}

func (errs ErrorMap) IsNil() bool {
	return len(errs) == 0
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

func (errs ErrorMap) ToStringSlice() [][2]string {
	_arr := [][2]string{}

	for k, v := range errs {
		_arr = append(_arr, [2]string{k, v.Error()})
	}

	return _arr
}

func (errs ErrorMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteByte('{')
	first := true

	for k, v := range errs {

		if !first {
			buf.WriteByte(',')
		}

		buf.WriteString(`"` + k + `":`)

		errsBytes, err := json.Marshal(strings.Split(v.Error(), ","))
		if err != nil {
			buf.WriteByte(' ')
		} else {

			buf.Write(errsBytes)
		}

		first = false
	}

	buf.WriteByte('}')

	return buf.Bytes(), nil
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
