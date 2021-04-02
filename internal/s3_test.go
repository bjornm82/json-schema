package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s3Object = `{
	"type": "s3",
	"name": "ABCDEFGH",
	"path": "/path/to/file",
	"bucket": "ABCDEFGHIJKLMNOPQRSTU",
	"key": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"secret": "ABCDEFGHIJKLM",
	"endpoint": "http://docker.for.mac.localhost"
  }`

func TestS3Validate(t *testing.T) {
	s := S3{}

	ok, err := s.Validate([]byte(s3Object))
	if err != nil {
		t.Error(err)
	}
	assert.True(t, ok)
}

func TestS3Unmarshal(t *testing.T) {
	s := S3{}
	err := json.Unmarshal([]byte(s3Object), &s)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "s3", s.Type)
	assert.Equal(t, "ABCDEFGH", s.Name)
	assert.Equal(t, "/path/to/file", s.Path)
	assert.Equal(t, "ABCDEFGHIJKLMNOPQRSTU", s.Bucket)
	assert.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", s.Key)
	assert.Equal(t, "ABCDEFGHIJKLM", s.Secret)
	assert.Equal(t, "http://docker.for.mac.localhost", s.Endpoint)
}

func TestS3Marshal(t *testing.T) {
	s := S3{}
	err := json.Unmarshal([]byte(s3Object), &s)
	if err != nil {
		t.Error(err)
	}
	b, err := s.Marshal()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, `{"type":"s3","name":"ABCDEFGH","path":"/path/to/file","bucket":"ABCDEFGHIJKLMNOPQRSTU","key":"ABCDEFGHIJKLMNOPQRSTUVWXYZ","secret":"ABCDEFGHIJKLM","endpoint":"http://docker.for.mac.localhost"}`, string(b))
}
