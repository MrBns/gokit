// BErr ~= BinarySniper Error. it will contain all Error utility
package berr

import (
	"errors"
	"net/http"
	"strings"
)

// Extended from [ error ]
type BError interface {
	GetStatus() int
	GetMessage() string
	GetError() error
	GetData() any
	Error() string
}

// Internal Server Error
type BErrorBase struct {
	Err     error  `json:"error"`
	Status  int    `json:"status_code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (v BErrorBase) GetStatus() int {
	return v.Status
}
func (v BErrorBase) GetMessage() string {
	return v.Message
}

func (v BErrorBase) GetData() any {
	return v.Data
}

func (v BErrorBase) GetError() error {
	return v.Err
}
func (v *BErrorBase) SetStatus(code int) *BErrorBase {
	v.Status = code
	return v
}
func (v *BErrorBase) SetError(err error) *BErrorBase {
	v.Err = err
	return v
}
func (v *BErrorBase) SetMessage(msg string) *BErrorBase {
	v.Message = msg
	return v
}

func (v *BErrorBase) With(data any) *BErrorBase {
	v.Data = data
	return v
}

func (v BErrorBase) Error() string {
	return v.Err.Error()
}

// Construct default Berror from error and string
func From(err error, msg string) *BErrorBase {
	if strings.HasSuffix(msg, ";") {
		msg = msg + err.Error()
	}
	return &BErrorBase{
		Message: msg,
		Err:     err,
		Status:  400,
	}
}

// Sometimes Developer needs to just send a message as error.
// this Function is dedicated for them.
// message would be something went wrong. but can be change via SetMessage()
func FromMsg(msg string) *BErrorBase {
	return &BErrorBase{
		Message: msg,
		Err:     errors.New("something went wrong"),
		Status:  400,
	}
}

// Bad Request error. extended from BErrorBase
// ByDefault Status code is 400
type BadRequest struct {
	BErrorBase
}

// Initiator of BadRequest Error
func NewBadRequest(err error, msg string) BadRequest {
	return BadRequest{
		BErrorBase{
			Err:     err,
			Status:  400,
			Message: msg,
		},
	}
}

// Error for unauthorized request.
// Status Code is 401
type AuthRequired struct {
	BErrorBase
}

// Initiator of AuthRequired
func NewAuthRequired(err error, msg string) *AuthRequired {
	return &AuthRequired{
		BErrorBase{
			Err:     err,
			Status:  http.StatusUnauthorized,
			Message: msg,
		},
	}
}

// Forbidden Request Error
// status code is 403
type ForbiddenRequest struct {
	BErrorBase
}

// Initiator of Forbidden Request Error
func NewForbiddenReq(err error, msg string) *ForbiddenRequest {
	return &ForbiddenRequest{
		BErrorBase{
			Err:     err,
			Status:  403,
			Message: msg,
		},
	}
}

// Content Not Found Error
// Status code si 404
type NotFound struct {
	BErrorBase
}

// Initiator of NotFound Error
func NewNotFound(err error, msg string) *NotFound {
	return &NotFound{
		BErrorBase{
			Err:     err,
			Status:  404,
			Message: msg,
		},
	}
}

// Content Not Found Error
// Status code si 406
type NotAcceptable struct {
	BErrorBase
}

// Initiator of NotFound Error
func NewNotAcceptable(err error, msg string) *NotAcceptable {
	return &NotAcceptable{
		BErrorBase{
			Err:     err,
			Status:  406,
			Message: msg,
		},
	}
}

// Error For internal Server Error
//
// By Default Status code is 500
type InternalError struct {
	BErrorBase
}

// InternalServerError - 500
//
// Initiator of InternalServer Error
func NewInternal(err error, msg string) *BadRequest {
	return &BadRequest{
		BErrorBase{
			Err:     err,
			Status:  500,
			Message: msg,
		},
	}
}

// Error For internal Server Error
// By Default Status code is 501
type NotImplemented struct {
	BErrorBase
}

// Initiator of InternalServer Error
func NewNotImplemented(err error, msg string) *NotImplemented {
	return &NotImplemented{
		BErrorBase{
			Err:     err,
			Status:  501,
			Message: msg,
		},
	}
}
