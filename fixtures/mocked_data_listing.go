package fixtures

import (
	glide "github.com/retitle/go-sdk/v3"
	"github.com/retitle/go-sdk/v3/core"
)

func ListingData() *glide.Listing {
	return &glide.Listing{
		Id:      "LISTING ID",
		MlsKind: "SOME MLS",
	}
}

func ListingListData() *glide.ListingList {
	return &glide.ListingList{
		Data:       []glide.Listing{*ListingData()},
		ListObject: "list object",
		Object:     "Object",
		HasMore:    false,
	}
}

func ListingError() core.ErrorObject {
	return core.ErrorObject{
		Message: "ERROR SENDING EMAIL",
		Object:  "ERROR OBJECT EMAIL",
	}
}
