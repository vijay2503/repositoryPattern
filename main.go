package main

import (
	"fmt"
	"go-postgres/driver"
	models "go-postgres/model"
	"go-postgres/repository"
	"go-postgres/repository/repoimpl"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "password"
	dbname   = "repo_pattern"
)

var userRepo repository.UserRepo

func init() {
	db := driver.Connect(host, port, user, password, dbname)
	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}
	userRepo = repoimpl.NewUserRepo(db.SQL)
}
func main() {
	Userst := []models.User{}
	choice := 0
loop:
	for {
		fmt.Print("\n")
		fmt.Println("Enter your Choice")
		fmt.Print("\n")
		fmt.Println("Press 1 for Insert ")
		fmt.Println("Press 2 for Read ")
		fmt.Println("Press 3 for Update")
		fmt.Println("Press 4 for Delete")
		fmt.Println("Press 5 for Exit")
		fmt.Print("\n")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			userRepo.Create(&Userst)
		case 2:
			userRepo.Select()
		case 3:
			userRepo.Update()
		case 4:
			userRepo.Delete()
		default:
			fmt.Println(" Exit !!! ")
			break loop
		}
	}

}
