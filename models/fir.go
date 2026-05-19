package models

import "time"

type FIR struct {

	ID int `json:"id"`

	PoliceStationCode string `json:"police_station_code"`

	LawEnforcementAgency string `json:"law_enforcement_agency"`

	NationalCode string `json:"national_code"`

	LastFIRNo string `json:"last_fir_no"`

	LastFIRYear string `json:"last_fir_year"`

	LargeChargeSheetDate time.Time `json:"large_charge_sheet_date"`

	District string `json:"district"`

	LawEnforcementType string `json:"law_enforcement_type"`

	DisplayStatus bool `json:"display_status"`
}