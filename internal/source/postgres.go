package source

import (
	"encoding/json"

	"github.com/bjornm82/json-schema/internal/validation"
	"github.com/pkg/errors"
)

const (
	schemaPath  = "file://../../schemas/postgres.json"
	msgValidate = "unable to validate schema"
)

type Postgres struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Table    string `json:"table"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func (p *Postgres) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Postgres) Unmarshal(b []byte) error {
	return json.Unmarshal(b, p)
}

func (p *Postgres) Validate() (validation.Response, error) {
	e := validation.Error{}
	b, err := p.Marshal()
	if err != nil {
		return validation.Response{}, errors.Wrap(err, "unable to marshal object")
	}
	result, err := e.Validate(schemaPath, b)
	if err != nil {
		return validation.Response{}, errors.Wrap(err, msgValidate)
	}

	return result.Parse()
}
