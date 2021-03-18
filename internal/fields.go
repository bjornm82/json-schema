package internal

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

type Avro struct {
}

func (a *Avro) Unmarshal(b []byte) error {
	return json.Unmarshal(b, a)
}

func Schema(fields string) {
	schemaLoader := gojsonschema.NewReferenceLoader("file:///Users/bmooijekind/Projects/go/src/github.com/bjornm82/json-schema/schemas/fieldSchema.json")
	documentLoader := gojsonschema.NewStringLoader(fields)

	i, err := schemaLoader.LoadJSON()
	log.Println(err)
	log.Println(i)

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
