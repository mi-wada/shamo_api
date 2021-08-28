package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Price    int32  `json:"price"`
	Room_id  string `json:"room_id"`
	User_id  uint   `json:"user_id"`
	What     string `json:"what"`
	Is_valid bool
}
