package models

type Agent struct {
	CompanyName   string `json:"company_name"`
	LicenseNumber string `json:"license_number"`
	NrdsNumber    string `json:"nrds_number"`
}
