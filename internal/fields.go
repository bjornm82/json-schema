package internal

import (
	"encoding/json"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Avro struct {
}

func (a *Avro) Unmarshal(b []byte) error {
	return json.Unmarshal(b, a)
}

func Schema(fields string) {
	schemaLoader := gojsonschema.NewReferenceLoader("file:///Users/bmooijekind/Projects/go/src/github.com/bjornm82/json-schema/schemas/field.json")
	documentLoader := gojsonschema.NewStringLoader(fields)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
