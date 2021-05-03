package compA

import (
	"gorm.io/gorm"
)

type CompA struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
