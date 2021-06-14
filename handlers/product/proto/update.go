package proto

import (
	"encoding/json"
	"io"
)

// Update body content
type Update struct {
	AmountAvailable uint64 `json:"amount_available"`
	Cost            uint64 `json:"cost"`
}

//FromJSON converts data to json struct
func (u *Update) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}
