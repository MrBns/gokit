package bns

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrbns/gokit/berr"
)

// Safely decoe body. becuase this fucntion will check if body is not empty.
//
// Return a  berr.ErrorMap map instance.
func SafeDecodeBodyToJson(r *http.Request, v any) berr.ErrorMap {
	errorMap := berr.NewErrorMap()

	if r.Body == http.NoBody {
		errorMap.Set("body", fmt.Errorf("request body is empty"))
		return errorMap
	}

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		errorMap.Set("body-decode", err)
	}

	return errorMap
}
