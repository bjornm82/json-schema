package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/xeipuuv/gojsonschema"
)

var postgres = `{
	"data": {
		"type": "bounded",
		"processing": "event",
		"format": {
			"type": "xml",
			"version": 1.1
		}
	},
	"config": {
		"type": "postgres",
		"name": "postgres",
		"username": "postgres",
		"password": "postgres",
		"database": "postgres",
		"table": "postgres",
		"host": "http://postgres.com",
		"port": 1234
	},
	"fields": [{
			"name": "serverProtocol",
			"type": ["null", "string"]
		},
		{
			"name": "serverProtocol",
			"type": ["string"]
		},
		{
			"name": "serverHash",
			"type": ["null", {
				"type": "fixed",
				"name": "MD5",
				"size": 16
			}, {
				"type": "enum",
				"name": "HandshakeMatch",
				"symbols": ["BOTH", "CLIENT", "NONE"]
			}]
		},
		{
			"name": "meta",
			"type": [{
				"type": "map",
				"values": "bytes"
			}, {
				"type": "map",
				"values": "bytes2"
			}]
		},
		{
			"name": "serverHash",
			"type": {
				"type": "fixed",
				"name": "MD5",
				"size": 16
			}
		},
		{
			"name": "match",
			"type": {
				"type": "enum",
				"name": "HandshakeMatch",
				"symbols": ["BOTH", "CLIENT", "NONE"]
			}
		},
		{
			"name": "serverProtocol",
			"type": "string"
		},
		{
			"name": "serverProtocol",
			"type": "boolean"
		},
		{
			"name": "serverProtocol",
			"type": "null"
		},
		{
			"name": "serverProtocol",
			"type": {
				"type": "fixed",
				"name": "MD5",
				"size": 16
			}
		},
		{
			"name": "serverProtocol",
			"type": "int"
		},
		{
			"name": "serverProtocol",
			"type": "long"
		},
		{
			"name": "serverProtocol2",
			"type": {
				"type": "map",
				"values": "MD5"
			}
		},
		{
			"name": "serverProtocol",
			"type": "string"
		}
	]
}`

var s3 = `{
	  "config":
		  {
			"type": "s3",
			"name": "s3",
			"path": "s3",
			"bucket": "s3",
			"key": "s3",
			"secret": "s3"
		}
	}`

var http = `{
		"config":
			{
				"type": "http",
				"name": "http",
				"host": "http",
				"path": "http",
				"query": "http"
		  }
	  }`

type Evaluator interface {
	evaluate() bool
}

type Storage struct {
	Storage map[string]interface{} `json:"storage"`
	Source  Evaluator
}

type S3 struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

func (s *S3) evaluate() bool {
	return true
}

type Postgres struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Table    string `json:"table"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func (p *Postgres) evaluate() bool {
	return true
}

type Http struct {
	Name  string `json:"name"`
	Host  string `json:"host"`
	Path  string `json:"path"`
	Query string `json:"query"`
}

func (h *Http) evaluate() bool {
	return true
}

func main() {
	schemaLoader := gojsonschema.NewReferenceLoader("file:///Users/bmooijekind/Projects/go/src/github.com/bjornm82/json-schema/schemas/baseSchema.json")
	documentLoader := gojsonschema.NewStringLoader(postgres)

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

	var m = Storage{}
	if err := json.Unmarshal([]byte(postgres), &m); err != nil {
		panic(err)
	}

	switch m.Storage["type"] {
	case "postgres":
		var p Postgres
		if err := mapstructure.Decode(m.Storage, &p); err != nil {
			panic(err)
		}
		m.Source = &p
		fmt.Printf("%T %+v\n", p, p)
	case "http":
		var h Http
		if err := mapstructure.Decode(m.Storage, &h); err != nil {
			panic(err)
		}
		m.Source = &h
		fmt.Printf("%T %+v\n", h, h)
	case "s3":
		var s S3
		if err := mapstructure.Decode(m.Storage, &s); err != nil {
			panic(err)
		}
		m.Source = &s
		fmt.Printf("%T %+v\n", s, s)
	}
}
