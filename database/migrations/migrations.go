package migrations

import (
	"imoveis-back/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Apartamentos{})
	db.AutoMigrate(models.Casa{})
	db.AutoMigrate(models.Terreno{})
}
