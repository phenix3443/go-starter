package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/phenix3443/go-starter/pkg/database/models"
)

type AppDB struct {
	*gorm.DB
}

func NewAppDB(dsn string) (*AppDB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect app database: %w", err)
	}
	if err = db.AutoMigrate(
		&models.Person{},
	); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate: %w", err)
	}
	return &AppDB{DB: db}, nil
}

func (db *AppDB) Name() string {
	return "App Database"
}

func (db *AppDB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sqlDB: %w", err)
	}
	return sqlDB.Close()
}
