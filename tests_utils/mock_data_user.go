package tests_utils

import (
	glide "github.com/retitle/go-sdk/v3"
)

func Address1() *glide.Address {
	return &glide.Address{
		Street:  "185 BERRY ST",
		Unit:    "STE 6600",
		City:    "SAN FRANCISCO",
		State:   "CA",
		ZipCode: "94107",
	}
}

func UserWithAddress() *glide.User {
	return &glide.User{
		Id:           "32",
		AgentAddress: Address1(),
		Contact:      ContactWithoutAddress(),
		Uuid:         "123e4567-e89b-12d3-a456-426614174000",
	}
}

func BillingInfoResponse() *glide.UserBillingInfo {
	return &glide.UserBillingInfo{
		StripeCustomerId: "123",
	}
}

func UserUpsertPayload() glide.UserManagementSchema {
	return glide.UserManagementSchema{
		Email:           "upsert_user@compass.com",
		FirstName:       "FirstName",
		LastName:        "LastName",
		LinkedSubjectId: "abc123",
	}
}
