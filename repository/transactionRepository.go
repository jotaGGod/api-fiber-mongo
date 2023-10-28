package repository

import (
	"api-fiber-mongo/database"
	"api-fiber-mongo/entity"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var TransactionRepository ITransactionRepository = &transactionRepository{CollectionName: "userAccount"}

type ITransactionRepository interface {
	GetAccountDb(id string) *entity.Account
	MakeDepositDb(conta *entity.Account) *entity.Account
	MakeWithdrawDb(conta *entity.Account) *entity.Account
}

type transactionRepository struct {
	CollectionName string
}

func (repo *transactionRepository) GetAccountDb(id string) *entity.Account {
	collection := database.GetCollection(repo.CollectionName)

	var conta entity.Account

	if err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&conta); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &conta
}

func (repo *transactionRepository) MakeDepositDb(conta *entity.Account) *entity.Account { //Realiza a alteração e realização de depósito no db
	collection := database.GetCollection(repo.CollectionName)

	update := bson.M{
		"$set": bson.M{"balance": conta.Balance},
	}

	if _, err := collection.UpdateOne(context.TODO(), bson.M{"id": "1"}, update); err != nil {
		log.Println(err.Error())
	}

	return conta
}

func (repo *transactionRepository) MakeWithdrawDb(conta *entity.Account) *entity.Account {
	collection := database.GetCollection(repo.CollectionName)

	update := bson.M{
		"$set": bson.M{
			"balance": conta.Balance,
		},
	}

	if _, err := collection.UpdateOne(context.TODO(), bson.M{"id": "1"}, update); err != nil {
		log.Println(err.Error())
	}

	return conta
}
