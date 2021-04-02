package validation

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

type Result struct {
	*gojsonschema.Result
}

type Response []Error

func (r *Result) Parse() (Response, error) {

	var response = Response{}

	for _, v := range r.Errors() {
		e := Error{}
		e.Field = v.Field()
		e.Type = v.Type()
		e.Context = v.Context().String()
		e.Code = "0001" // TODO: add code levels
		e.Description = v.Description()
		e.Value = v.Value()
		e.Details = v.Details()
		response = append(response, e)
	}

	if len(response) == 0 {
		return response, nil
	}

	return response, errors.New("validation errors")
}
