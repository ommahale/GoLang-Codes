package repo

import (
	"log"

	db_types "github.com/ommahale/gormtut/types"
	"gorm.io/gorm"
)

type UserRepo struct {
}

func (u *UserRepo) CreateUser(user db_types.User, db gorm.DB) {
	db.Create(&user)
}
func (u *UserRepo) FindUser(user db_types.User, db gorm.DB) {
	result := db.Find(&user)
	if result.Error == gorm.ErrRecordNotFound {
		log.Fatal("User not found")
	}
	log.Println("User found, uuid: ", user.Uuid)
}
