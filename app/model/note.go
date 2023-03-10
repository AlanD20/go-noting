package model

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteSchema struct {
	UserId  uuid.UUID `gorm:"type:uuid;" json:"user_id" form:"user_id"`
	Title   string    `gorm:"not null;" json:"title" form:"title"`
	Body    string    `gorm:"not null;" json:"body" form:"body"`
	Private bool      `gorm:"default:false;" json:"private" form:"private"`
}

type Note struct {
	BaseModel
	NoteSchema
	User User `gorm:"foreignKey:UserId;references:ID"`
}

func (n *Note) BeforeCreate(tx *gorm.DB) (err error) {

	n.BaseModel.New()
	n.Title = strings.ToTitle(n.Title)

	return
}

func (n *NoteSchema) Create(conn *gorm.DB) error {

	note := Note{
		NoteSchema: NoteSchema{
			Title:   n.Title,
			Body:    n.Body,
			Private: n.Private,
			UserId:  n.UserId,
		},
	}

	err := conn.Create(&note).Error

	return err
}
