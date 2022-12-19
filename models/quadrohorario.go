package models

import "github.com/lib/pq"

type Quadroshorarios struct {
	Id      int
	Domingo pq.Int64Array `gorm:"type:integer[]"`
	Segunda pq.Int64Array `gorm:"type:integer[]"`
	Terca   pq.Int64Array `gorm:"type:integer[]"`
	Quarta  pq.Int64Array `gorm:"type:integer[]"`
	Quinta  pq.Int64Array `gorm:"type:integer[]"`
	Sexta   pq.Int64Array `gorm:"type:integer[]"`
	Sabado  pq.Int64Array `gorm:"type:integer[]"`
}
