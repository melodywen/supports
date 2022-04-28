package str

import (
	"fmt"
	"testing"
)

func BenchmarkE(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "<html></html>",
			args: args{value: "<html></html>"},
			want: "&lt;html&gt;&lt;/html&gt;",
		}, {
			name: "&lt;html&gt;&lt;/html&gt;",
			args: args{value: "<html></html>"},
			want: "&lt;html&gt;&lt;/html&gt;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := E(tt.args.value); got != tt.want {
				t.Errorf("E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRegularReplaceArray(t *testing.B) {
	type args struct {
		pattern      string
		replacements []string
		subject      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "The event will take place between :start and :end",
			args: args{
				pattern:      ":[a-z_]+",
				replacements: []string{"8:30", "9:00"},
				subject:      "The event will take place between :start and :end",
			},
			want: "The event will take place between 8:30 and 9:00",
		}, {
			name: "The event will take place between :start and :end",
			args: args{
				pattern:      ":[a-z_]+",
				replacements: []string{"8:30"},
				subject:      "The event will take place between :start and :end",
			},
			want: "The event will take place between 8:30 and :end",
		}, {
			name: "The event will take place between :start and :end",
			args: args{
				pattern:      ":[a-z_]+",
				replacements: []string{"8:30", "9:00", "10:00"},
				subject:      "The event will take place between :start and :end",
			},
			want: "The event will take place between 8:30 and 9:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := RegularReplaceArray(tt.args.pattern, tt.args.replacements, tt.args.subject); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("RegularReplaceArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAfter(t *testing.B) {
	type args struct {
		subject string
		search  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "This isa",
			},
			want: "This is my name,This is you name",
		}, {
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "",
			},
			want: "This is my name,This is you name",
		}, {
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "This is",
			},
			want: "my name,This is you name",
		}, {
			name: "This is my name",
			args: args{
				subject: "aa This is my name,This is you name",
				search:  "This is",
			},
			want: "my name,This is you name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := After(tt.args.subject, tt.args.search); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("After() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAfterLast(t *testing.B) {
	type args struct {
		subject string
		search  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "This isa",
			},
			want: "This is my name,This is you name",
		}, {
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "",
			},
			want: "This is my name,This is you name",
		}, {
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "This is",
			},
			want: "you name",
		}, {
			name: "This is my name",
			args: args{
				subject: "aa This is my name,This is you name",
				search:  "This is",
			},
			want: "you name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := AfterLast(tt.args.subject, tt.args.search); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("AfterLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_before(t *testing.B) {
	type args struct {
		subject string
		search  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "my name",
			},
			want: "This is ",
		}, {
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "",
			},
			want: "This is my name,This is you name",
		}, {
			name: "This is my name,This is you name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "name",
			},
			want: "This is my ",
		}, {
			name: "This is my name,This is you name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "namea",
			},
			want: "This is my name,This is you name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Before(tt.args.subject, tt.args.search); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("before() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_beforeLast(t *testing.B) {
	type args struct {
		subject string
		search  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "This isa",
			},
			want: "This is my name,This is you name",
		}, {
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "",
			},
			want: "This is my name,This is you name",
		}, {
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "my name",
			},
			want: "This is ",
		}, {
			name: "This is my name",
			args: args{
				subject: "This is my name,This is you name",
				search:  "name",
			},
			want: "This is my name,This is you ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := BeforeLast(tt.args.subject, tt.args.search); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("beforeLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBetween(t *testing.B) {
	type args struct {
		subject string
		from    string
		to      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name,This is you name",
			args: args{
				subject: "This is my name,This is you name",
				from:    "name",
				to:      "name",
			},
			want: ",This is you ",
		}, {
			name: "This is my name,This is you name",
			args: args{
				subject: "This is my name,This is you name",
				from:    "This",
				to:      "is",
			},
			want: "is my name,This ",
		}, {
			name: "This is my name,This is you name",
			args: args{
				subject: "This is my name,This is you name",
				from:    "This",
				to:      "",
			},
			want: "This is my name,This is you name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Between(tt.args.subject, tt.args.from, tt.args.to); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("Between() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBetweenFirst(t *testing.B) {
	type args struct {
		subject string
		from    string
		to      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name,This is you name",
			args: args{
				subject: "This is my name,This is you name",
				from:    "This",
				to:      "name",
			},
			want: "is my ",
		}, {
			name: "This is my name,This is you name",
			args: args{
				subject: "This is my name,This is you name",
				from:    "This",
				to:      "",
			},
			want: "This is my name,This is you name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := BetweenFirst(tt.args.subject, tt.args.from, tt.args.to); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("BetweenFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_replace(t *testing.B) {
	type args struct {
		search  string
		replace string
		subject string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name,This is you name",
			args: args{
				search:  "is",
				replace: "are",
				subject: "This is my name,This is you name",
			},
			want: "Thare are my name,Thare are you name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Replace(tt.args.search, tt.args.replace, tt.args.subject); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_replaceFirst(t *testing.B) {
	type args struct {
		search  string
		replace string
		subject string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name,This is you name",
			args: args{
				search:  "is",
				replace: "are",
				subject: "This is my name,This is you name",
			},
			want: "Thare is my name,This is you name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := ReplaceFirst(tt.args.search, tt.args.replace, tt.args.subject); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("replaceFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_replaceArray(t *testing.B) {
	type args struct {
		search  string
		replace []string
		subject string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name,This is you name",
			args: args{
				search: "is",
				replace: []string{
					"are1", "are2", "are3",
				},
				subject: "This is my name,This is you name",
			},
			want: "Thare1 are2 my name,Thare3 is you name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := ReplaceArray(tt.args.search, tt.args.replace, tt.args.subject); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("replaceArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_replaceLast(t *testing.B) {
	type args struct {
		search  string
		replace string
		subject string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "This is my name,This is you name",
			args: args{
				search:  "is",
				replace: "are",
				subject: "This is my name,This is you name",
			},
			want: "This is my name,This are you name",
		}, {
			name: "This is my name,This is you name",
			args: args{
				search:  "",
				replace: "are",
				subject: "This is my name,This is you name",
			},
			want: "This is my name,This is you name",
		}, {
			name: "This is my name,This is you name",
			args: args{
				search:  "aa",
				replace: "are",
				subject: "This is my name,This is you name",
			},
			want: "This is my name,This is you name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := ReplaceLast(tt.args.search, tt.args.replace, tt.args.subject); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("replaceLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_substr(t *testing.B) {
	type args struct {
		subject string
		start   int
		length  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "The Tool Framework",
			args: args{
				subject: "The Tool Framework",
				start:   4,
				length:  4,
			},
			want: "Tool",
		}, {
			name: "The Tool Framework",
			args: args{
				subject: "The Tool Framework",
				start:   19,
				length:  4,
			},
			want: "",
		}, {
			name: "The Tool Framework",
			args: args{
				subject: "The Tool Framework",
				start:   4,
				length:  40,
			},
			want: "Tool Framework",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Substr(tt.args.subject, tt.args.start, tt.args.length); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("substr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_substrCount(t *testing.B) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "If you like ice cream, you will like snow cones.",
			args: args{
				haystack: "If you like ice cream, you will like snow cones.",
				needle:   "like",
			},
			want: 2,
		}, {
			name: "If you like ice cream, you will like snow cones.",
			args: args{
				haystack: "If you like ice cream, you will like snow cones.",
				needle:   "",
			},
			want: 0,
		}, {
			name: "If you like ice cream, you will like snow cones.",
			args: args{
				haystack: "",
				needle:   "like",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := SubstrCount(tt.args.haystack, tt.args.needle); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("substrCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSubstrReplace(t *testing.B) {
	type args struct {
		subject string
		replace string
		offset  int
		length  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1300",
			args: args{
				subject: "1300",
				replace: ":",
				offset:  2,
				length:  1,
			},
			want: "13:0",
		}, {
			name: "1300",
			args: args{
				subject: "1300",
				replace: ":",
				offset:  2,
				length:  0,
			},
			want: "13:00",
		}, {
			name: "1300",
			args: args{
				subject: "1300",
				replace: ":",
				offset:  -100,
				length:  2,
			},
			want: ":00",
		}, {
			name: "1300",
			args: args{
				subject: "1300",
				replace: ":",
				offset:  200,
				length:  2,
			},
			want: "1300:",
		}, {
			name: "1300",
			args: args{
				subject: "1300",
				replace: ":",
				offset:  -100,
				length:  -100,
			},
			want: ":1300",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := SubstrReplace(tt.args.subject, tt.args.replace, tt.args.offset, tt.args.length); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("SubstrReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkContains(t *testing.B) {
	type args struct {
		haystack   string
		needles    string
		ignoreCase bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "This is my name",
			args: args{
				haystack:   "This is my name",
				needles:    "my",
				ignoreCase: false,
			},
			want: true,
		}, {
			name: "This is my name",
			args: args{
				haystack:   "This is my name",
				needles:    "mya",
				ignoreCase: false,
			},
			want: false,
		}, {
			name: "This is my name",
			args: args{
				haystack:   "This is my name",
				needles:    "MY",
				ignoreCase: true,
			},
			want: true,
		}, {
			name: "This is my name",
			args: args{
				haystack:   "This is my name",
				needles:    "",
				ignoreCase: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Contains(tt.args.haystack, tt.args.needles, tt.args.ignoreCase); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsAny(t *testing.B) {
	type args struct {
		haystack   string
		needles    []string
		ignoreCase bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "This is my name",
			args: args{
				haystack:   "This is my name",
				needles:    []string{"MY", "name"},
				ignoreCase: false,
			},
			want: true,
		}, {
			name: "This is my name",
			args: args{
				haystack:   "This is my name",
				needles:    []string{"Mine", "Name"},
				ignoreCase: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := ContainsAny(tt.args.haystack, tt.args.needles, tt.args.ignoreCase); got != tt.want {
				t.Errorf("ContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsAll(t *testing.B) {
	type args struct {
		haystack   string
		needles    []string
		ignoreCase bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "This is my name",
			args: args{
				haystack:   "This is my name",
				needles:    []string{"MY", "name"},
				ignoreCase: true,
			},
			want: true,
		}, {
			name: "This is my name",
			args: args{
				haystack:   "This is my name",
				needles:    []string{"Mine", "Name"},
				ignoreCase: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := ContainsAll(tt.args.haystack, tt.args.needles, tt.args.ignoreCase); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkEndsWith(t *testing.B) {
	type args struct {
		haystack string
		needles  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "This is my name",
			args: args{
				haystack: "This is my name",
				needles:  "name",
			},
			want: true,
		}, {
			name: "This is my name",
			args: args{
				haystack: "This is my name",
				needles:  "name tom",
			},
			want: false,
		}, {
			name: "This is my name",
			args: args{
				haystack: "This is my name",
				needles:  "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := EndsWith(tt.args.haystack, tt.args.needles); got != tt.want {
				t.Errorf("EndsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLength(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "my name",
			args: args{value: "my name"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Length(tt.args.value); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_is(t *testing.B) {
	type args struct {
		pattern string
		value   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "http://c.biancheng.net/php",
			args: args{
				pattern: "http://c.biancheng.net/*",
				value:   "http://c.biancheng.net/php",
			},
			want: true,
		}, {
			name: "foobar",
			args: args{
				pattern: "baz*",
				value:   "foobar",
			},
			want: false,
		}, {
			name: "foobar",
			args: args{
				pattern: "foobar",
				value:   "foobar",
			},
			want: true,
		}, {
			name: "foobar",
			args: args{
				pattern: "",
				value:   "foobar",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Is(tt.args.pattern, tt.args.value); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFinish(t *testing.B) {
	type args struct {
		value string
		cap   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "this/string/",
			args: args{
				value: "this/string.*/ab",
				cap:   "ab",
			},
			want: "this/string.*/ab",
		}, {
			name: "this/string/",
			args: args{
				value: "this/string.*/",
				cap:   "ab",
			},
			want: "this/string.*/ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Finish(tt.args.value, tt.args.cap); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("Finish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLimit(t *testing.B) {
	type args struct {
		value string
		limit int
		end   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "limit too long",
			args: args{
				value: "golang",
				limit: 2,
				end:   "...",
			}, want: "go...",
		}, {
			name: "limit too long",
			args: args{
				value: "golang",
				limit: 20,
				end:   "...",
			}, want: "golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Limit(tt.args.value, tt.args.limit, tt.args.end); got != tt.want {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLower(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GoLang",
			args: args{value: "GoLang"},
			want: "golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Lower(tt.args.value); got != tt.want {
				t.Errorf("Lower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkUpper(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GoLang",
			args: args{value: "GoLang"},
			want: "GOLANG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Upper(tt.args.value); got != tt.want {
				t.Errorf("Upper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPadLeft(t *testing.B) {
	type args struct {
		value  string
		length int
		pad    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tom",
			args: args{
				value:  "tom",
				length: 2,
				pad:    "-",
			},
			want: "tom",
		}, {
			name: "tom",
			args: args{
				value:  "tom",
				length: 10,
				pad:    "-",
			},
			want: "-------tom",
		}, {
			name: "tom",
			args: args{
				value:  "tom",
				length: 10,
				pad:    "-=",
			},
			want: "-=-=-=-tom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := PadLeft(tt.args.value, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPadRight(t *testing.B) {
	type args struct {
		value  string
		length int
		pad    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tom",
			args: args{
				value:  "tom",
				length: 2,
				pad:    "-",
			},
			want: "tom",
		}, {
			name: "tom",
			args: args{
				value:  "tom",
				length: 10,
				pad:    "-",
			},
			want: "tom-------",
		}, {
			name: "tom",
			args: args{
				value:  "tom",
				length: 10,
				pad:    "-=",
			},
			want: "tom-=-=-=-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := PadRight(tt.args.value, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPadBoth(t *testing.B) {
	type args struct {
		value  string
		length int
		pad    string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
	}{
		{
			name: "tom",
			args: args{
				value:  "tom",
				length: 2,
				pad:    "-",
			},
			wantResponse: "tom",
		}, {
			name: "tom",
			args: args{
				value:  "tom",
				length: 10,
				pad:    "-",
			},
			wantResponse: "---tom----",
		}, {
			name: "tom",
			args: args{
				value:  "tom",
				length: 10,
				pad:    "-=",
			},
			wantResponse: "=-=tom-=-=",
		}, {
			name: "tom",
			args: args{
				value:  "tom",
				length: 8,
				pad:    "-=",
			},
			wantResponse: "-=tom-=-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if gotResponse := PadBoth(tt.args.value, tt.args.length, tt.args.pad); gotResponse != tt.wantResponse {
				t.Errorf("PadBoth() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func BenchmarkStart(t *testing.B) {
	type args struct {
		value  string
		prefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "this/string",
			args: args{
				value:  "this/string",
				prefix: "/",
			},
			want: "/this/string",
		}, {
			name: "this/string",
			args: args{
				value:  "/this/string",
				prefix: "/",
			},
			want: "/this/string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Start(tt.args.value, tt.args.prefix); got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStartsWith(t *testing.B) {
	type args struct {
		haystack string
		needles  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "This is my name",
			args: args{
				haystack: "This is my name",
				needles:  "is",
			},
			want: false,
		}, {
			name: "This is my name",
			args: args{
				haystack: "This is my name",
				needles:  "This",
			},
			want: true,
		}, {
			name: "This is my name",
			args: args{
				haystack: "This is my name",
				needles:  "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := StartsWith(tt.args.haystack, tt.args.needles); got != tt.want {
				t.Errorf("StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsAscii(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Tom",
			args: args{value: "Tom"},
			want: true,
		}, {
			name: "陈",
			args: args{value: "陈"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := IsAscii(tt.args.value); got != tt.want {
				t.Errorf("IsAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRandom(t *testing.B) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "16",
			args: args{length: 16},
			want: 16,
		}, {
			name: "16",
			args: args{length: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Random(tt.args.length); len(got) != tt.want {
				t.Errorf("Random() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkUcFirst(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "golang orm",
			args: args{value: "golang orm"},
			want: "Golang orm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := UcFirst(tt.args.value); got != tt.want {
				t.Errorf("UcFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkLcFirst(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "golang orm",
			args: args{value: "Golang orm"},
			want: "golang orm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := LcFirst(tt.args.value); got != tt.want {
				t.Errorf("UcFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReverse(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "null",
			args: args{value: ""},
			want: "",
		}, {
			name: "hello world",
			args: args{value: "hello world"},
			want: "dlrow olleh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Reverse(tt.args.value); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSwap(t *testing.B) {
	type args struct {
		swapMap map[string]string
		subject string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Tacos are great!",
			args: args{
				swapMap: map[string]string{
					"Tacos": "Burritos",
					"great": "fantastic",
				},
				subject: "Tacos are great!"},
			want: "Burritos are fantastic!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Swap(tt.args.swapMap, tt.args.subject); got != tt.want {
				t.Errorf("Swap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRemove(t *testing.B) {
	type args struct {
		search        string
		subject       string
		caseSensitive bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "this is my name",
			args: args{
				search:        "",
				subject:       "this is my name",
				caseSensitive: false,
			},
			want: "this is my name",
		}, {
			name: "this is my name",
			args: args{
				search:        "",
				subject:       "this is my name",
				caseSensitive: true,
			},
			want: "this is my name",
		}, {
			name: "this is my name",
			args: args{
				search:        "Is",
				subject:       "this Is my name",
				caseSensitive: true,
			},
			want: "this  my name",
		}, {
			name: "this is my name",
			args: args{
				search:        "Is",
				subject:       "thisAIsMyName",
				caseSensitive: false,
			},
			want: "thAmyname",
		}, {
			name: "this is my name",
			args: args{
				search:        "Is",
				subject:       "this is my name",
				caseSensitive: false,
			},
			want: "th  my name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Remove(tt.args.search, tt.args.subject, tt.args.caseSensitive); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReplaceOfArraySearch(t *testing.B) {
	type args struct {
		search  []string
		replace string
		subject string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
	}{
		{
			name: "The event will take place between :start and :end",
			args: args{
				search:  []string{"en", "a"},
				replace: "*",
				subject: "The event will take place between :start and :end",
			},
			wantResponse: "The ev*t will t*ke pl*ce betwe* :st*rt *nd :*d",
		}, {
			name: "The event will take place between :start and :end",
			args: args{
				search:  []string{},
				replace: "*",
				subject: "The event will take place between :start and :end",
			},
			wantResponse: "The event will take place between :start and :end",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if gotResponse := ReplaceOfArraySearch(tt.args.search, tt.args.replace, tt.args.subject); gotResponse != tt.wantResponse {
				t.Errorf("ReplaceOfArraySearch() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func BenchmarkStudly(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
	}{
		{
			name:         "golang_orm_pkg",
			args:         args{value: "golang_orm_pkg"},
			wantResponse: "GolangOrmPkg",
		}, {
			name:         "golang_orm_pkg",
			args:         args{value: ""},
			wantResponse: "",
		}, {
			name:         "golang",
			args:         args{value: "golang"},
			wantResponse: "Golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if gotResponse := Studly(tt.args.value); gotResponse != tt.wantResponse {
				t.Errorf("Studly() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func BenchmarkCamel(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
	}{
		{
			name:         "golang_orm_pkg",
			args:         args{value: "golang_orm_pkg"},
			wantResponse: "golangOrmPkg",
		}, {
			name:         "golang_orm_pkg",
			args:         args{value: ""},
			wantResponse: "",
		}, {
			name:         "golang",
			args:         args{value: "golang"},
			wantResponse: "golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if gotResponse := Camel(tt.args.value); gotResponse != tt.wantResponse {
				t.Errorf("Camel() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func BenchmarkHeadline(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
	}{
		{
			name:         "steve_jobs",
			args:         args{value: "steve_jobs"},
			wantResponse: "Steve Jobs",
		}, {
			name:         "steve_jobs",
			args:         args{value: ""},
			wantResponse: "",
		}, {
			name:         "EmailNotificationSent",
			args:         args{value: "EmailNotificationSent"},
			wantResponse: "Email Notification Sent",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if gotResponse := Headline(tt.args.value); gotResponse != tt.wantResponse {
				t.Errorf("Headline() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func BenchmarkSnake(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
	}{
		{
			name:         "golangOrmPkg",
			args:         args{value: "golangOrmPkg"},
			wantResponse: "golang_orm_pkg",
		}, {
			name:         "golang_orm_pkg",
			args:         args{value: ""},
			wantResponse: "",
		}, {
			name:         "golang",
			args:         args{value: "golang"},
			wantResponse: "golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if gotResponse := Snake(tt.args.value); gotResponse != tt.wantResponse {
				t.Errorf("Snake() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func BenchmarkKebab(t *testing.B) {
	type args struct {
		value string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
	}{
		{
			name:         "golangOrmPkg",
			args:         args{value: "golangOrmPkg"},
			wantResponse: "golang-orm-pkg",
		}, {
			name:         "golang_orm_pkg",
			args:         args{value: ""},
			wantResponse: "",
		}, {
			name:         "golang",
			args:         args{value: "golang"},
			wantResponse: "golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if gotResponse := Kebab(tt.args.value); gotResponse != tt.wantResponse {
				t.Errorf("Kebab() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func BenchmarkMask(t *testing.B) {
	type args struct {
		subject   string
		character string
		index     int
		length    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tel",
			args: args{
				subject:   "186777123456",
				character: "*",
				index:     3,
				length:    3,
			},
			want: "186***123456",
		}, {
			name: "tel",
			args: args{
				subject:   "186777123456",
				character: "*",
				index:     30,
				length:    -30,
			},
			want: "186777123456",
		}, {
			name: "tel",
			args: args{
				subject:   "186777123456",
				character: "*",
				index:     3,
				length:    -3,
			},
			want: "186******456",
		}, {
			name: "tel",
			args: args{
				subject:   "186777123456",
				character: "*",
				index:     -30,
				length:    -30,
			},
			want: "186777123456",
		}, {
			name: "tel",
			args: args{
				subject:   "186777123456",
				character: "*",
				index:     -30,
				length:    30,
			},
			want: "************",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := Mask(tt.args.subject, tt.args.character, tt.args.index, tt.args.length); got != tt.want {
				t.Errorf("Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkExcerpt(t *testing.B) {
	type args struct {
		text    string
		phrase  string
		options map[string]string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
	}{
		{
			name: "This is my name",
			args: args{
				text:    "This is my name",
				phrase:  "you",
				options: map[string]string{"radius": "3", "omission": "(...)"},
			},
			wantResponse: "",
		}, {
			name: "This is my name",
			args: args{
				text:    "This is my name",
				phrase:  "my",
				options: map[string]string{"radius": "4", "omission": "...."},
			},
			wantResponse: ".... is my nam....",
		}, {
			name: "This is my name",
			args: args{
				text:    "This is my name",
				phrase:  "is",
				options: map[string]string{"radius": "3"},
			},
			wantResponse: "This is...",
		}, {
			name: "This is my name",
			args: args{
				text:    "This is my name",
				phrase:  "na",
				options: map[string]string{"radius": "3"},
			},
			wantResponse: "...my name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if gotResponse := Excerpt(tt.args.text, tt.args.phrase, tt.args.options); gotResponse != tt.wantResponse {
				t.Errorf("Excerpt() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
