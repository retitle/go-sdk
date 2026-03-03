package fixtures

import (
	glide "github.com/retitle/go-sdk/v4"
	"github.com/retitle/go-sdk/v4/core"
)

func PartyData() *glide.Party {
	return &glide.Party{
		Id:     "LISTING ID",
		UserId: "Goku",
	}
}

func PartyListData() *glide.PartyList {
	return &glide.PartyList{
		Data:       []glide.Party{*PartyData()},
		ListObject: "party object",
		Object:     "Object",
		HasMore:    false,
	}
}

func PartyError() core.ErrorObject {
	return core.ErrorObject{
		Message: "ERROR GETTING PARTIES",
		Object:  "ERROR OBJECT PARTIES",
	}
}
