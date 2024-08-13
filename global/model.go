package global

import (
	"gorm.io/gorm"
	"time"
)

type LYY_MODEL struct {
	ID        uint `gorm:"primarkey" json:"ID"`
	CreateDAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
