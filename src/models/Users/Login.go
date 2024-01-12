package users

import (
	"crypto/subtle"
	"errors"
	"net/http"

	encryption "github.com/seronz/api/src/utils/Encryption"
	jwt "github.com/seronz/api/src/utils/JWT"
	"gorm.io/gorm"
)

func (l *User) getUserCredentials(db *gorm.DB) error {
	err := db.Select("first_name, last_name, user_role,email, password, salt").Where("email = ?", l.Email).First(&l).Error
	if err != nil {
		return err
	}
	return nil
}

func Login(db *gorm.DB, w http.ResponseWriter, user User) (string, error) {
	var u User

	u.Email = user.Email
	err := u.getUserCredentials(db)
	if err != nil {
		return "", err
	}

	if u.Password == "" {
		return "", errors.New("user not found")
	}

	s, err := encryption.DecryptPassword(u.Password, u.Salt)
	if err != nil {
		return "", err
	}

	if subtle.ConstantTimeCompare([]byte(u.Password), []byte(s)) != 1 {
		return "", errors.New("invalid username or password")
	}

	params := jwt.Params{
		W:         w,
		Firstname: u.FirstName,
		Lastname:  u.LastName,
		Userrole:  u.UserRole,
		Email:     u.Email,
	}
	token, err := jwt.CreateToken(params)
	if err != nil {
		return "", err
	}
	return token, nil
}
