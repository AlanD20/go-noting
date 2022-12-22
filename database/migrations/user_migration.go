package migrations

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint64
	Username   string
	Email      string
	Password   string
	Created_at string
	Updated_at string
	Notes      []Note
}
