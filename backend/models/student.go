package models

import (
	"goClean/backend/utils"
	"html"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	RollNo   string `gorm:"type:varchar(100);not null" json:"roll_no"`
	Hostel   string `gorm:"type:varchar(100);not null" json:"hostel"`
	RoomNo   string `gorm:"type:varchar(100);not null" json:"room_no"`
}

func (s *Student) SaveUser() (*Student, error) {
	// Saves the student to the database
	err := DB.Create(&s).Error
	if err != nil {
		return &Student{}, err
	}
	return s, nil
}

func (s *Student) BeforeSave() error {
	// Hashes the password before saving
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	s.Password = string(hashedPwd)
	s.Name = html.EscapeString(s.Name)
	return nil
}

func (s *Student) VerifyPwd(pwd, hashedPwd string) error {
	// Compares the password with the hashed password
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}

func LoginCheck(name string, pwd string) (string, error) {
	// Checks if the student is registered and returns a token if true
	var err error

	s := Student{}

	err = DB.Model(&Student{}).Where("name = ?", name).Take(&s).Error
	if err != nil {
		return "", err
	}

	err = s.VerifyPwd(pwd, s.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(s.ID, utils.Student)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetStudents(RoomId uint) (*[]Student, error) {
	// Gets all the students in a particular room
	var s []Student
	err := DB.Model(&Student{}).Where("room_no = ?", RoomId).Find(&s).Error
	if err != nil {
		return &[]Student{}, err
	}
	return &s, nil
}
