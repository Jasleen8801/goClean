package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Hostel  string `gorm:"type:varchar(100);not null" json:"hostel"`
	RoomNo  string `gorm:"type:varchar(100);not null" json:"room_no"`
	Floor   string `gorm:"type:varchar(100);not null" json:"floor"`
	Cleaned bool   `gorm:"type:boolean;not null" json:"cleaned" default:"false"`
}

func (r *Room) SaveRoom() (*Room, error) {
	// Check if the room already exists
	var existingRoom Room
	err := DB.Model(&Room{}).Where("hostel = ? AND room_no = ? AND floor = ?", r.Hostel, r.RoomNo, r.Floor).Take(&existingRoom).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == nil {
		return &existingRoom, nil
	}

	err = DB.Create(&r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}
