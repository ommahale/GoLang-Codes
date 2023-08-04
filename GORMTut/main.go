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
L:
	for {
		fmt.Print("Please insert command:")

		var input string
		fmt.Scanln(&input)
		switch input {
		case "create":
			var username string
			var passowrd string
			fmt.Print("Username:")
			fmt.Scanln(&username)
			fmt.Print("Password:")
			fmt.Scanln(&passowrd)
			user := db_types.GetNewUser(username, passowrd)
			r.userRepo.CreateUser(*user, *newDB)
		case "login":
			var username string
			var passowrd string
			fmt.Print("Username:")
			fmt.Scanln(&username)
			fmt.Print("Password:")
			fmt.Scanln(&passowrd)
			db_types.HashPassword(&passowrd)
			user := &db_types.User{Username: username, Password: passowrd}
			r.userRepo.FindUser(*user, *newDB)
		case "exit":
			break L
		}
	}
}
