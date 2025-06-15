package dto

type Response[T any] struct {
	Status int       `json:"status"`
	Body   T         `json:"body"`
	Errors *[]*Error `json:"errors,omitempty"`
}

func ResponseBuilder[T any](status int, body T) *Response[T] {
	return &Response[T]{
		Status: status,
		Body:   body,
		Errors: &[]*Error{},
	}
}

func (r *Response[T]) AddError(err *Error) *Response[T] {
	*r.Errors = append(*r.Errors, err)
	return r
}

func (r *Response[T]) SetStatus(status int) *Response[T] {
	r.Status = status
	return r
}

func (r *Response[T]) SetBody(body T) *Response[T] {
	r.Body = body
	return r
}

func (r *Response[T]) AddErrors(errors *[]*Error) *Response[T] {
	*r.Errors = append(*r.Errors, *errors...)
	return r
}

func (r *Response[T]) SetErrors(errors *[]*Error) *Response[T] {
	r.Errors = errors
	return r
}
