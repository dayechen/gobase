package dbmodel

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	Label string
	Url   string
}
