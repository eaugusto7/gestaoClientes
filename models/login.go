package models

import "gopkg.in/validator.v2"

type Login struct {
	Id       int
	Username string
	Password string
}

func ValidaDadosLogin(login *Login) error {
	if err := validator.Validate(login); err != nil {
		return err
	}
	return nil
}
