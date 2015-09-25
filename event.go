package onfido

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
)

var ResourceType = struct {
	Report string
	Check  string
}{
	Report: "report",
	Check:  "check",
}

var EventType = struct {
	Completion string
	Withdrawal string
	InProgress string
}{
	Completion: "completion",
	Withdrawal: "withdrawal",
	InProgress: "in_progress",
}

type Object struct {
	ID          string    `json:"id"`
	Status      string    `json:"status"`
	CompletedAt time.Time `json:"completed_at"`
	Href        string    `json:"href"`
}

type Event struct {
	Payload struct {
		ResourceType string `json:"resource_type"`
		Action       string `json:"action"`
		Object       Object `json:"object"`
	} `json:"payload"`
}

func (c *Client) UnmarshalEvent(sig string, body io.Reader) (*Event, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	if c.WebhookToken != "" {
		if !compareMAC(b, []byte(sig), []byte(c.WebhookToken)) {
			return nil, ErrBadSignature
		}
	}

	var e Event
	err = json.Unmarshal(b, &e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

// compareMAC reports whether expectedMAC is a valid HMAC tag for message.
func compareMAC(message, expectedMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	messageMAC := make([]byte, hex.EncodedLen(mac.Size()))
	hex.Encode(messageMAC, mac.Sum(nil))
	// fmt.Println(string(expectedMAC), string(messageMAC))
	return subtle.ConstantTimeCompare(messageMAC, expectedMAC) == 1
}
