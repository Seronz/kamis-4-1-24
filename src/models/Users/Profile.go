package users

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/seronz/api/src/utils/JWT"
	otp "github.com/seronz/api/src/utils/OTP"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type User struct {
	ID            string `bson:"_id" gorm:"size:255;not null;uniqueIndex;primary_key" json:"id"`
	Users         []Address
	FirstName     string    `bson:"first_name" gorm:"size:100;not null" json:"first_name"`
	LastName      string    `bson:"last_name" gorm:"size:100;not null" json:"last_name"`
	Email         string    `bson:"email" gorm:"size:100;not null" json:"email"`
	Password      string    `bson:"password" gorm:"size:100;not null" json:"password"`
	UserRole      string    `bson:"user_role" gorm:"size:10;not null;default:'user'" json:"user_role"`
	OtpCode       string    `bson:"otp_code" gorm:"size:10" json:"otp_code"`
	IsActive      bool      `bson:"is_active" gorm:"type:boolean;default:true" json:"is_active"`
	RememberToken string    `bson:"remember_token" gorm:"type:text;not null" json:"remember_token"`
	CreatedAt     time.Time `bson:"CreatedAt" json:"created_at"`
	UpdatedAt     time.Time `bson:"UpdatedAt" json:"updated_at"`
	ExpiredAt     time.Time `bson:"expired_at" json:"expired_at"`
	DeletedAt     gorm.DeletedAt
}

var ctx = context.TODO()

func ConnectionMongo(mg *mongo.Client) *mongo.Collection {
	collection := mg.Database("tokopedia").Collection("user_not_active")
	return collection
}

func UserRegister(mg *mongo.Client, user User) (*mongo.InsertOneResult, string, error) {

	param := jwt.Params{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		Userrole:  user.UserRole,
	}
	token, err := jwt.CreateToken(param)
	if err != nil {
		return nil, "", err
	}

	otp, err := otp.GenerateOTP(user.Email)
	if err != nil {
		return nil, "", err
	}

	expireTime := time.Now().Add(2 * time.Hour)
	user.RememberToken = token
	user.OtpCode = otp
	user.ExpiredAt = expireTime

	userMap := make(map[string]interface{})
	userBytes, err := json.Marshal(user)
	if err != nil {
		return nil, "", err
	}
	if err := json.Unmarshal(userBytes, &userMap); err != nil {
		return nil, "", err
	}

	collection := ConnectionMongo(mg)
	insertResult, err := collection.InsertOne(ctx, userMap)
	if err != nil {
		return nil, "", err
	}

	return insertResult, token, nil
}

func findUser(mg *mongo.Client, otp string) (User, error) {

	fmt.Println("ini otp user", otp)
	collection := ConnectionMongo(mg)
	otpTrim := strings.TrimSpace(otp)
	filter := bson.D{{Key: "otp_code", Value: otpTrim}}

	fmt.Printf("Filter Type: %T, Filter Value: %v\n", filter, filter)

	var users User
	err := collection.FindOne(ctx, filter).Decode(&users)
	if err != nil {
		fmt.Println("ini error:", err)
		return User{}, err
	}

	fmt.Printf("ini otp : %s", users.ID)

	err = verificationOTP(users, otpTrim)
	if err != nil {
		return User{}, err
	}
	return users, nil
}

func verificationOTP(result User, otp string) error {

	fmt.Println("ini otp mongo", result.OtpCode)
	fmt.Println("ini otp input", result.ExpiredAt)
	if result.ExpiredAt.Before(time.Now()) {
		return errors.New("OTP anda telah kadaluarsa")
	}

	if result.OtpCode != otp {
		return errors.New("kode OTP anda salah")
	}

	return nil
}

func ActivationAccount(db *gorm.DB, mg *mongo.Client, otp string) error {
	result, err := findUser(mg, otp)
	if err != nil {
		return err
	}

	res := db.Create(&result)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
