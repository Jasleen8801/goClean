package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	StudentName  string    `gorm:"type:varchar(100);not null" json:"student_name"`
	Hostel       string    `gorm:"type:varchar(100);not null" json:"hostel"`
	RoomNo       string    `gorm:"type:varchar(100);not null" json:"room_no"`
	Floor        string    `gorm:"type:varchar(100);not null" json:"floor"`
	Feedback     string    `gorm:"type:varchar(100);not null" json:"feedback" default:""`
	CleaningTime time.Time `gorm:"not null" json:"cleaning_time"`
}

func (l *Log) SaveLog() (*Log, error) {
	// Saves the log to the database
	err := DB.Create(&l).Error
	if err != nil {
		return &Log{}, err
	}
	return l, nil
}
