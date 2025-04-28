package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);column:name;uniqueIndex:idx_person_name" json:"name"`
	Age  uint64 `gorm:"column:age" json:"age"`
}
