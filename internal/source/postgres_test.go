package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	pg := Postgres{
		Name:     "name",
		Username: "Username",
		Password: "Password",
		Database: "Database",
		Table:    "Table",
		Host:     "Host",
		Port:     1234,
	}

	resp, err := pg.Validate()

	assert.Error(t, err)
	assert.Equal(t, "required", resp[0].Type)
	assert.Equal(t, "0001", resp[0].Code)
	assert.Equal(t, "does not equal format", resp[1].Type)
	assert.Equal(t, "actual", "expected")
}
