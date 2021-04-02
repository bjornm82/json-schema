package internal

import "encoding/json"

const (
	maskSchemaPath = "file://../schemas/mask.json"
)

type Mask struct {
	Type        string `json:"type"`
	CharsToMask int    `json:"chars_to_mask"`
	MaskingChar string `json:"masking_char"`
	FromEnd     bool   `json:"from_end"`
}

func (m *Mask) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Mask) Validate(b []byte) (bool, error) {
	return validate(maskSchemaPath, b)
}
