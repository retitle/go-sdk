package models

type Contact struct {
	Agent      Agent  `json:"agent"`
	CellPhone  string `json:"cell_phone"`
	Email      string `json:"email"`
	EntityName string `json:"entity_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Title      string `json:"title"`
}
