package onfido

import (
	"flag"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/tylerb/is"
)

var token = flag.String("token", "", "your sandbox API token for onfido")

func checkToken(is *is.Is) {
	if *token == "" {
		is.TB.Fatal(`You must provide a sandbox API token (hint: --token=[your token]) to run this test. If you don't have a sandbox token, you may set the --short flag to skip.`)
	}
}

var nowDate = Date(time.Now().Add(-5 * time.Hour * 24))

var (
	testIDNumbers = []IDNumber{
		{
			Type:      IDNumberType.DrivingLicense,
			Value:     "I1234562",
			StateCode: "CA",
		},
		{
			Type:  IDNumberType.SSN,
			Value: "433-54-3934",
		},
	}
	testAddress = []Address{
		{
			Street:    "123 Fake Street",
			Town:      "Sacramento",
			State:     "CA",
			Postcode:  "92443",
			Country:   "USA",
			StartDate: &nowDate,
		},
	}
)

func TestMain(m *testing.M) {

	rand.Seed(time.Now().UnixNano())

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestClientNew(t *testing.T) {
	is := is.New(t)

	c := New("testing")
	is.Equal(c.apiToken, "testing")
	is.Equal(c.apiTokenHeader, "Token token=testing")
}

func TestClientApplicants(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	is := is.New(t)
	checkToken(is)

	c := New(*token)
	is.NotNil(c)

	nameNum := rand.Int31()

	nowDate = Date(time.Now().Add(-5 * time.Hour * 24))

	a := &Applicant{
		Country:   "USA",
		FirstName: "Tyler" + strconv.Itoa(int(nameNum)),
		LastName:  "Bunnell" + strconv.Itoa(int(nameNum)),
		Email:     "tyler" + strconv.Itoa(int(nameNum)) + "@fake.com",
		Gender:    "male",
		Dob:       &nowDate,
		IDNumbers: testIDNumbers,
		Telephone: "555-738-6874",
		Addresses: testAddress,
	}

	a, err := c.CreateApplicant(a)
	is.NotErr(err)
	is.NotZero(a.ID)
	is.NotZero(a.Href)

	a, err = c.ReadApplicant(a.ID)
	is.NotErr(err)
	is.NotZero(a.ID)
	is.NotZero(a.Href)

	apps, err := c.ReadApplicants()
	is.NotErr(err)
	is.NotEqual(len(apps), 0)
}

func TestClientChecks(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	is := is.New(t)
	checkToken(is)

	c := New(*token)
	is.NotNil(c)

	nowDate = Date(time.Now().Add(-5 * time.Hour * 24))

	nameNum := rand.Int31()

	a := &Applicant{
		Country:   "USA",
		FirstName: "Tyler" + strconv.Itoa(int(nameNum)),
		LastName:  "Bunnell" + strconv.Itoa(int(nameNum)),
		Email:     "tyler" + strconv.Itoa(int(nameNum)) + "@fake.com",
		Gender:    "male",
		Dob:       &nowDate,
		IDNumbers: testIDNumbers,
		Telephone: "555-738-6874",
		Addresses: testAddress,
	}

	a, err := c.CreateApplicant(a)
	is.NotErr(err)
	is.NotZero(a.ID)
	is.NotZero(a.Href)

	ch, err := c.CreateCheck(a.ID,
		NewCheckRequest(CheckType.Express,
			ReportType.USA.Identity, ReportType.USA.DrivingRecord))
	is.NotErr(err)
	is.NotNil(ch)
	is.NotZero(ch.ID)
	is.NotZero(ch.CreatedAt)
	is.Equal(ch.Type, "express")
	is.Equal(ch.Result, "consider")
	is.NotZero(ch.Href)
	is.NotNil(ch.Reports)
	is.Equal(len(ch.Reports), 2)

	ch, err = c.ReadCheck(a.ID, ch.ID)
	is.NotErr(err)
	is.NotNil(ch)
	is.NotZero(ch.ID)
	is.NotZero(ch.CreatedAt)
	is.Equal(ch.Type, "express")
	is.Equal(ch.Result, "consider")
	is.NotZero(ch.Href)
	is.NotNil(ch.Reports)
	is.Equal(len(ch.Reports), 2)

	chs, err := c.ReadChecks(a.ID)
	is.NotErr(err)
	is.NotEqual(len(chs), 0)
}
