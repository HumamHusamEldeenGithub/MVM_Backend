package errors

import (
	"encoding/json"
	"fmt"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	msgInvalidEmailError    = "invalid username , make sure you enter a username of length of 4 or more characters"
	msgInvalidUsernameError = "invalid username , make sure you enter a username of length of 4 or more characters"
	msgInvalidPasswordError = "invalid password , make sure you enter a correct password of length of 8 or more characters"
	msgExpiredTokenError    = "expired token"
)

type errorDesc struct {
	msg      string
	grpcCode codes.Code
}

var errors = map[ErrorCode]*errorDesc{
	UnknownError:         {msg: msgUnknownError, grpcCode: codes.Unknown},
	Unauthenticated:      {msg: msgUnauthenticated, grpcCode: codes.PermissionDenied},
	UserIDInvalid:        {msg: msgUserIDInvalid, grpcCode: codes.InvalidArgument},
	InvalidPasswordError: {msg: msgInvalidPasswordError, grpcCode: codes.InvalidArgument},
	InvalidUsernameError: {msg: msgInvalidUsernameError, grpcCode: codes.InvalidArgument},
	ExpiredTokenError:    {msg: msgExpiredTokenError, grpcCode: codes.Code(401)},
}

// Errorf creates error with msg arguments for grpc response
func Errorf(code ErrorCode, a ...interface{}) error {
	if errors[code] == nil {
		return status.Error(codes.Unknown, errors[UnknownError].msg)
	}
	d := errors[code]
	m := d.msg
	if a != nil {
		m = fmt.Sprintf(m, a)
	}
	return status.Error(d.grpcCode, m)
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
