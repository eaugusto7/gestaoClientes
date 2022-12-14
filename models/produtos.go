package models

import "gopkg.in/validator.v2"

type Produtos struct {
	Id         int
	Nome       string  `json:"nome"`
	Quantidade int     `json:"quantidade"`
	Descricao  string  `json:"descricao"`
	Fabricante string  `json:"fabricante"`
	Valorcusto float64 `json:"valorcusto"`
	Valorvenda float64 `json:"valorvenda"`
}

func ValidaDadosProdutos(produtos *Produtos) error {
	if err := validator.Validate(produtos); err != nil {
		return err
	}
	return nil
}
