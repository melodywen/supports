package exceptions

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// BaseError
// @Description: base error
type BaseError struct {
	errorType string // err type

	skipStack    int    // skip runtime stack num
	functionName string // The function name of the runtime
	filename     string // The absolute path to the runtime
	lineno       int    // The line number of the file at runtime

	message string // message
	data    string // Json serialize string
}

// NewBaseError
// @Description: construct
// @param errorType
// @param message
// @param data
// @param skipStack
// @return err
func NewBaseError(errorType string, message string, data string, skipStack int) (err *BaseError) {
	err = &BaseError{
		errorType: errorType,
		skipStack: skipStack,
		message:   message,
		data:      data,
	}
	if skipStack != 0 {
		err.initRuntimeEnv(skipStack)
	}
	return err
}

// initRuntimeEnv
// @Description:Initialize the operating environment
// @receiver err
// @param skipStack
func (err *BaseError) initRuntimeEnv(skipStack int) {
	pc, filename, lineno, _ := runtime.Caller(skipStack)
	err.filename = filename
	err.lineno = lineno
	functionName := runtime.FuncForPC(pc).Name()
	err.functionName = functionName
}

// Error
// @Description: Output error message
// @receiver err
// @return string
func (err *BaseError) Error() string {
	return err.message
}

// ErrorInspect
// @Description: Output error inspect message
// @receiver err
// @return string
func (err *BaseError) ErrorInspect() string {
	inspect := map[string]interface{}{
		"error_type":    err.errorType,
		"function_name": err.functionName,
		"filename":      err.filename,
		"lineno":        err.lineno,
		"skip_stack":    err.skipStack,
		"message":       err.message,
		"data":          err.data,
	}
	response, e := json.MarshalIndent(inspect, "", "  ")
	if e != nil {
		fmt.Printf("Error: %s", e)
		panic(e)
	}
	return string(response)
}

// GetErrorType
// @Description: get error type
// @receiver err
// @return string
func (err *BaseError) GetErrorType() string {
	return err.errorType
}

// GetFunctionName
// @Description: get function name
// @receiver err
// @return string
func (err *BaseError) GetFunctionName() string {
	return err.functionName
}

// GetLineNo
// @Description: get line no
// @receiver err
// @return int
func (err *BaseError) GetLineNo() int {
	return err.lineno
}

// GetFilename
// @Description: get filename
// @receiver err
// @return string
func (err *BaseError) GetFilename() string {
	return err.filename
}

// GetData
// @Description: get data
// @receiver err
// @return string
func (err *BaseError) GetData() string {
	return err.data
}
