package onfido

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func newError(err string) error {
	return fmt.Errorf("onfido: %s", err)
}

var (
	ErrBadSignature = newError("webhook signature validation failed")
)

type APIError struct {
	Err struct {
		ID      string                 `json:"id"`
		Fields  map[string]interface{} `json:"fields"`
		Message string                 `json:"message"`
		Type    string                 `json:"type"`
	} `json:"error"`
}

func (a *APIError) Error() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(a.Err.Message)
	if len(a.Err.Fields) == 0 {
		return buf.String()
	}
	buf.WriteString(" ~ ")
	b, err := json.Marshal(a.Err.Fields)
	if err == nil {
		buf.Write(b)
	}
	return buf.String()
}
