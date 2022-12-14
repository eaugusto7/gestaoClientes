package models

import "gopkg.in/validator.v2"

type Atendimento struct {
	Id          int
	Nome        string  `json:"nome"`
	Horario     float64 `json:"horario"`
	Idservico   int     `json:"idservico"`
	Idatendente int     `json:"idatendente"`
}

func ValidaDadosAtendimento(atendimento *Atendimento) error {
	if err := validator.Validate(atendimento); err != nil {
		return err
	}
	return nil
}
