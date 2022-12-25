package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `gorm:"type:uuid;primaryKey;not null" json:"id,omitempty"`
	CreatedAt time.Time      `gorm:"not null;" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"not null;" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (b *BaseModel) New() {

	b.ID = uuid.NewString()
}
