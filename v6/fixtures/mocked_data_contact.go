package fixtures

import (
	glide "github.com/retitle/go-sdk/v6"
)

func ContactWithAddress() *glide.Contact {
	return &glide.Contact{
		Id:            "123",
		Address:       Address1(),
		CellPhone:     "987654321",
		Email:         "test_email@compass.com",
		FirstName:     "Contact1",
		LastName:      "ContactLast",
		Object:        "user",
		ContactSource: ContactSource(),
	}
}

func ContactRequest() *glide.ContactRequest {
	return &glide.ContactRequest{
		Address:       Address1(),
		CellPhone:     "987654321",
		Email:         "test_email@compass.com",
		FirstName:     "Contact1",
		LastName:      "ContactLast",
		ContactSource: ContactSource(),
	}
}

func ContactSource() *glide.ContactSource {
	return &glide.ContactSource{
		Id:     "123",
		Origin: "GLIDE",
		Object: "contact_source",
	}
}

func ContactCreate() *glide.ContactCreate {
	return &glide.ContactCreate{
		Contact: &glide.ContactRequest{
			Address:   Address1(),
			CellPhone: "987654321",
			Email:     "test_email@compass.com",
			FirstName: "Contact1",
			LastName:  "ContactLast",
		},
	}
}

func ContactCreateResponseData() *glide.ContactCreateResponse {
	return &glide.ContactCreateResponse{
		Contact: ContactWithAddress(),
	}
}

func ContactUpdateResponse() *glide.ContactUpdateResponse {
	return &glide.ContactUpdateResponse{
		Contact: ContactWithoutAddress(),
	}
}

func ContactUpdateRequest() glide.ContactUpdate {
	return glide.ContactUpdate{
		Contact: &glide.ContactRequest{
			Address:   Address1(),
			CellPhone: "987654321",
			Email:     "test_email@compass.com",
			FirstName: "Contact1",
			LastName:  "ContactLast",
		},
	}
}

func ContactWithoutAddress() *glide.Contact {
	return &glide.Contact{
		Id:        "321",
		CellPhone: "987654321",
		Email:     "test_email@compass.com",
		FirstName: "Contact2",
		LastName:  "ContactLast",
		Object:    "user",
	}
}

func ContactList() *glide.ContactList {
	return &glide.ContactList{
		Data: []glide.Contact{
			{
				Id:        "111",
				CellPhone: "987654321",
				Email:     "test_email@compass.com",
				FirstName: "Contact1",
				LastName:  "ContactLast",
				Object:    "user",
			},
			{
				Id:        "112",
				CellPhone: "987654321",
				Email:     "test_email@compass.com",
				FirstName: "Contact2",
				LastName:  "ContactLast",
				Object:    "user",
			},
			{
				Id:        "113",
				CellPhone: "987654321",
				Email:     "test_email@compass.com",
				FirstName: "Contact3",
				LastName:  "ContactLast",
				Object:    "user",
			},
		},
		HasMore: false,
	}
}
