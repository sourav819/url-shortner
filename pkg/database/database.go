package database

import (
	"fmt"
	"url-shortner/models"
	"url-shortner/pkg/config"
	"url-shortner/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
)

func GetDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gl.Default.LogMode(gl.Info),
	})
	if err != nil {
		logger.Error("unable to make db connrection")
		return nil, err
	}

	err = db.AutoMigrate(models.GetMigrationModels()...)
	if err != nil {
		logger.Error("unable to migrate models")
		return nil, err
	}
	return db, nil
}
