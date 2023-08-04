package db_types

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid     string `gorm:"index"`
	Username string
	Password string
}

func GetNewUser(username string, password string) *User {
	h := sha256.New()
	h.Write([]byte(password))
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))
	id := uuid.New()
	return &User{
		Uuid:     id.String(),
		Username: username,
		Password: hash,
	}
}
