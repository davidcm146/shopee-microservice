package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Email       string    `bson:"email" json:"email"`
	Password    string    `bson:"password,omitempty"`
	PhoneNumber string    `bson:"phoneNumber" json:"phoneNumber"`
	Address     []string  `bson:"address" json:"address"`
	Role        string    `bson:"role" json:"role"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
}

func NewUser(email, password, name, role string) *User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{Email: email, Name: name, Role: role, Password: string(hashed)}
}

func (u *User) CheckPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw)) == nil
}
