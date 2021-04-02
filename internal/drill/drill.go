package drill

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

const (
	drillSchemaPath = "file://./schemas/drill.json"
)

type Drill struct {
	Name   string  `json:"name"`
	Sql    string  `json:"sql"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Precision  int    `json:"precision"`
	IsNullable bool   `json:"isNullable"`
}

func (d *Drill) Unmarshal(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *Drill) Marshal(data []byte) ([]byte, error) {
	return json.Marshal(d)
}

func (d *Drill) Validate() (bool, error) {
	ls := gojsonschema.NewReferenceLoader(drillSchemaPath)
	ld := gojsonschema.NewGoLoader(d)

	result, err := gojsonschema.Validate(ls, ld)
	if err != nil {
		return false, errors.Wrap(err, "unable to validate schema")
	}

	if !result.Valid() {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}

		return false, errors.New("document is not valid")
	}

	// TODO: Not validating, should be probably with Drill by just querying the API
	// _, err = sqlparser.Parse(d.Sql)
	// if err != nil {
	// 	return false, errors.Wrap(err, "not able to parse query")
	// }

	return true, nil
}
