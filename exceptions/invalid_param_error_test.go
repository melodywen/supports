package exceptions

import (
	"github.com/melodywen/supports/str"
	"testing"
)

func BenchmarkNewInvalidParamErrorCustomStack(t *testing.B) {
	type args struct {
		message   string
		data      string
		skipStack int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not runtime",
			args: args{
				message:   "please input data",
				data:      "data",
				skipStack: 0,
			},
			want: "please input data",
		}, {
			name: "not runtime",
			args: args{
				message:   "please input data",
				data:      "data",
				skipStack: 3,
			},
			want: "please input data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			defer func() {
				if err := recover(); err != nil {
					e, _ := err.(ErrorInterface)
					if got := e.Error(); got != tt.want {
						t.Errorf("Error() = %v, want %v", got, tt.want)
					}
				}
			}()
			panic(NewInvalidParamErrorCustomStack(tt.args.message, tt.args.data, tt.args.skipStack))
		})
	}
}
func BenchmarkErrorInspect(t *testing.B) {
	type args struct {
		message   string
		data      string
		skipStack int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not runtime",
			args: args{
				message:   "please input data",
				data:      "data",
				skipStack: 0,
			},
			want: `{
  "data": "data",
  "error_type": "invalid_param",
  "filename": "",
  "function_name": "",
  "lineno": 0,
  "message": "please input data",
  "skip_stack": 0
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			defer func() {
				if err := recover(); err != nil {
					e, _ := err.(ErrorInterface)
					if got := e.ErrorInspect(); got != tt.want {
						t.Errorf("Error() = %v, want %v", got, tt.want)
					}
				}
			}()
			panic(NewInvalidParamErrorCustomStack(tt.args.message, tt.args.data, tt.args.skipStack))
		})
	}
}
func BenchmarkGetErrorType(t *testing.B) {
	type args struct {
		message   string
		data      string
		skipStack int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not runtime",
			args: args{
				message:   "please input data",
				data:      "data",
				skipStack: 0,
			},
			want: "invalid_param",
		}, {
			name: "not runtime",
			args: args{
				message:   "please input data",
				data:      "data",
				skipStack: 3,
			},
			want: "invalid_param",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			defer func() {
				if err := recover(); err != nil {
					e, _ := err.(ErrorInterface)
					if got := e.GetErrorType(); got != tt.want {
						t.Errorf("Error() = %v, want %v", got, tt.want)
					}
				}
			}()
			panic(NewInvalidParamErrorCustomStack(tt.args.message, tt.args.data, tt.args.skipStack))
		})
	}
}
func BenchmarkGetFunctionName(t *testing.B) {
	type args struct {
		message   string
		data      string
		skipStack int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not runtime skip 1",
			args: args{
				message:   "please input data",
				data:      "data",
				skipStack: 3,
			},
			want: "BenchmarkGetFunctionName.func1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			defer func() {
				if err := recover(); err != nil {
					e, _ := err.(ErrorInterface)
					got := e.GetFunctionName()
					if got == "" || !str.EndsWith(got, tt.want) {
						t.Errorf("Error() = %v, want %v", got, tt.want)
					}
				}
			}()
			panic(NewInvalidParamErrorCustomStack(tt.args.message, tt.args.data, tt.args.skipStack))
		})
	}
}
func BenchmarkGetLineNo(t *testing.B) {
	type args struct {
		message   string
		data      string
		skipStack int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not runtime skip 1",
			args: args{
				message:   "please input data",
				data:      "data",
				skipStack: 3,
			},
			want: 205,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			defer func() {
				if err := recover(); err != nil {
					e, _ := err.(ErrorInterface)
					got := e.GetLineNo()
					if got < tt.want {
						t.Errorf("Error() = %v, want %v", got, tt.want)
					}
				}
			}()
			panic(NewInvalidParamErrorCustomStack(tt.args.message, tt.args.data, tt.args.skipStack))
		})
	}
}
func BenchmarkGetFilename(t *testing.B) {
	type args struct {
		message   string
		data      string
		skipStack int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not runtime skip 1",
			args: args{
				message:   "please input data",
				data:      "data",
				skipStack: 3,
			},
			want: "invalid_param_error_test.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			defer func() {
				if err := recover(); err != nil {
					e, _ := err.(ErrorInterface)
					got := e.GetFilename()
					if got == "" || !str.EndsWith(got, tt.want) {
						t.Errorf("Error() = %v, want %v", got, tt.want)
					}
				}
			}()
			panic(NewInvalidParamErrorCustomStack(tt.args.message, tt.args.data, tt.args.skipStack))
		})
	}
}
func BenchmarkGetData(t *testing.B) {
	type args struct {
		message   string
		data      string
		skipStack int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not runtime skip 1",
			args: args{
				message:   "please input data",
				data:      "data",
				skipStack: 3,
			},
			want: "data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			defer func() {
				if err := recover(); err != nil {
					e, _ := err.(ErrorInterface)
					got := e.GetData()
					if got != tt.want {
						t.Errorf("Error() = %v, want %v", got, tt.want)
					}
				}
			}()
			panic(NewInvalidParamErrorCustomStack(tt.args.message, tt.args.data, tt.args.skipStack))
		})
	}
}

func BenchmarkNewInvalidParamError(t *testing.B) {
	type args struct {
		message string
		data    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "not runtime skip 1",
			args: args{
				message: "please input data",
			},
			want: "please input data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			defer func() {
				if err := recover(); err != nil {
					e, _ := err.(ErrorInterface)
					if got := e.Error(); got != tt.want {
						t.Errorf("Error() = %v, want %v", got, tt.want)
					}
				}
			}()
			panic(NewInvalidParamError(tt.args.message))
		})
	}
}
