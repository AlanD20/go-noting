package migrations

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	ID         uint64
	Title      string
	Body       string
	Private    bool
	Created_at string
	Updated_at string
	UserId     uint
}
