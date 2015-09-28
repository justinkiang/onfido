package onfido

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/tylerb/is"
)

func TestEventUnmarshalEvent(t *testing.T) {
	is := is.New(t)

	body := []byte(`{"payload":{"resource_type":"report","action":"completed","object":{"id":"50e65977-496f-464a-be73-836def7bea6c","status":"completed","completed_at":"2006-01-02 15:04:05 MST","href":"https://api.onfido.com/checks/dd29776f-26fa-452f-90f1-9c8ce4fcec3e/reports/50e65977-496f-464a-be73-836def7bea6c"}}}`)

	c := New("none")
	c.WebhookToken = "kCxrT1iqSqlxf1OeQi2On5M-1fA6xnr3"

	r, err := http.NewRequest("POST", "/webhook", bytes.NewReader(body))
	is.NotErr(err)
	r.Header.Set("X-Signature", "8fe7130c7e193b5049daafb507ad246047c69bc4")

	event, err := c.UnmarshalEvent(r)
	is.NotErr(err)
	is.Equal(event.Payload.ResourceType, "report")
	is.Equal(event.Payload.Action, "completed")
	is.NotZero(event.Payload.Object.ID)
	is.Equal(event.Payload.Object.Status, "completed")
	tc, err := time.Parse(eventTimeFormat, "2006-01-02 15:04:05 MST")
	is.NotErr(err)
	is.Equal(event.Payload.Object.CompletedAt, tc)
	is.Equal(event.Payload.Object.Href, "https://api.onfido.com/checks/dd29776f-26fa-452f-90f1-9c8ce4fcec3e/reports/50e65977-496f-464a-be73-836def7bea6c")
}
