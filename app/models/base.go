package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID string `gorm:"type:uuid;primaryKey;not null"`
}

func (b *BaseModel) New() {

	b.ID = uuid.NewString()
}
