package proto

import (
	"encoding/json"
	"io"
)

// Create body content
type Create struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

//FromJSON converts data to json struct
func (c *Create) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}
