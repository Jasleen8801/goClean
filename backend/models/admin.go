package models

import (
	"goClean/backend/utils"
	"html"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
}

func (a *Admin) SaveAdmin() (*Admin, error) {
	// Saves the admin to the database
	err := DB.Create(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	return a, nil
}

func (a *Admin) BeforeSave() error {
	// Hashes the password before saving
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hashedPwd)
	a.Name = html.EscapeString(a.Name)
	return nil
}

func (a *Admin) VerifyPwd(pwd, hashedPwd string) error {
	// Compares the password with the hashed password
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}

func AdminLoginCheck(name string, pwd string) (string, error) {
	// Checks if the admin is registered and returns a token if true
	var err error

	a := Admin{}

	err = DB.Model(&Admin{}).Where("name = ?", name).Take(&a).Error
	if err != nil {
		return "", err
	}

	err = a.VerifyPwd(pwd, a.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return utils.GenerateToken(a.ID, utils.Admin)
}

func AuthenticateAdmin(username, password string) error {
	admin := Admin{}
	err := DB.Where("name = ?", username).Take(&admin).Error
	if err != nil {
		return err
	}

	err = admin.VerifyPwd(password, admin.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}

	return nil
}

func GetAdminById(id uint) (*Admin, error) {
	// Gets the admin by id
	var a *Admin
	err := DB.Model(&Admin{}).Where("id = ?", id).Take(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	return a, nil
}

func (a *Admin) AddHostel(h *Hostel) (*Hostel, error) {
	// Adds a hostel to the database
	err := DB.Create(&h).Error
	if err != nil {
		return &Hostel{}, err
	}
	return h, nil
}
