package users

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	encryption "github.com/seronz/api/src/utils/Encryption"
	jwt "github.com/seronz/api/src/utils/JWT"
	otp "github.com/seronz/api/src/utils/OTP"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type User struct {
	ID            string `bson:"_id" gorm:"size:255;not null;uniqueIndex;primary_key" json:"id"`
	Users         []Address
	FirstName     string    `bson:"first_name" gorm:"size:100;not null" json:"first_name"`
	LastName      string    `bson:"last_name" gorm:"size:100;not null" json:"last_name"`
	Email         string    `bson:"email" gorm:"size:100;not null" json:"email"`
	Password      string    `bson:"password" gorm:"size:100;not null" json:"password"`
	Salt          string    `bson:"salt" gorm:"size:100;not null" json:"salt"`
	UserRole      string    `bson:"user_role" gorm:"size:10;not null;default:'user'" json:"user_role"`
	OtpCode       string    `bson:"otp_code" gorm:"size:10" json:"otp_code"`
	IsActive      bool      `bson:"is_active" gorm:"type:boolean;default:true" json:"is_active"`
	RememberMe    bool      `bson:"remember_me" gorm:"type:boolean;default:false" json:"remember_me"`
	RememberToken string    `bson:"remember_token" gorm:"type:text;not null" json:"remember_token"`
	CreatedAt     time.Time `bson:"CreatedAt" json:"created_at"`
	UpdatedAt     time.Time `bson:"UpdatedAt" json:"updated_at"`
	ExpiredAt     time.Time `bson:"expired_at" json:"expired_at"`
	DeletedAt     gorm.DeletedAt
}

type MongoParam struct {
	update bson.D
	filter bson.D
}

var ctx = context.TODO()

func ConnectionMongo(mg *mongo.Client) *mongo.Collection {
	collection := mg.Database("tokopedia").Collection("user_not_active")
	return collection
}

func UserRegister(mg *mongo.Client, user User) (*mongo.InsertOneResult, string, string, error) {
	param := jwt.Params{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		Userrole:  user.UserRole,
	}
	token, err := jwt.CreateToken(param)
	if err != nil {
		return nil, "", "", err
	}

	otp, err := otp.GenerateOTP(user.Email)
	if err != nil {
		return nil, "", "", err
	}

	pw, salt, err := encryption.EncryptPassword(user.Password)
	if err != nil {
		return nil, "", "", err
	}

	expireTime := time.Now().Add(2 * time.Hour)
	user.Password = pw
	user.Salt = salt
	user.OtpCode = otp
	user.ExpiredAt = expireTime

	userMap := make(map[string]interface{})
	userBytes, err := json.Marshal(user)
	if err != nil {
		return nil, "", "", err
	}
	if err := json.Unmarshal(userBytes, &userMap); err != nil {
		return nil, "", "", err
	}

	collection := ConnectionMongo(mg)
	emailIndex := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(ctx, emailIndex)
	if err != nil {
		return nil, "", "", err
	}

	insertResult, err := collection.InsertOne(ctx, userMap)
	if err != nil {
		return nil, "", "", err
	}

	return insertResult, token, otp, nil
}

func findUser(mg *mongo.Client, otp string) (User, error) {
	collection := ConnectionMongo(mg)
	otpTrim := strings.TrimSpace(otp)
	filter := bson.D{{Key: "otp_code", Value: otpTrim}}

	var users User
	err := collection.FindOne(ctx, filter).Decode(&users)
	if err != nil {
		fmt.Println("ini error:", err)
		return User{}, err
	}

	err = verificationOTP(users, otpTrim)
	if err != nil {
		return User{}, err
	}
	return users, nil
}

func verificationOTP(result User, otp string) error {
	if result.ExpiredAt.Before(time.Now()) {
		return errors.New("OTP anda telah kadaluarsa")
	}

	if result.OtpCode != otp {
		return errors.New("kode OTP anda salah")
	}

	return nil
}

func deleteDataMongo(mg *mongo.Client, otp string) error {
	collection := ConnectionMongo(mg)

	trimOtp := strings.TrimSpace(otp)

	filter := bson.D{{Key: "otp_code", Value: trimOtp}}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func ActivationAccount(db *gorm.DB, mg *mongo.Client, otp string) error {

	result, err := findUser(mg, otp)
	if err != nil {
		return err
	}

	log.Printf("Activate Account : %s %s ", result.FirstName, result.LastName)

	res := db.Create(&result)
	if res.Error != nil {
		return res.Error
	}

	err = deleteDataMongo(mg, otp)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoParam) updateOTP(mg *mongo.Client) error {
	collection := ConnectionMongo(mg)

	_, err := collection.UpdateOne(ctx, m.filter, m.update)
	if err != nil {
		return err
	}
	return nil
}

func RegenerateOTP(mg *mongo.Client, token string) (string, error) {
	claims, err := jwt.JWTGetClaims(token)
	if err != nil {
		return "", err
	}

	otp, err := otp.GenerateOTP(claims.Email)
	if err != nil {
		return "", err
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "otp_code", Value: otp},
		}},
	}

	filter := bson.D{{Key: "email", Value: claims.Email}}

	var m MongoParam
	m.update = update
	m.filter = filter
	err = m.updateOTP(mg)
	if err != nil {
		return "", err
	}

	return otp, nil
}
