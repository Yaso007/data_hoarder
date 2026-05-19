package services

import (
	"data_hoarder_go/config"
	"data_hoarder_go/models"
	"data_hoarder_go/utils"
	"errors"
)

func RegisterUser(
	user models.User,
) error {

	hashedPassword, err :=
		utils.HashPassword(
			user.Password,
		)

	if err != nil {

		return err
	}

	query := `
	INSERT INTO users
	(
		username,
		password
	)
	VALUES
	(
		$1,
		$2
	)
	`

	_, err = config.DB.Exec(
		query,
		user.Username,
		hashedPassword,
	)

	return err
}

func LoginUser(
	user models.User,
) (string, error) {

	query := `
	SELECT password
	FROM users
	WHERE username=$1
	`

	var hashedPassword string

	err := config.DB.QueryRow(
		query,
		user.Username,
	).Scan(&hashedPassword)

	if err != nil {

		return "",
			errors.New("invalid username")
	}

	valid := utils.CheckPassword(
		user.Password,
		hashedPassword,
	)

	if !valid {

		return "",
			errors.New("invalid password")
	}

	token, err :=
		utils.GenerateToken(
			user.Username,
		)

	return token, err
}