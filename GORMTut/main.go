package main

import (
	"fmt"
	"log"

	"github.com/ommahale/gormtut/db"
	"github.com/ommahale/gormtut/db/repo"
	db_types "github.com/ommahale/gormtut/types"
)

type repos struct {
	userRepo repo.UserRepo
}

func initRepo() *repos {
	userRepo := repo.UserRepo{}
	return &repos{
		userRepo: userRepo,
	}
}

func main() {
	newDB := db.Connect()
	log.Println("db connected")
	db.InitModels(newDB)
	log.Println("db initialized")
	r := initRepo()
	for {
		fmt.Println("Please insert command:")

		var input string
		fmt.Scanln(&input)
		switch input {
		case "create":
			var username string
			var passowrd string
			fmt.Scanln(&username)
			fmt.Scanln(&passowrd)
			user := db_types.GetNewUser(username, passowrd)
			r.userRepo.CreateUser(*user, *newDB)
		case "login":
			var username string
			var passowrd string
			fmt.Scanln(&username)
			fmt.Scanln(&passowrd)
			user := db_types.GetNewUser(username, passowrd)
			r.userRepo.FindUser(*user, *newDB)
		}
	}
}
