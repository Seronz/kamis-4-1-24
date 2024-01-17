package users

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"log"
	"net/http"

	encryption "github.com/seronz/api/src/utils/Encryption"
	jwt "github.com/seronz/api/src/utils/JWT"
	"gorm.io/gorm"
)

func (l *User) getUserCredentials(db *gorm.DB) error {
	err := db.Select("id, first_name, last_name, user_role, email, password, salt").Where("email = ?", l.Email).First(&l).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *User) updateActivateRememberMe(db *gorm.DB, remember string) error {
	result := db.Model(&User{}).Where("email = ?", l.Email).
		Update("remember_me", l.RememberMe).
		Update("remember_token", remember)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteRememberToken(db *gorm.DB, user User) error {
	var u User

	u.Email = user.Email
	err := u.getUserCredentials(db)
	if err != nil {
		return err
	}
	remember := ""
	err = u.updateActivateRememberMe(db, remember)
	if err != nil {
		return err
	}
	return nil
}

func CreateRememberToken(db *gorm.DB, w http.ResponseWriter, user User) (string, error) {
	var u User

	u.Email = user.Email
	u.RememberMe = user.RememberMe
	err := u.getUserCredentials(db)
	if err != nil {
		return "", err
	}

	params := jwt.Params{
		W:          w,
		ID:         u.ID,
		Firstname:  u.FirstName,
		Lastname:   u.LastName,
		Userrole:   u.UserRole,
		Email:      u.Email,
		RememberMe: u.RememberMe,
	}
	remember, err := jwt.CreateRememberToken(params)
	if err != nil {
		return "", err
	}
	err = u.updateActivateRememberMe(db, remember)
	if err != nil {
		return "", err
	}
	return remember, nil
}

func Login(db *gorm.DB, w http.ResponseWriter, user User) (string, error) {
	var u User

	u.Email = user.Email
	err := u.getUserCredentials(db)
	if err != nil {
		log.Println(fmt.Errorf("%s %s login failed \n %s", u.FirstName, u.LastName, err))
		return "", err
	}

	if u.Password == "" {
		log.Println(fmt.Errorf("%s %s login failed : user not found", u.FirstName, u.LastName))
		return "", errors.New("user not found")
	}

	s, err := encryption.DecryptPassword(user.Password, u.Salt)
	if err != nil {
		log.Println(fmt.Errorf("%s %s login failed \n %s", u.FirstName, u.LastName, err))
		return "", err
	}

	if subtle.ConstantTimeCompare([]byte(u.Password), []byte(s)) != 1 {
		log.Println(fmt.Errorf("%s %s login failed. invalid username or password", u.FirstName, u.LastName))
		return "", errors.New("invalid username or password")
	}

	params := jwt.Params{
		W:          w,
		ID:         u.ID,
		Firstname:  u.FirstName,
		Lastname:   u.LastName,
		Userrole:   u.UserRole,
		Email:      u.Email,
		RememberMe: u.RememberMe,
	}

	token, err := jwt.CreateToken(params)
	if err != nil {
		log.Println(fmt.Errorf("%s %s login failed \n %s", u.FirstName, u.LastName, err))
		return "", err
	}

	log.Printf("%s %s login success... \n", u.FirstName, u.LastName)
	return token, nil
}
