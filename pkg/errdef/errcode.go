package errdef

import "net/http"

// Common errors, each error must has HttpStatus!
var (
	OK                  = &ErrDef{Code: 0, Msg: "OK", HttpStatus: http.StatusOK}
	InternalServerError = &ErrDef{Code: 10001, Msg: "Internal server error.", HttpStatus: http.StatusForbidden}
	ErrBind             = &ErrDef{
		Code:       10002,
		Msg:        "Error occurred while binding the request body to the struct.",
		HttpStatus: http.StatusBadRequest,
	}
	ErrParams = &ErrDef{Code: 10003, Msg: "Error params.", HttpStatus: http.StatusBadRequest}
	ErrParse  = &ErrDef{Code: 10004, Msg: "Error parse.", HttpStatus: http.StatusBadRequest}

	// database errors
	ErrValidation = &ErrDef{Code: 20001, Msg: "Validation failed.", HttpStatus: http.StatusUnauthorized}
	ErrDatabase   = &ErrDef{Code: 20002, Msg: "Database error.", HttpStatus: http.StatusForbidden}
	ErrToken      = &ErrDef{Code: 20003,
		Msg:        "Error occurred while signing the JSON web token.",
		HttpStatus: http.StatusUnauthorized,
	}

	// user errors
	ErrEncrypt = &ErrDef{Code: 20101,
		Msg:        "Error occurred while encrypting the user password.",
		HttpStatus: http.StatusBadRequest,
	}
	ErrUserNotFound      = &ErrDef{Code: 20102, Msg: "The user was not found.", HttpStatus: http.StatusUnauthorized}
	ErrTokenInvalid      = &ErrDef{Code: 20103, Msg: "The token was invalid.", HttpStatus: http.StatusUnauthorized}
	ErrPasswordIncorrect = &ErrDef{Code: 20104, Msg: "The password was incorrect.", HttpStatus: http.StatusUnauthorized}

	// other errors
)
