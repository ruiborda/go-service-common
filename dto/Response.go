package dto

import (
	"net/http"
)

type ResponseErrorProvider interface {
	GetErrors() *[]*Error
}

type Response[T any] struct {
	Status  int       `json:"status"`
	Message *string   `json:"message,omitempty"`
	Body    T         `json:"body"`
	Errors  *[]*Error `json:"errors,omitempty"`
}

func NewResponse[T any]() *Response[T] {
	return &Response[T]{
		Status: http.StatusProcessing,
		Errors: &[]*Error{},
	}
}

func ResponseBuilder[T any](status int, body T) *Response[T] {
	return &Response[T]{
		Status: status,
		Body:   body,
		Errors: &[]*Error{},
	}
}

func (r *Response[T]) GetErrors() *[]*Error {
	return r.Errors
}

func (r *Response[T]) MergeErrors(other ResponseErrorProvider) *Response[T] {
	if other == nil || other.GetErrors() == nil {
		return r
	}
	return r.AddErrors(other.GetErrors())
}

func (r *Response[T]) AddError(err *Error) *Response[T] {
	if err == nil {
		return r
	}
	if r.Errors == nil {
		r.Errors = &[]*Error{}
	}
	*r.Errors = append(*r.Errors, err)
	return r
}

func (r *Response[T]) AddErrors(errors *[]*Error) *Response[T] {
	if errors == nil {
		return r
	}
	if r.Errors == nil {
		r.Errors = &[]*Error{}
	}
	*r.Errors = append(*r.Errors, *errors...)
	return r
}

func (r *Response[T]) SetStatus(status int) *Response[T] {
	r.Status = status
	return r
}

func (r *Response[T]) SetMessage(message string) *Response[T] {
	r.Message = &message
	return r
}

func (r *Response[T]) SetBody(body T) *Response[T] {
	r.Body = body
	return r
}

func (r *Response[T]) SetErrors(errors *[]*Error) *Response[T] {
	if errors == nil {
		return r
	}
	r.Errors = errors
	return r
}

func (r *Response[T]) HasErrors() bool {
	return r.Errors != nil && len(*r.Errors) > 0
}

func (r *Response[T]) IsInfo() bool {
	return r.Status >= 100 && r.Status < 200
}

func (r *Response[T]) IsOK() bool {
	return r.Status >= 200 && r.Status < 300
}

func (r *Response[T]) IsRedirect() bool {
	return r.Status >= 300 && r.Status < 400
}

func (r *Response[T]) IsClientError() bool {
	return r.Status >= 400 && r.Status < 500
}

func (r *Response[T]) IsServerError() bool {
	return r.Status >= 500 && r.Status < 600
}

func Map[T any, R any](original *Response[T], transform func(T) R) *Response[R] {
	if original == nil {
		return nil
	}
	return &Response[R]{
		Status:  original.Status,
		Message: original.Message,
		Errors:  original.Errors,
		Body:    transform(original.Body),
	}
}
