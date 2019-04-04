package onfido

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jmcvetta/napping"
    "image"
    "bytes"
    "image/jpeg"
)

var (
	applicantBaseURL = "https://api.onfido.com/v2/applicants"
	checkBaseURL     = "https://api.onfido.com/v2/applicants"
	reportBaseURL    = "https://api.onfido.com/v2/checks"
	sdkToeknBaseURL  = "https://api.onfido.com/v2/sdk_token"
	livePhotoBaseURL = "https://api.onfido.com/v2/live_photos"
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
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println("failed unmarshalling s", err)
	}
	t, err := time.Parse(dateFormat, s)
	*d = Date(t)
	if err != nil {
		t, err = time.Parse(time.RFC3339, s)
		*d = Date(t)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*d).Format(dateFormat))), nil
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

func (c *Client) SDKToken(appID, referrer string) (SDKToken, error) {
	s := c.newSession()
	apiErr := APIError{}

	token := SDKToken{}
	payload := struct {
		ApplicantId string `json:"applicant_id"`
		Referrer    string `json:"referrer"`
	}{
		appID,
		referrer,
	}
	res, err := s.Post(sdkToeknBaseURL, payload, &token, &apiErr)
	if err != nil {
		fmt.Println(err.Error())
		return token, err
	}

	if res.Status() != 200 {
		return token, &apiErr
	}
	return token, nil

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

func (c *Client) GetLivePhotos(appID string) ([]LivePhoto, error){
    s := c.newSession()
    apiErr := APIError{}
    a := livePhotosResponse{}
    q := url.Values{}
    q.Add("applicant_id",appID)
    res, err := s.Get(livePhotoBaseURL, &q, &a, &apiErr)
    if err != nil {
        return nil, err
    }
    if res.Status() != 200 {
        return nil, &apiErr
    }
    return a.LivePhotos, nil
}
func (c *Client) GetFile(url string)(){
    httpClient := http.Client{}

    req,_ := http.NewRequest("GET",url,nil)
    req.Header.Set("Authorization", c.apiTokenHeader)

    response, err := httpClient.Do(req)
    fmt.Println(err)
    defer response.Body.Close()
    m, _, err := image.Decode(response.Body)

    fmt.Println(err)
    buf := new(bytes.Buffer)
    err = jpeg.Encode(buf, m, nil)

    fmt.Println(err)
    fmt.Println(buf.String())
    //send_s3 := buf.Bytes()
    //imgBase64Str := base64.StdEncoding.EncodeToString(response.Body)

}
func (c *Client) GetDocuments(appID string) ([]Document, error) {
    s := c.newSession()
    apiErr := APIError{}
    a := documentsResponse{}
    res, err := s.Get(assembleURL(checkBaseURL, appID, "documents"), &url.Values{}, &a, &apiErr)
    if err != nil {
        return nil, err
    }
    if res.Status() != 200 {
        return nil, &apiErr
    }
    return a.Documents, nil
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

func (c *Client) ReadReport(url string) (*Report, error) {
	s := c.newSession()
	apiErr := APIError{}
	report := Report{}
	res, err := s.Get(url, expandParam, &report, &apiErr)
	if err != nil {
		return nil, err
	}
	if res.Status() != 200 {
		return nil, &apiErr
	}
	return &report, nil
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
