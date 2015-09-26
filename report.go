package onfido

import "time"

var ReportType = struct {
	GBR struct {
		Identity            string
		Document            string
		Employment          string
		Education           string
		NegativeMedia       string
		Directorship        string
		CriminalHistory     string
		Watchlist           string
		AntiMoneyLaundering string
		StreetLevel         string
	}
	USA struct {
		Identity         string
		Document         string
		SexOffender      string
		Watchlist        string
		NationalCriminal string
		Eviction         string
		CountyCriminal   string
		DrivingRecord    string
	}
	Europe struct {
		Identity        string
		Document        string
		CriminalHistory string
	}
}{
	GBR: struct {
		Identity            string
		Document            string
		Employment          string
		Education           string
		NegativeMedia       string
		Directorship        string
		CriminalHistory     string
		Watchlist           string
		AntiMoneyLaundering string
		StreetLevel         string
	}{
		Identity:            "identity",
		Document:            "document",
		Employment:          "employment",
		Education:           "education",
		NegativeMedia:       "negative_media",
		Directorship:        "directorship",
		CriminalHistory:     "criminal_history",
		Watchlist:           "watchlist",
		AntiMoneyLaundering: "anti_money_laundering",
		StreetLevel:         "street_level",
	},
	USA: struct {
		Identity         string
		Document         string
		SexOffender      string
		Watchlist        string
		NationalCriminal string
		Eviction         string
		CountyCriminal   string
		DrivingRecord    string
	}{
		Identity:         "identity",
		Document:         "document",
		SexOffender:      "sex_offender",
		Watchlist:        "watchlist",
		NationalCriminal: "national_criminal",
		Eviction:         "eviction",
		CountyCriminal:   "county_criminal",
		DrivingRecord:    "driving_record",
	},
	Europe: struct {
		Identity        string
		Document        string
		CriminalHistory string
	}{
		Identity:        "identity",
		Document:        "document",
		CriminalHistory: "criminal_history",
	},
}

var ReportStatus = struct {
	Pending  string
	Complete string
}{
	Pending:  "pending",
	Complete: "complete",
}

var ResultKind = struct {
	Clear        string
	Consider     string
	Fail         string
	Unidentified string
	None         string
}{
	Clear:        "clear",
	Consider:     "consider",
	Fail:         "fail",
	Unidentified: "unidentified",
	None:         "none",
}

type Report struct {
	ID         string                 `json:"id"`
	Href       string                 `json:"href"`
	Name       string                 `json:"name"`
	Result     string                 `json:"result"`
	Status     string                 `json:"status"`
	Breakdown  map[string]interface{} `json:"breakdown"`
	Properties map[string]interface{} `json:"properties"`
	CreatedAt  time.Time              `json:"created_at"`
}

// func (r *Report) DrivingRecord() *DrivingRecord {
// 	return nil
// }
