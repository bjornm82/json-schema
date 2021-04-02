package internal

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

const (
	PiiSchemaPath = "file://../schemas/pii.json"
)

type Pii []PiiObject

type PiiObject struct {
	Type      string
	Score     float64
	Language  string
	Anonymise map[string]interface{}
	Source    Marhaller
}

type Marhaller interface {
	Marshal() ([]byte, error)
}

type Replace struct {
	Type     string `json:"type"`
	NewValue string `json:"new_value"`
}

type Hash struct {
	Type     string `json:"type"`
	HashType string `json:"hash_type"`
}

type Redact struct {
	Type string `json:"type"`
}

type None struct {
	Type string `json:"type"`
}

func (n *None) Marshal() ([]byte, error) {
	return json.Marshal(n)
}

func (p *PiiObject) Unmarshal(b []byte) error {
	if err := json.Unmarshal(b, p); err != nil {
		return err
	}
	config := &mapstructure.DecoderConfig{
		TagName: "json",
	}

	switch p.Anonymise["type"] {
	case "mask":
		var m Mask
		config.Result = &m
		decoder, err := mapstructure.NewDecoder(config)
		if err != nil {
			return err
		}
		if err := decoder.Decode(p.Anonymise); err != nil {
			return err
		}

		p.Source = &m
	case "none":
		var n None
		config.Result = &n
		decoder, err := mapstructure.NewDecoder(config)
		if err != nil {
			return err
		}
		if err := decoder.Decode(p.Anonymise); err != nil {
			return err
		}

		p.Source = &n
	}

	return nil
}

func (p *Pii) Validate(byteObject []byte) (bool, error) {
	return validate(PiiSchemaPath, byteObject)
}
