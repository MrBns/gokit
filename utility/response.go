package bns

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrbns/gokit/berr"
)

func _ternary[T any](condition bool, truth T, falsy T) T {
	if condition {
		return truth
	} else {
		return falsy
	}
}

// structural way to send Response trough http.Response Writer
type Response struct {
	Data    any    `json:"data"`
	Err     error  `json:"error"`
	Status  int    `json:"status_code"`
	Msg     string `json:"message"`
	Success bool   `json:"success"`
	Meta    any    `json:"metadata"`
}

func (d Response) Write(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(d.Status)

	//Fixing Status Code
	if d.Status == 0 && d.Success {
		d.Status = 200
	} else if d.Status == 0 && !d.Success {
		d.Status = 400
	}

	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		w.Write(fmt.Appendf(nil, `{"data":null, "success":false, "status_code": 500, "error":%v,"message":%v, "metadata":null}`, Ternary(d.Err != nil, d.Err.Error(), "failed to parse response json"), Ternary(d.Msg != "", d.Msg, "internal server error")))
	}
}

// Overrite or Set Status in Struct
func (d *Response) SetStatus(status int) *Response {
	d.Status = status
	return d
}

// overrite or set Metadata in response struct;
func (r *Response) SetMeta(meta any) *Response {
	r.Meta = meta
	return r
}

// overrite or set Message in response struct;
func (r *Response) SetMessage(msg string) *Response {
	r.Msg = msg
	return r
}

// overrite or set Message in response struct;
//
// ⚠️ It will set success false and if message is empty it will set "failed to execute request."
func (r *Response) SetError(err error) *Response {
	if err == nil {
		return r
	}

	if r.Msg == "" {
		r.Msg = "failed to execute request."
	}
	r.Success = false
	r.Err = err
	return r
}

// Construct Error response but with Data field exist.
func ErrResponseWithData[T any](data T, err error, msg string) *Response {

	response := Response{
		Data:    data,
		Err:     err,
		Msg:     msg,
		Status:  500,
		Success: false,
	}
	return &response
}

// Construct Error Response with Error and msg;
func ErrResponse(err error, msg string) *Response {
	response := Response{
		Data:    nil,
		Err:     err,
		Msg:     msg,
		Status:  400,
		Success: false,
	}
	return &response
}

// Quick way to write Ok Response to responseWriter
func WriteErrResponse(err error, w http.ResponseWriter) {
	ErrResponse(err, err.Error()).SetStatus(500).Write(w)
}

// Construct OkResponse and it returns *Response
func OkResponse[T any](data T, msg string) *Response {
	response := Response{
		Data:    data,
		Msg:     msg,
		Err:     nil,
		Status:  200,
		Success: true,
	}
	return &response
}

// Quick way to write Ok Response to responseWriter
func WriteOkResponse[T any](data T, w http.ResponseWriter) {
	OkResponse(data, "request processed successfully").Write(w)
}

// Http Handler with Error Handling. Just Return Error. and if you don't want to return error
// then directly write SuccessResponse to ResponseWriter.
// Error also can be written diretly from controller but not recommended for code readablity.
//
// Support default error and also some Utility Error such as InternalError, BadRequest, ForbiddenRequest, AuthorizationRequired
// etc. check more on docs.
func HttpHandler(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)

		// Do not want extra conidtion checking if erro is empty
		if err == nil {
			return
		}

		// Checking if error is Berror. (Unified MrBns Error)
		if val, ok := err.(berr.BError); ok {

			ErrResponseWithData(
				val.GetData(),
				val.GetError(),
				_ternary(val.GetMessage() == "", val.GetError().Error(), val.GetMessage()),
			).
				SetStatus(val.GetStatus()).
				Write(w)

			return

		} else if val, ok := err.(berr.ErrorMap); ok {

			ErrResponse(val, "Something went wrong.").SetStatus(400).Write(w)
			return

		} else {

			ErrResponse(err, err.Error()).SetStatus(400).Write(w)
			return

		}

	}
}
