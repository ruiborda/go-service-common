package dto

import (
	"reflect"
	"testing"
)

func TestResponse_AddErrors(t *testing.T) {
	type args struct {
		errors *[]*Error
	}
	type testCase[T any] struct {
		name string
		r    Response[T]
		args args
		want *Response[T]
	}
	tests := []testCase[string]{
		{
			name: "Add nil errors",
			r:    *ResponseBuilder[string](200, ""),
			args: args{errors: nil},
			want: &Response[string]{
				Status: 200,
				Body:   "",
				Errors: &[]*Error{},
			},
		},
		{
			name: "Add empty error list",
			r:    *ResponseBuilder[string](200, ""),
			args: args{errors: &[]*Error{}},
			want: &Response[string]{
				Status: 200,
				Body:   "",
				Errors: &[]*Error{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.AddErrors(tt.args.errors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}
