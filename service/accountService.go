package service

import (
	"api-fiber-mongo/entity"
	"api-fiber-mongo/repository"
)

func CreateAccounts(conta *entity.Account) *entity.Account {
	return repository.AccountRepository.CreateAccount(conta)
}
