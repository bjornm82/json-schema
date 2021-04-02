package drill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sqlQuery = "SELECT CAST('id' AS VARCHAR(20)) AS `id`, CAST('age' AS bigint) AS `age`, CAST('sex' AS varchar(20)) AS `sex`, CAST('region' AS varchar(20)) AS `region`, CAST('income' AS DOUBLE) AS `income` FROM `s3.DEFAULT`.`bank-data-with-headers.csvh`"

var testView = `{
	"name": "bank-data-func-test",
	"sql": "` + sqlQuery + `",
	"fields": [
	   {
		"name": "id",
		"type": "VARCHAR",
		"precision": 20,
		"isNullable": true
	  },
	  {
		"name": "age",
		"type": "BIGINT",
		"isNullable": true
	  },
	  {
		"name": "income",
		"type": "DOUBLE",
		"isNullable": true
	  }
	]
  }`

func TestValidation(t *testing.T) {
	d := Drill{}
	err := d.Unmarshal([]byte(testView))
	if err != nil {
		t.Error(t, err)
	}

	ok, err := d.Validate()
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestUpsert(t *testing.T) {
	cl := NewClient("localhost", 8047, false)
	d := Drill{}
	err := d.Unmarshal([]byte(testView))
	if err != nil {
		t.Error(t, err)
	}

	err = cl.UpsertView(d.Name, d.Sql)
	if err != nil {
		t.Error(t, err)
	}
}
func TestDelete(t *testing.T) {
	cl := NewClient("localhost", 8047, false)
	d := Drill{}
	err := d.Unmarshal([]byte(testView))
	if err != nil {
		t.Error(t, err)
	}

	err = cl.DeleteView(d.Name)
	if err != nil {
		t.Error(t, err)
	}
}
