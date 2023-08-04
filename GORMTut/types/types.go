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
	HashPassword(&password)
	id := uuid.New()
	return &User{
		Uuid:     id.String(),
		Username: username,
		Password: password,
	}
}
func HashPassword(password *string) {
	h := sha256.New()
	h.Write([]byte(*password))
	*password = base64.URLEncoding.EncodeToString(h.Sum(nil))
}
