package users

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type User struct {
	Id            string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	Users         []Address
	FirstName     string `gorm:"size:100;not null" json:"first_name"`
	LastName      string `gorm:"size:100;not null" json:"last_name"`
	Email         string `gorm:"size:100;not null" json:"email"`
	Password      string `gorm:"size:100;not null" json:"password"`
	UserRole      string `gorm:"size:10;not null;default:'guest'" json:"user_role"`
	OtpCode       string `gorm:"size:10" json:"otp_code"`
	IsActive      bool   `json:"is_active"`
	RememberToken string `gorm:"size:255;not null" json:"remember_token"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

var ctx = context.TODO()

func UserRegister(db *gorm.DB, mongo *mongo.Client, user User) (*mongo.InsertOneResult, error) {
	collection := mongo.Database("tokopedia").Collection("user_not_active")
	userMap := make(map[string]interface{})
	userBytes, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(userBytes, &userMap); err != nil {
		return nil, err
	}

	insertResult, err := collection.InsertOne(ctx, userMap)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}
