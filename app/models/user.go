package models

import (
	"strings"

	"github.com/aland20/go-noting/database"
	"gorm.io/gorm"
)

type UserSchema struct {
	Username string `gorm:"unique;not null;" json:"username" form:"username"`
	Email    string `gorm:"unique;not null;" json:"email" form:"email"`
	Password string `gorm:"not null;" json:"password" form:"password"`
}

type User struct {
	BaseModel
	UserSchema
	Notes []Note `gorm:"foreignKey:UserId;Constraints:onDelete:Cascade;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	u.BaseModel.New()
	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)

	return
}

func (u *UserSchema) Create() error {

	user := User{
		UserSchema: UserSchema{
			Username: u.Username,
			Email:    u.Email,
			Password: u.Password,
		},
	}

	err := database.Connect().Create(&user).Error

	return err
}
