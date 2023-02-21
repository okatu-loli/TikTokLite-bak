package model

import (
	"time"
)

type BasePo struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
