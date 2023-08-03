package service

import (
	"api-fiber-mongo/entity"
	"api-fiber-mongo/repository"
)

func GetAccounts() []entity.Account {
	return []entity.Account{repository.AccountRepository.GetAccountDb()}
}

func MakeDeposits(num int) *entity.Account {

	contas := GetAccounts()

	for _, conta := range contas {
		conta.Balance += num
		repository.AccountRepository.MakeDepositDb(&conta)
		return &conta
	}
	return nil
}

func CreateAccounts(conta *entity.Account) *entity.Account {
	return repository.AccountRepository.CreateAccount(conta)
}

func MakeWithdraw(value int) *entity.Account {

	Calculo(value)

	contas := GetAccounts()

	for _, conta := range contas {
		conta.Balance -= value
		repository.AccountRepository.MakeWithdrawDb(&conta)
		return &conta
	}

	return nil
}

func Calculo(valor int) map[int]int {
	var NotasTotais = map[int]int{200: 0, 100: 0, 50: 0, 20: 0, 10: 0, 5: 0, 2: 0} //logica do retorno do valor de notas a ser sacado
	var notasReais = []int{200, 100, 50, 20, 10, 5, 2}
	for _, nota := range notasReais {
		if valor-nota != 1 && valor-nota != 3 {
			for valor-nota >= 0 {
				valor -= nota
				(NotasTotais)[nota] += 1
			}
		}
	}
	return NotasTotais
}
