package services

import (
	"data_hoarder_go/config"
	"data_hoarder_go/models"
)

func GetAllFIRs() (
	[]models.FIR,
	error,
) {

	rows, err := config.DB.Query(
		`
		SELECT *
		FROM fir_records
		`,
	)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var firs []models.FIR

	for rows.Next() {

		var fir models.FIR

		err := rows.Scan(
			&fir.ID,
			&fir.PoliceStationCode,
			&fir.LawEnforcementAgency,
			&fir.NationalCode,
			&fir.LastFIRNo,
			&fir.LastFIRYear,
			&fir.LargeChargeSheetDate,
			&fir.District,
			&fir.LawEnforcementType,
			&fir.DisplayStatus,
		)

		if err != nil {

			return nil, err
		}

		firs = append(
			firs,
			fir,
		)
	}

	return firs, nil
}

func CreateFIR(
	fir models.FIR,
) error {

	query := `
	INSERT INTO fir_records
	(
		police_station_code,
		law_enforcement_agency,
		national_code,
		last_fir_no,
		last_fir_year,
		large_charge_sheet_date,
		district,
		law_enforcement_type,
		display_status
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,$7,$8,$9
	)
	`

	_, err := config.DB.Exec(
		query,

		fir.PoliceStationCode,
		fir.LawEnforcementAgency,
		fir.NationalCode,
		fir.LastFIRNo,
		fir.LastFIRYear,
		fir.LargeChargeSheetDate,
		fir.District,
		fir.LawEnforcementType,
		fir.DisplayStatus,
	)

	return err
}

func UpdateFIR(
	id string,
	fir models.FIR,
) error {

	query := `
	UPDATE fir_records
	SET
		police_station_code=$1,
		law_enforcement_agency=$2,
		national_code=$3,
		last_fir_no=$4,
		last_fir_year=$5,
		large_charge_sheet_date=$6,
		district=$7,
		law_enforcement_type=$8,
		display_status=$9
	WHERE id=$10
	`

	_, err := config.DB.Exec(
		query,

		fir.PoliceStationCode,
		fir.LawEnforcementAgency,
		fir.NationalCode,
		fir.LastFIRNo,
		fir.LastFIRYear,
		fir.LargeChargeSheetDate,
		fir.District,
		fir.LawEnforcementType,
		fir.DisplayStatus,
		id,
	)

	return err
}

func DeleteFIR(
	id string,
) error {

	_, err := config.DB.Exec(
		`
		DELETE FROM fir_records
		WHERE id=$1
		`,
		id,
	)

	return err
}