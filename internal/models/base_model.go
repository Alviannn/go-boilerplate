package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64          `gorm:"primaryKey;type:bigserial" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt" swaggertype:"string"`
}
