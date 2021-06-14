package proto

import (
	"encoding/json"
	"io"
)

// Create body content
type Create struct {
	AmountAvailable uint64 `json:"amount_available"`
	Cost            uint64 `json:"cost"`
	ProductName     string `json:"product_name"`
}

//FromJSON converts data to json struct
func (c *Create) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}
