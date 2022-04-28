package utils

import "testing"

func BenchmarkDump(t *testing.B) {
	type args struct {
		value any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "dump-bool",
			args: args{value: true},
		}, {
			name: "dump-string",
			args: args{value: "hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			Dump(tt.args.value)
		})
	}
}
