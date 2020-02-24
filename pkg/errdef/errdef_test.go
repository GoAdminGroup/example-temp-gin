package errdef

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_err_def_NewCode(t *testing.T) {
	errCode := NewErr(ErrParams)
	err := errCode.Add("params id not found")
	if err == nil {
		t.Errorf("errCode.Add err %v", err)
	}
	// do
	code := errCode.Code
	msg := errCode.Msg
	errorInfo := errCode.Error()
	// verify
	assert.Equal(t, 10003, code)
	assert.Equal(t, "Error params. Error params. params id not found", msg)
	assert.Equal(t, "Err - code: 10003, message: Error params. Error params. params id not found, error: Error params.", errorInfo)
}

func Test_err_def_NEW(t *testing.T) {
	// mock
	// do
	paramsErr := fmt.Errorf("error params")
	err := New(ErrParams, paramsErr)
	code := err.Code
	errInfo := err.Error()
	// verify
	assert.Equal(t, 10003, code)
	assert.Equal(t, "Err - code: 10003, message: Error params., error: error params", errInfo)
}

func Test_err_def_add(t *testing.T) {
	// mock
	paramsErr := fmt.Errorf("error params")
	addMessage := "add message"
	errCode := New(ErrParams, paramsErr)

	err := errCode.Add(addMessage)
	if err == nil {
		t.Errorf("errCode.Add err %v", err)
	}

	// do
	code := errCode.Code
	msg := errCode.Msg
	// verify
	assert.Equal(t, 10003, code)
	assert.Equal(t, "Error params. error params add message", msg)

}
func Test_err_def_IsErrUserNotFound(t *testing.T) {
	// mock
	userNotFound := NewErr(ErrUserNotFound)
	bindErr := NewErr(ErrBind)

	// do
	foundOne := IsErrUserNotFound(userNotFound)
	foundTwo := IsErrUserNotFound(bindErr)

	// verify
	assert.Equal(t, true, foundOne)
	assert.Equal(t, false, foundTwo)
}
