package models

import "gopkg.in/validator.v2"

type Atendimento struct {
	Id             int
	Idservico      int     `json:"idservico"`
	Idatendente    int     `json:"idatendente"`
	Idcliente      int     `json:"idcliente"`
	Nome           string  `json:"nome"`
	Horario        float64 `json:"horario"`
	Status         string
	Statusfixo     bool
	Formapagamento string
}

func ValidaDadosAtendimento(atendimento *Atendimento) error {
	if err := validator.Validate(atendimento); err != nil {
		return err
	}
	return nil
}
