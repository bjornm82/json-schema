package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fields = `[
		{
			"name": "serverProtocol",
			"type": ["null", "string", "int"]
		},
		{
			"name": "serverProtocol",
			"type": ["string", "boolean", "double"]
		},
		{
			"name": "serverHash",
			"type": ["null", {
				"type": "fixed",
				"name": "MD5",
				"size": 16
			},{
				"type": "enum",
				"name": "HandshakeMatch",
				"symbols": ["BOTH", 1, "NONE"]
			}, "string", {
				"type": "map",
				"values": "bytes"
			}]
		},
		{
			"name": "meta",
			"type": [{
				"type": "map",
				"values": "bytes"
			}, {
				"type": "map",
				"values": "other"
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
				"name": "emails",
				"symbols": ["perosnal", "company", "other"]
			}
		},
		{
			"name": "match",
			"type": {
				"type": "enum",
				"name": "HandshakeMatch",
				"symbols": [2, 1, 4]
			}
		},
		{
			"name": "email",
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
			"type": "float"
		},
		{
			"name": "serverProtocol",
			"type": "double"
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
	]`

func TestSchema(t *testing.T) {
	Schema(fields)
	assert.Equal(t, "", "")
}
