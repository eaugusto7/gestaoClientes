package models

import "gopkg.in/validator.v2"

type Atendente struct {
	Id              int
	Nome            string `json:"nome"`
	Celular         string `json:"celular" validate:"regexp=^[(]+[0-9]+[)]+[0-9]{9}"`
	Idquadrohorario int
}

func ValidaDadosAtendente(atendente *Atendente) error {
	if err := validator.Validate(atendente); err != nil {
		return err
	}
	return nil
}
