package main

import (
	"fmt"
	"log"
	"sample/application/repositories"
	"sample/domain"
	"sample/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{

		Name:     "Manoel",
		Email:    "manoel@jp.com",
		Password: "998",
	}

	user1 := domain.User{

		Name:     "Maria",
		Email:    "maria@jp.com",
		Password: "456",
	}

	user2 := domain.User{

		Name:     "Jose",
		Email:    "jose@jp.com",
		Password: "789",
	}

	userRepo := repositories.UserRepositoryDB{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	result, err = userRepo.Insert(&user1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	result, err = userRepo.Insert(&user2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
