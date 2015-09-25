package onfido

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jmcvetta/napping"
)

var (
	applicantBaseURL = "https://api.onfido.com/v1/applicants"
	checkBaseURL     = "https://api.onfido.com/v1/applicants"
	expandParam      = &url.Values{}
)

const (
	dateFormat = "2006-01-02"
)

func init() {
	expandParam.Add("expand", "reports")
}

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) error {
	var err error
	t, err := time.Parse(dateFormat, string(data))
	*d = Date(t)
	if err != nil {
		*d = Date(time.Time{})
	}
	return nil
}

// Client provides methods for interacting with the onfido API
type Client struct {
	apiToken       string
	apiTokenHeader string
	WebhookToken   string
}

func New(apiToken string) *Client {
	return &Client{
		apiToken:       apiToken,
		apiTokenHeader: "Token token=" + apiToken,
	}
}

func (c *Client) CreateApplicant(a *Applicant) (*Applicant, error) {
	s := c.newSession()
	apiErr := APIError{}
	res, err := s.Post(applicantBaseURL, a, a, &apiErr)
	if err != nil {
		return nil, err
	}

	if res.Status() != 201 {
		return nil, &apiErr
	}
	return a, nil
}
func (c *Client) ReadApplicant(id string) (*Applicant, error) {
	s := c.newSession()
	apiErr := APIError{}
	a := Applicant{}
	res, err := s.Get(assembleURL(applicantBaseURL, id), expandParam, &a, &apiErr)
	if err != nil {
		return nil, err
	}

	if res.Status() != 200 {
		return nil, &apiErr
	}
	return &a, nil
}
func (c *Client) ReadApplicants() ([]*Applicant, error) {
	s := c.newSession()
	apiErr := APIError{}
	a := applicantsResponse{}
	res, err := s.Get(applicantBaseURL, expandParam, &a, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, &apiErr
	}
	return a.Applicants, nil
}
func (c *Client) CreateCheck(appID string, cr *CheckRequest) (*Check, error) {
	s := c.newSession()
	apiErr := APIError{}
	check := Check{}
	res, err := s.Post(assembleURL(checkBaseURL, appID, "checks"), cr, &check, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 201 {
		return nil, &apiErr
	}
	return &check, nil
}
func (c *Client) ReadCheck(appID, checkID string) (*Check, error) {
	s := c.newSession()
	apiErr := APIError{}
	check := Check{}
	res, err := s.Get(assembleURL(checkBaseURL, appID, "checks", checkID), expandParam, &check, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, &apiErr
	}
	return &check, nil
}

func (c *Client) ReadChecks(appID string) ([]*Check, error) {
	s := c.newSession()
	apiErr := APIError{}
	checks := checksResponse{}
	res, err := s.Get(assembleURL(checkBaseURL, appID, "checks"), expandParam, &checks, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, &apiErr
	}
	return checks.Checks, nil
}

func (c *Client) newSession() *napping.Session {
	n := napping.Session{
		Header: &http.Header{},
	}
	n.Header.Set("Authorization", c.apiTokenHeader)
	return &n
}

func assembleURL(parts ...string) string {
	return strings.Join(parts, "/")
}
