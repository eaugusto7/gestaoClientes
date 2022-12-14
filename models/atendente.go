package models

import "gopkg.in/validator.v2"

type Atendente struct {
	Id       int
	Nome     string `json:"nome"`
	Telefone string `json:"telefone" validate:"regexp=^[(]+[0-9]+[)]+[0-9]{9}"`
}

func ValidaDadosAtendente(atendente *Atendente) error {
	if err := validator.Validate(atendente); err != nil {
		return err
	}
	return nil
}
