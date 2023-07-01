package errors

import (
	"encoding/json"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

// ErrorCode predefined error codes
type ErrorCode uint32

const (
	// UnknownError ...
	UnknownError ErrorCode = iota
	// Unauthorized ...
	Unauthenticated
	// UserIDInvalid ...
	UserIDInvalid
	// InvalidEmailError  ...
	InvalidEmailError
	// InvalidPasswordError ...
	InvalidPasswordError
	// InvalidUsernameError ...
	InvalidUsernameError
	// ExpiredTokenError ...
	ExpiredTokenError
)

const (
	msgUnknownError         = "unknown error"
	msgUnauthenticated      = "unauthenticated access"
	msgUserIDInvalid        = "invalid user id"
	msgInvalidEmailError    = "invalid username , make sure you enter a valid email"
	msgInvalidUsernameError = "invalid username , make sure you enter a username of length of 4 or more characters"
	msgInvalidPasswordError = "invalid password , make sure you enter a correct password of length of 8 or more characters"
	msgExpiredTokenError    = "expired token"
)

type ErrorDesc struct {
	Message string
	Code    int32
}

var ErrorsList = map[ErrorCode]ErrorDesc{
	UnknownError:         {Message: msgUnknownError, Code: 520},
	Unauthenticated:      {Message: msgUnauthenticated, Code: http.StatusUnauthorized},
	UserIDInvalid:        {Message: msgUserIDInvalid, Code: http.StatusBadRequest},
	InvalidPasswordError: {Message: msgInvalidPasswordError, Code: http.StatusBadRequest},
	InvalidUsernameError: {Message: msgInvalidUsernameError, Code: http.StatusBadRequest},
	ExpiredTokenError:    {Message: msgExpiredTokenError, Code: 403},
}

func (err ErrorDesc) Error() string {
	return err.Message
}

func NewErrorDesc(msg string, code int32) ErrorDesc {
	return ErrorDesc{
		Message: msg,
		Code:    code,
	}
}

func NewSocketError(message string, code int64) *mvmPb.SocketMessage_ErrorMessage {
	return &mvmPb.SocketMessage_ErrorMessage{
		ErrorMessage: &mvmPb.ErrorMessage{
			Error:      message,
			StatusCode: code,
		},
	}
}

func NewError(message string, code int64) *mvmPb.ErrorMessage {
	return &mvmPb.ErrorMessage{
		Error:      message,
		StatusCode: code,
	}
}

func NewHTTPError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
