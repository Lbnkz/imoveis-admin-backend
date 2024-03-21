package models

import (
	"gorm.io/gorm"
)

type House struct {
    gorm.Model
    Endereco            string  `gorm:"column:endereco"`
    NumQuartos          int     `gorm:"column:num_quartos"`
    NumBanheiros        int     `gorm:"column:num_banheiros"`
    VagasGaragem        int     `gorm:"column:vg_garagem"`
    AreaMetrosQuadrados int     `gorm:"column:area_metros_quadrados"`
    Descricao           string  `gorm:"column:descricao"`
    Preco               float64 `gorm:"column:preco"`
}
