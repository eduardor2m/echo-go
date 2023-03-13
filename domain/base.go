package domain

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint `gorm:"uuid,primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("CreatedAt", time.Now())
	tx.Statement.SetColumn("UpdatedAt", time.Now())

	return nil

}
