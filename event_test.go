package onfido

import (
	"bytes"
	"testing"
	"time"

	"github.com/tylerb/is"
)

func TestEventUnmarshalEvent(t *testing.T) {
	is := is.New(t)

	body := []byte(`{"payload":{"resource_type":"report","action":"completed","object":{"id":"50e65977-496f-464a-be73-836def7bea6c","status":"completed","completed_at":"2015-09-25T16:02:36+00:00","href":"https://api.onfido.com/checks/dd29776f-26fa-452f-90f1-9c8ce4fcec3e/reports/50e65977-496f-464a-be73-836def7bea6c"}}}`)

	c := New("none")
	c.WebhookToken = "kCxrT1iqSqlxf1OeQi2On5M-1fA6xnr3"

	event, err := c.UnmarshalEvent("f8b06c5d4bdbe5639ad6f7a5a1f6a7b7523ff945", bytes.NewReader(body))
	is.NotErr(err)
	is.Equal(event.Payload.ResourceType, "report")
	is.Equal(event.Payload.Action, "completed")
	is.NotZero(event.Payload.Object.ID)
	is.Equal(event.Payload.Object.Status, "completed")
	tc, err := time.Parse(time.RFC3339, "2015-09-25T16:02:36+00:00")
	is.NotErr(err)
	is.Equal(event.Payload.Object.CompletedAt, tc)
	is.Equal(event.Payload.Object.Href, "https://api.onfido.com/checks/dd29776f-26fa-452f-90f1-9c8ce4fcec3e/reports/50e65977-496f-464a-be73-836def7bea6c")
}
