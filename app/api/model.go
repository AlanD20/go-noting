package api

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string          `gorm:"type:uuid;primaryKey;not null" json:"id,omitempty"`
	CreatedAt time.Time       `gorm:"not null;"                     json:"created_at,omitempty"`
	UpdatedAt time.Time       `gorm:"not null;"                     json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index"                         json:"deleted_at,omitempty"`
}

func (b *BaseModel) GenerateId() {
	b.ID = uuid.NewString()
}

// ---------------
// User Model
// ---------------
type UserSchema struct {
	Username string `gorm:"unique;not null;" json:"username,omitempty" form:"username"`
	Email    string `gorm:"unique;not null;" json:"email,omitempty"    form:"email"`
	Password string `gorm:"not null;"        json:"password,omitempty" form:"password"`
}

type User struct {
	BaseModel
	UserSchema
	Notes []Note `gorm:"foreignKey:UserId;references:ID;Constraints:onDelete:Cascade;" json:"notes"`
}

// ---------------
// Note Model
// ---------------
type NoteSchema struct {
	Title   string `gorm:"not null;"      json:"title,omitempty"   form:"title"`
	Body    string `gorm:"not null;"      json:"body,omitempty"    form:"body"`
	Private bool   `gorm:"default:false;" json:"private,omitempty" form:"private"`
}

type Note struct {
	BaseModel
	NoteSchema
	User   *User  `gorm:"foreignKey:UserId;references:ID" json:"user,omitempty"`
	UserId string `gorm:"type:uuid;"                      json:"user_id,omitempty" form:"user_id"`
}
