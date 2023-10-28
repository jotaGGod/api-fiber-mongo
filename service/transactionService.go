package service

import (
	"api-fiber-mongo/entity"
	"api-fiber-mongo/repository"
)

func GetAccounts() []entity.Account {
	return []entity.Account{*repository.TransactionRepository.GetAccountDb("1")}
}

func MakeDeposits(num int) *entity.Account {
	conta := repository.TransactionRepository.GetAccountDb("1")

	conta.Balance += num
	repository.TransactionRepository.MakeDepositDb(conta)

	return conta
}

func MakeWithdraw(value int) *entity.Account {
	// qual o motivo da função calculo estar sendo chamada aqui sendo q seu retorno n está sendo utilizado?
	Calculo(value)

	contas := GetAccounts()

	for _, conta := range contas {
		conta.Balance -= value
		repository.TransactionRepository.MakeWithdrawDb(&conta)
		return &conta
	}

	return nil
}

func Calculo(valor int) map[int]int {
	var NotasTotais = map[int]int{200: 0, 100: 0, 50: 0, 20: 0, 10: 0, 5: 0, 2: 0}
	var notasReais = []int{200, 100, 50, 20, 10, 5, 2}
	// será q essa lógica n poderia ser melhorada em questão de performance??
	// dica: busque debugar e entender passo a passo para idenificar possíveis melhorias.
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
