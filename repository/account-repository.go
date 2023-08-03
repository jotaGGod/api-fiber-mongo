package repository

import (
	"api-fiber-mongo/database"
	"api-fiber-mongo/entity"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var AccountRepository IAccountRepository = &accountRepository{}

type IAccountRepository interface {
	GetAccountDb() entity.Account
	MakeDepositDb(conta *entity.Account) *entity.Account
	CreateAccount(conta *entity.Account) *entity.Account
	MakeWithdrawDb(conta *entity.Account) *entity.Account
}

type accountRepository struct {
	ID      string
	Name    string
	Balance int
}

func (repo accountRepository) GetAccountDb() entity.Account {
	database := database.MongoDB.Database("Banco-Atlas").Collection("userAccount")

	var conta entity.Account

	if err := database.FindOne(context.TODO(), bson.M{"id": "1"}).Decode(&conta); err != nil {
		log.Println(err.Error())
	}

	return conta
}

func (repo *accountRepository) MakeDepositDb(conta *entity.Account) *entity.Account { //Realiza a alteração e realização de depósito no db
	database := database.MongoDB.Database("Banco-Atlas").Collection("userAccount")

	update := bson.M{
		"$set": bson.M{"balance": conta.Balance},
	}

	if _, err := database.UpdateOne(context.TODO(), bson.M{"id": "1"}, update); err != nil {
		log.Println(err.Error())
	}

	return conta
}

func (repo *accountRepository) CreateAccount(conta *entity.Account) *entity.Account { //Criar conta no banco de dados
	database := database.MongoDB.Database("Banco-Atlas").Collection("userAccount")

	if _, err := database.InsertOne(context.TODO(), conta); err != nil {
		log.Println(err.Error())
	}

	return conta
}

func (repo *accountRepository) MakeWithdrawDb(conta *entity.Account) *entity.Account {
	database := database.MongoDB.Database("Banco-Atlas").Collection("userAccount")

	update := bson.M{
		"$set": bson.M{"balance": conta.Balance},
	}

	if _, err := database.UpdateOne(context.TODO(), bson.M{"id": "1"}, update); err != nil {
		log.Println(err.Error())
	}

	return conta
}
