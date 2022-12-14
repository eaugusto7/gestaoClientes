package models

import "gopkg.in/validator.v2"

type Servico struct {
	Id    int
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
	Tempo float64 `json:"tempo"`
}

func ValidaDadosServico(servico *Servico) error {
	if err := validator.Validate(servico); err != nil {
		return err
	}
	return nil
}
