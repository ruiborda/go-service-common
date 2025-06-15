package dto

type Error struct {
	Message string      `json:"message"`
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
}

func ErrorBuilder(field string, message string) *Error {
	return &Error{
		Message: message,
		Field:   field,
		Value:   nil,
	}
}

func (e *Error) SetMessage(message string) *Error {
	e.Message = message
	return e
}

func (e *Error) SetField(field string) *Error {
	e.Field = field
	return e
}

func (e *Error) SetValue(value interface{}) *Error {
	e.Value = value
	return e
}
