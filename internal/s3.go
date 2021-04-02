package internal

import (
	"encoding/json"
)

const (
	s3schemaPath = "file://../schemas/s3.json"
)

type S3 struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Bucket   string `json:"bucket"`
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	Endpoint string `json:"endpoint"`
}

func (s *S3) Validate(byteObject []byte) (bool, error) {
	return validate(s3schemaPath, byteObject)
}

func (s *S3) Marshal() ([]byte, error) {
	return json.Marshal(s)
}
