package exceptions

// ErrorInterface custom error interface
// @Description:
type ErrorInterface interface {
	error
	ErrorInspect() string
	GetErrorType() string
	GetFunctionName() string
	GetLineNo() int
	GetFilename() string
	GetData() string
}
