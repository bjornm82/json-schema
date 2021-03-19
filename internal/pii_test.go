package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// DATE_TIME
// CREDIT_CARD
// DOMAIN_NAME
// IP_ADDRESS
// EMAIL_ADDRESS
// IBAN_CODE
// LOCATION
// LATITUDE
// LONGITUDE
// PERSON
// PHONE_NUMBER
// US_BANK_NUMBER
// US_DRIVER_LICENCE
// US_PASSPORT
// NL_DRIVER_LICENCE
// NL_PASSPORT

// NONE
// HASH
// - hash_type: ["sha256", "sha512", "md5"]
// REPLACE
// - new_value: string
// REDACT
// MASK
// - chars_to_mask: integer
// - masking_char: string
// - from_end: bool

var piiFields = `[{
	"type": "CREDIT_CARD",
	"score": 0.85,
	"language": "nl",
	"anonymise": {
		"type": "hash",
		"hash_type": "md5"
	}
}, {
	"type": "PHONE_NUMBER",
	"score": 0.85,
	"language": "nl",
	"anonymise": "none"
}, {
	"type": "IBAN_CODE",
	"score": 0.85,
	"language": "nl",
	"anonymise": "redact"
}, {
	"type": "IP_ADDRESS",
	"score": 0.82,
	"language": "nl",
	"anonymise": {
		"type": "replace",
		"new_value": "<IP_ADDRESS>"
	}
}, {
	"type": "US_DRIVER_LICENCE",
	"score": 0.82,
	"language": "en",
	"anonymise": {
		"type": "mask",
		"chars_to_mask": 8,
		"masking_char": "*"
	}
}, {
	"type": "NL_PASSPORT",
	"score": 0.82,
	"language": "nl",
	"anonymise": {
		"type": "mask",
		"chars_to_mask": 8,
		"masking_char": "*",
		"from_end": true
	}
}]`

func TestSchemaPii(t *testing.T) {
	SchemaPii(piiFields)
	assert.Equal(t, "", "1")
}
