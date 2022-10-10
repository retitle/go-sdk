package fixtures

import (
	glide "github.com/retitle/go-sdk/v3"
	"github.com/retitle/go-sdk/v4/core"
)

func TransactionDocumentData() *glide.TransactionDocument {
	return &glide.TransactionDocument{
		Id:    "LISTING ID",
		Title: "Goku's document",
	}
}

func TransactionDocumentListData() *glide.TransactionDocumentList {
	return &glide.TransactionDocumentList{
		Data:       []glide.TransactionDocument{*TransactionDocumentData()},
		ListObject: "trx doc object",
		Object:     "Object",
		HasMore:    false,
	}
}

func TransactionDocumentError() core.ErrorObject {
	return core.ErrorObject{
		Message: "ERROR GETTING TRX DOCS",
		Object:  "ERROR OBJECT TRX DOCS",
	}
}
