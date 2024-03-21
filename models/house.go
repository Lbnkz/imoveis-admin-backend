package models

import (
	"gorm.io/gorm"
)

type Casa struct {
	gorm.Model
	Endereco            string  `gorm:"column:endereco"`
    NumSuites           int     `gorm:"column:num_suites"`
	NumQuartos          int     `gorm:"column:num_quartos"`
	NumBanheiros        int     `gorm:"column:num_banheiros"`
	VagasGaragem        int     `gorm:"column:vg_garagem"`
	AreaMetrosQuadrados int     `gorm:"column:area_metros_quadrados"`
	AreaConstruida      int     `gorm:"column:area_construida"`
	Descricao           string  `gorm:"column:descricao"`
	Preco               float64 `gorm:"column:preco"`
}
