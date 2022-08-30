package tests_utils

import (
  glide "github.com/retitle/go-sdk/v3"
)

func ContactWithAddress() *glide.Contact {
  return &glide.Contact{
    Id:        "123",
    Address:   Address1(),
    CellPhone: "987654321",
    Email:     "test_email@compass.com",
    FirstName: "Contact1",
    LastName:  "ContactLast",
    Object:    "user",
  }
}

func ContactCreate() glide.ContactCreate {
  return glide.ContactCreate{
    Contact: &glide.ContactRequest{
      Address:   Address1(),
      CellPhone: "987654321",
      Email:     "test_email@compass.com",
      FirstName: "Contact1",
      LastName:  "ContactLast",
    },
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
      glide.Contact{
        Id:        "111",
        CellPhone: "987654321",
        Email:     "test_email@compass.com",
        FirstName: "Contact1",
        LastName:  "ContactLast",
        Object:    "user",
      },
      glide.Contact{
        Id:        "112",
        CellPhone: "987654321",
        Email:     "test_email@compass.com",
        FirstName: "Contact2",
        LastName:  "ContactLast",
        Object:    "user",
      },
      glide.Contact{
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
