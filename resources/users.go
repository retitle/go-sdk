package resources

import (
	"fmt"

	"github.com/retitle/go-sdk/glide_errors"
	"github.com/retitle/go-sdk/interfaces"
	"github.com/retitle/go-sdk/models"
)

type UsersResource struct {
	GlideClient interface{}
}

func GetUsersResource(glideClient interface{}) UsersResource {
	return UsersResource{
		GlideClient: glideClient,
	}
}

func (r UsersResource) client() interfaces.Client {
	return r.GlideClient.(interfaces.Client)
}

func (r UsersResource) Current() (*models.User, *glide_errors.ApiError) {
	res := models.User{}
	if err := r.client().Get(&res, true, fmt.Sprintf("/users/current"), nil, nil); err != nil {
		return nil, err
	}
	return &res, nil
}
