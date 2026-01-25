package dto

import (
	"net/http"
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

func TestMapWithoutBody(t *testing.T) {

	message := "Operación exitosa"
	errors := &[]*Error{{Message: "Error de prueba", Field: "Test"}}

	type args[T any] struct {
		original *Response[T]
	}

	type testCase[T any, R any] struct {
		name string
		args args[T]
		want *Response[R]
	}

	tests := []testCase[string, any]{
		{
			name: "Debe retornar nil si el response original es nil",
			args: args[string]{
				original: nil,
			},
			want: nil,
		},
		{
			name: "Debe cambiar el body a nil (zero value) y conservar el status 200",
			args: args[string]{
				original: &Response[string]{
					Status: http.StatusOK,
					Body:   "Contenido que será ignorado",
					Errors: &[]*Error{},
				},
			},
			want: &Response[any]{
				Status: http.StatusOK,
				Body:   nil,
				Errors: &[]*Error{},
			},
		},
		{
			name: "Debe conservar Mensaje, Errores y Status, pero eliminar Body",
			args: args[string]{
				original: &Response[string]{
					Status:  http.StatusBadRequest,
					Message: &message,
					Body:    "Contenido sensible",
					Errors:  errors,
				},
			},
			want: &Response[any]{
				Status:  http.StatusBadRequest,
				Message: &message,
				Body:    nil,
				Errors:  errors,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapWithoutBody[string, any](tt.args.original); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToNilBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapWithBody(t *testing.T) {
	message := "Operación exitosa"
	errors := &[]*Error{{Message: "Error de prueba", Field: "Test"}}

	type args[T any, R any] struct {
		original *Response[T]
		newBody  R
	}

	type testCase[T any, R any] struct {
		name string
		args args[T, R]
		want *Response[R]
	}

	tests := []testCase[string, int]{
		{
			name: "Debe retornar nil si el response original es nil",
			args: args[string, int]{
				original: nil,
				newBody:  100,
			},
			want: nil,
		},
		{
			name: "Debe reemplazar el body con el nuevo valor (123) y conservar status 200",
			args: args[string, int]{
				original: &Response[string]{
					Status: http.StatusOK,
					Body:   "Texto original",
					Errors: &[]*Error{},
				},
				newBody: 123,
			},
			want: &Response[int]{
				Status: http.StatusOK,
				Body:   123,
				Errors: &[]*Error{},
			},
		},
		{
			name: "Debe conservar Mensaje, Errores y Status al cambiar el body",
			args: args[string, int]{
				original: &Response[string]{
					Status:  http.StatusBadRequest,
					Message: &message,
					Body:    "Error texto",
					Errors:  errors,
				},
				newBody: -1,
			},
			want: &Response[int]{
				Status:  http.StatusBadRequest,
				Message: &message,
				Body:    -1,
				Errors:  errors,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapWithBody(tt.args.original, tt.args.newBody); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapWithBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
