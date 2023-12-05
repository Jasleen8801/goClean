package models

import (
	"fmt"
	"goClean/backend/utils"
	"html"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Hostel struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Hostel   string `gorm:"type:varchar(100);not null" json:"hostel"`
}

func (h *Hostel) VerifyPwd(pwd, hashedPwd string) error {
	// Compares the password with the hashed password
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}

func HostelLoginCheck(name string, pwd string) (string, error) {
	// Checks if the hostel is registered and returns a token if true
	var err error

	h := Hostel{}

	err = DB.Model(&Hostel{}).Where("name = ?", name).Take(&h).Error
	if err != nil {
		return "", err
	}

	err = h.VerifyPwd(pwd, h.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return utils.GenerateToken(h.ID, utils.Hostel)
}

// func (h *Hostel) BeforeSave(tx *gorm.DB) error {
func (h *Hostel) BeforeSave() error {
	// Hashes the password before saving
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(h.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	h.Password = string(hashedPwd)
	h.Name = html.EscapeString(h.Name)
	return nil
}

func GetHostelById(id uint) (*Hostel, error) {
	// Gets the hostel by id
	var h *Hostel
	err := DB.Model(&Hostel{}).Where("id = ?", id).Take(&h).Error
	if err != nil {
		return &Hostel{}, err
	}
	return h, nil
}

func (h *Hostel) ClearAllLogs() error {
	// Clears all the logs of the hostel
	err := DB.Where("hostel = ?", h.Hostel).Delete(&Log{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *Hostel) GetAllUncleanedRoomNos() ([]string, error) {
	// Gets all the uncleaned room numbers of the hostel
	var rooms []string
	err := DB.Model(&Room{}).Where("hostel = ? AND cleaned = ?", h.Hostel, false).Pluck("room_no", &rooms).Error
	if err != nil {
		return []string{}, err
	}
	// fmt.Println(rooms)
	return rooms, nil
}

func (h *Hostel) GetAllCleanedRoomNos() (*[]Room, error) {
	// Gets all the cleaned room numbers of the hostel
	var r []Room
	err := DB.Model(&Room{}).Where("hostel = ? AND cleaned = ?", h.Hostel, true).Find(&r).Error
	if err != nil {
		return &[]Room{}, err
	}
	fmt.Println(r)
	return &r, nil
}

func (h *Hostel) GetAllLogs() (*[]Log, error) {
	// Gets all the logs of the hostel
	var logs []Log
	err := DB.Model(&Log{}).Where("hostel = ?", h.Hostel).Find(&logs).Error
	if err != nil {
		return &[]Log{}, err
	}
	return &logs, nil
}
