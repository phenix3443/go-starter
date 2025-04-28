package database

import (
	"fmt"

	"github.com/phenix3443/go-starter/pkg/database/models"
	"gorm.io/gorm/clause"
)

func (db *AppDB) SavePerson(person *models.Person) error {
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		UpdateAll: true,
	}).Create(person).Error; err != nil {
		return fmt.Errorf("failed to save person: %w", err)
	}
	return nil
}

func (db *AppDB) GetPerson(name string) (*models.Person, error) {
	var person models.Person
	if err := db.Where("name = ?", name).First(&person).Error; err != nil {
		return nil, fmt.Errorf("failed to get person: %w", err)
	}
	return &person, nil
}
