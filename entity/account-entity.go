package entity

type Account struct {
	ID      string `json:"id" bson:"id"` // Id could be int or uuid. Usage of mongo objectId is an option also, but no that recomended.
	Name    string `json:"name" bson:"name"`
	Balance int    `json:"balance" bson:"balance"`
}

type Result struct {
	Status      string
	ValorSaque  int
	NotasUsadas map[int]int
}
