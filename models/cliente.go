package models

import "gopkg.in/validator.v2"

type Cliente struct {
	Id             int
	Nome           string `json:"nome" validate:"nonzero"`
	Cpf            string `json:"cpf" validate:"len=14"`
	Rg             string `json:"rg" validate:"max=13"`
	Email          string `json:"email" validate:"regexp=^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$"`
	Telefone       string `json:"telefone" validate:"regexp=^[(]+[0-9]+[)]+[0-9]{9}"`
	Celular        string `json:"celular" validate:"regexp=^[(]+[0-9]+[)]+[0-9]{9}"`
	Datanascimento string `json:"datanascimento" validade:"len=10"`
	Sexo           string `json:"sexo"`
	Profissao      string `json:"profissao"`
}

func ValidaDadosClientes(cliente *Cliente) error {
	if err := validator.Validate(cliente); err != nil {
		return err
	}
	return nil
}
