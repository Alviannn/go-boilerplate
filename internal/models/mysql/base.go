package models_mysql

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt" swaggertype:"string"`
}

func (m *BaseModel) IsExist() bool {
	return m.ID != 0
}
