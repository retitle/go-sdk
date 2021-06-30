package models

type Transaction struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type TransactionList struct {
	Data       []Transaction `json:"data"`
	ListObject string        `json:"list_object"`
	Object     string        `json:"object"`
}
