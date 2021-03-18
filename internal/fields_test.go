package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fields = `[
		{
			"name": "serverProtocol",
			"type": ["null", "null", "string"]
		},
		{
			"name": "serverProtocol",
			"type": ["string", "string"]
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
				"values": "bytes"
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
	]`

func TestSchema(t *testing.T) {
	Schema(fields)
	assert.Equal(t, "", "1")
}
