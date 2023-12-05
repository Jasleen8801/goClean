package models

import (
	"fmt"
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
	Floor    string `gorm:"type:varchar(100);not null" json:"floor"`
}

func (s *Student) SaveUser() (*Student, error) {
	// Saves the student to the database
	var existingStudent Student
	err := DB.Model(&Student{}).Where("roll_no = ?", s.RollNo).Take(&existingStudent).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == nil {
		return &existingStudent, nil
	}

	err = DB.Create(&s).Error
	if err != nil {
		return &Student{}, err
	}
	return s, nil
}

// func (s *Student) BeforeSave(tx *gorm.DB) error {
func (s *Student) BeforeSave() error {
	// Hashes the password before saving
	fmt.Println(s.Password)
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	s.Password = string(hashedPwd)
	s.Name = html.EscapeString(s.Name)
	return nil
}

func (s *Student) VerifyPwd(pwd string, hashedPwd []byte) error {
	// Compares the password with the hashed password
	return bcrypt.CompareHashAndPassword(hashedPwd, []byte(pwd))
}

func LoginCheck(name string, pwd string) (string, error) {
	// Checks if the student is registered and returns a token if true
	var err error

	s := Student{}

	err = DB.Model(&Student{}).Where("name = ?", name).Take(&s).Error
	if err != nil {
		return "", err
	}

	// fmt.Println(s.Password)
	err = s.VerifyPwd(pwd, []byte(s.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(s.ID, utils.Student)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetStudentById(id uint) (*Student, error) {
	// Gets the student by id
	var s *Student
	err := DB.Model(&Student{}).Where("id = ?", id).Take(&s).Error
	if err != nil {
		return &Student{}, err
	}
	return s, nil
}

func GetStudentsbyRoomID(RoomId string) (*[]Student, error) {
	// Gets all the students in a particular room
	var s []Student
	err := DB.Model(&Student{}).Where("room_no = ?", RoomId).Find(&s).Error
	if err != nil {
		return &[]Student{}, err
	}
	return &s, nil
}

func UpdateRoomCleaned(RoomId string) error {
	// Updates the cleaned status of a room
	// err := DB.Model(&Student{}).Where("room_no = ?", RoomId).Update("cleaned", gorm.Expr("NOT cleaned")).Error
	// if err != nil {
	// 	return err
	// }

	err := DB.Model(&Room{}).Where("room_no = ?", RoomId).Update("cleaned", gorm.Expr("NOT cleaned")).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Student) GetAllLogs() (*[]Log, error) {
	// Gets all the logs of the student
	var l []Log
	err := DB.Model(&Log{}).Where("student_name = ?", s.Name).Find(&l).Error
	if err != nil {
		return &[]Log{}, err
	}
	return &l, nil
}
