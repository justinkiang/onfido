package onfido

import "time"

var IDNumberType = struct {
	Nino           string
	SSN            string
	DrivingLicense string
}{
	Nino:           "nino",
	SSN:            "ssn",
	DrivingLicense: "driving_license",
}

type Address struct {
	BuildingName   string `json:"building_name,omitempty"`
	BuildingNumber string `json:"building_number,omitempty"`
	FlatNumber     string `json:"flat_number,omitempty"`
	Street         string `json:"street"`
	SubStreet      string `json:"sub_street,omitempty"`
	Town           string `json:"town"`
	State          string `json:"state,omitempty"`
	Postcode       string `json:"postcode"`
	Country        string `json:"country"`
	StartDate      *Date  `json:"start_date"`
	EndDate        *Date  `json:"end_date,omitempty"`
}

type IDNumber struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	StateCode string `json:"state_code,omitempty"`
}

type Applicant struct {
	ID         string     `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	MiddleName string     `json:"middle_name,omitempty"`
	Email      string     `json:"email,omitempty"`
	Gender     string     `json:"gender"`
	Dob        *Date      `json:"dob"`
	Country    string     `json:"country,omitempty"`
	Mobile     string     `json:"mobile,omitempty"`
	Telephone  string     `json:"telephone,omitempty"`
	IDNumbers  []IDNumber `json:"id_numbers"`
	Addresses  []Address  `json:"addresses"`

	Href      string     `json:"href,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type applicantsResponse struct {
	Applicants []*Applicant `json:"applicants"`
}
