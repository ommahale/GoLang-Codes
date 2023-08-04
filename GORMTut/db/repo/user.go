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
	var dest db_types.User
	result := db.Find(&dest, &user)
	if result.RowsAffected == 0 {
		log.Println("User not found")
		return
	}
	log.Println("User found, uuid: ", dest.Uuid)
}
