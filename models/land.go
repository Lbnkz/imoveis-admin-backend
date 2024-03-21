package models

import "gorm.io/gorm"

type Land struct {
    gorm.Model
    Endereco            string  `gorm:"column:endereco"`
    AreaMetrosQuadrados int     `gorm:"column:area_metros_quadrados"`
    Descricao           string  `gorm:"column:descricao"`
    Preco               float64 `gorm:"column:preco"`
}