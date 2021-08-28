package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Room_id string
}
