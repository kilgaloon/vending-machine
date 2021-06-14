package proto

import (
	"encoding/json"
	"io"
)

// Update body content
type Update struct {
	Password string `json:"password"`
	Role     string `json:"role"`
}

//FromJSON converts data to json struct
func (u *Update) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}
