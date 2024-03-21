package models

import (
	"gorm.io/gorm"
)

type Apartament struct {
    gorm.Model
    Endereco            string  `gorm:"column:endereco"`
    NumQuartos          int     `gorm:"column:num_quartos"`
    NumBanheiros        int     `gorm:"column:num_banheiros"`
    VagasGaragem        int     `gorm:"column:vg_garagem"`
    AreaMetrosQuadrados int     `gorm:"column:area_metros_quadrados"`
    Descricao           string  `gorm:"column:descricao"`
    Andar               int     `gorm:"column:andar"`
    Preco               float64 `gorm:"column:preco"`
}
