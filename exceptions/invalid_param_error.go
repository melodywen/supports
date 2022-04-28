package exceptions

// InvalidParamError
// @Description:
type InvalidParamError struct {
	BaseError
}

var invalidParamErrorTypeName = "invalid_param"

// NewInvalidParamError
// @Description: invalid param error construct
// @return *InvalidParamError
func NewInvalidParamError(message string) *InvalidParamError {
	err := NewBaseError(invalidParamErrorTypeName, message, "", 3)
	return &InvalidParamError{BaseError: *err}
}

// NewInvalidParamErrorWithData
// @Description: invalid param error construct with data
// @param message
// @param data
// @return *InvalidParamError
func NewInvalidParamErrorWithData(message string, data string) *InvalidParamError {
	err := NewBaseError(invalidParamErrorTypeName, message, data, 3)
	return &InvalidParamError{BaseError: *err}
}

// NewInvalidParamErrorCustomStack
// @Description: invalid param error construct ,and input skip stack
// @param message
// @param data
// @param skipStack
// @return *InvalidParamError
func NewInvalidParamErrorCustomStack(message string, data string, skipStack int) *InvalidParamError {
	err := NewBaseError(invalidParamErrorTypeName, message, data, skipStack)
	return &InvalidParamError{BaseError: *err}
}
