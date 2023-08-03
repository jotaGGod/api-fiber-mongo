package entity

type Account struct {
	ID      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Balance int    `json:"balance" bson:"balance"`
}

type Result struct {
	Status      string
	ValorSaque  int
	NotasUsadas map[int]int
}
