package models

type Atendimento struct {
	Id          int
	Nome        string
	Horario     float64
	Idservico   int
	Idatendente int
}
