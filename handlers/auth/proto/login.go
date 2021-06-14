package proto

import (
	"encoding/json"
	"io"
)

// Login body content
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//FromJSON converts data to json struct
func (l *Login) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(l)
}
