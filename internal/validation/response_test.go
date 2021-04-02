package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const anyOfTestdata = `[1,2.5,3,1.5]`

func TestParse(t *testing.T) {

	e := Error{}
	result, err := e.Validate("file://./schema.json", []byte(anyOfTestdata))
	if err != nil {
		t.Errorf("json schema result errors: %s", err.Error())
	}

	l, err := result.Parse()
	if err != nil {
		t.Errorf("json schema parse errors: %s", err.Error())
	}

	assert.Equal(t, "(root)", l[0].Field)
	assert.Equal(t, "invalid_type", l[0].Type)
	assert.Equal(t, "(root)", l[0].Context)
	assert.Equal(t, "0001", l[0].Code)
	assert.Equal(t, "Invalid type. Expected: string, given: array", l[0].Description)
	assert.Equal(t, `["1", "2.5", "3", "1.5"]`, l[0].Value)
	assert.Equal(t, map[string]interface{}(map[string]interface{}{"context": "(root)", "expected": "string", "field": "(root)", "given": "array"}), l[0].Details)

	assert.Error(t, err)

	assert.Equal(t, "expected", "actual")
}
