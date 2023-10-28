package repository

import (
	"api-fiber-mongo/database"
	"api-fiber-mongo/entity"
	"context"
	"log"
)

var AccountRepository IAccountRepository = &accountRepository{CollectionName: "userAccount"}

type IAccountRepository interface {
	CreateAccount(conta *entity.Account) *entity.Account
}

type accountRepository struct {
	CollectionName string
}

func (repo *accountRepository) CreateAccount(conta *entity.Account) *entity.Account { //Criar conta no banco de dados
	collection := database.GetCollection(repo.CollectionName)

	if _, err := collection.InsertOne(context.TODO(), conta); err != nil {
		log.Println(err.Error())
	}

	return conta
}
