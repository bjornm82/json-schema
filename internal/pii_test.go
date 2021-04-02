package internal

import (
	"errors"
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

var pii = `[{
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
	"anonymise": {
		"type": "none"
	}
}, {
	"type": "IBAN_CODE",
	"score": 0.85,
	"language": "nl",
	"anonymise": {
		"type": "redact"
	}
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

var piiObjectMask = `{
	"type": "NL_PASSPORT",
	"score": 0.82,
	"language": "nl",
	"anonymise": {
		"type": "mask",
		"chars_to_mask": 8,
		"masking_char": "*",
		"from_end": true
	}
}`

var piiObjectNone = `{
	"type": "PHONE_NUMBER",
	"score": 0.85,
	"language": "nl",
	"anonymise": {
		"type": "none"
	}
}`

func TestValidatePii(t *testing.T) {
	p := Pii{}
	ok, err := p.Validate([]byte(pii))
	if err != nil {
		t.Error(err)
	}
	assert.True(t, ok)
}

func TestUnmarshalPiiMask(t *testing.T) {
	p := PiiObject{}
	err := p.Unmarshal([]byte(piiObjectMask))
	if err != nil {
		t.Error(err)
	}

	m := p.Source.(*Mask)

	assert.Equal(t, "mask", m.Type)
	assert.Equal(t, 8, m.CharsToMask)
	assert.Equal(t, "*", m.MaskingChar)
	assert.Equal(t, true, m.FromEnd)

	assert.Equal(t, "mask", p.Anonymise["type"])
	// Why we shouldn't use this
	assert.Equal(t, float64(8), p.Anonymise["chars_to_mask"])
	assert.Equal(t, "*", p.Anonymise["masking_char"])
	assert.Equal(t, true, p.Anonymise["from_end"])
}

func TestUnmarshalPiiNone(t *testing.T) {
	p := PiiObject{}
	err := p.Unmarshal([]byte(piiObjectNone))
	if err != nil {
		t.Error(err)
	}

	switch ta := p.Source.(type) {
	default:
		t.Error(t, errors.New("not supposed to be default"))
	case *None:
		assert.Equal(t, &None{Type: "none"}, ta)
	}
}
