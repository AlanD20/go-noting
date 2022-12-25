package model

import (
	"strings"

	"gorm.io/gorm"
)

type UserSchema struct {
	Username string `gorm:"unique;not null;" json:"username,omitempty" form:"username"`
	Email    string `gorm:"unique;not null;" json:"email,omitempty" form:"email"`
	Password string `gorm:"not null;" json:"password,omitempty" form:"password"`
}

type User struct {
	BaseModel
	UserSchema
	Notes []Note `gorm:"foreignKey:UserId;Constraints:onDelete:Cascade;" json:"notes"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	u.BaseModel.New()
	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)

	return
}

func (u *UserSchema) NewUser() *User {

	user := &User{
		UserSchema: UserSchema{
			Username: u.Username,
			Email:    u.Email,
			Password: u.Password,
		},
	}

	return user
}

func (u *UserSchema) UpdateUser(user *User) {

	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password

}
