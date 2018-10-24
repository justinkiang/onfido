package onfido

import "time"

var CheckType = struct {
	Express  string
	Standard string
}{
	Express:  "express",
	Standard: "standard",
}

type ReportRequest struct {
	Name string `json:"name"`
}
type CheckRequest struct {
	Type    string          `json:"type"`
	Reports []ReportRequest `json:"reports"`
	Tags    []string        `json:"tags"`
}

type Check struct {
	ID        string    `json:"id"`
	Result    string    `json:"result"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	Href      string    `json:"href"`
	Reports   []*Report `json:"reports"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Check) ReportForName(name string) *Report {
	for _, v := range c.Reports {
		if v.Name == name {
			return v
		}
	}
	return nil
}

type checksResponse struct {
	Checks []*Check `json:"checks"`
}

func NewCheckRequest(checkType string, reports ...string) *CheckRequest {
	cr := &CheckRequest{
		Type:    CheckType.Express,
		Reports: make([]ReportRequest, len(reports)),
	}
	for i := 0; i < len(reports); i++ {
		cr.Reports[i].Name = reports[i]
	}
	return cr
}
