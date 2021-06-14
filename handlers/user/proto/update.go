package proto

import (
	"encoding/json"
	"io"
)

// Update body content
type Deposit struct {
	Amount uint64 `json:"amount"`
}

//FromJSON converts data to json struct
func (d *Deposit) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}
