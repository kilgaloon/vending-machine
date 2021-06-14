package proto

import (
	"encoding/json"
	"io"
)

// Buy body content
type Buy struct {
	ProductID uint   `json:"product_id"`
	Amount    uint64 `json:amount`
	
}

//FromJSON converts data to json struct
func (b *Buy) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}
