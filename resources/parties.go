package resources

import (
	"fmt"
	"strings"

	"github.com/retitle/go-sdk/glide_errors"
	"github.com/retitle/go-sdk/http_utils"
	"github.com/retitle/go-sdk/interfaces"
	"github.com/retitle/go-sdk/models"
)

type PartiesResource struct {
	GlideClient interface{}
}

func GetPartiesResource(glideClient interface{}) PartiesResource {
	return PartiesResource{
		GlideClient: glideClient,
	}
}

func (r PartiesResource) client() interfaces.Client {
	return r.GlideClient.(interfaces.Client)
}

func (r PartiesResource) GetDetail(transactionId string, id string) (*models.Party, *glide_errors.ApiError) {
	res := models.Party{}
	if err := r.client().Get(&res, true, fmt.Sprintf("/transactions/%s/parties/%s", transactionId, id), nil, nil); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r PartiesResource) GetMulti(transactionId string, ids []string) (*models.PartyList, *glide_errors.ApiError) {
	res := models.PartyList{}
	if err := r.client().Get(&res, true, fmt.Sprintf("/transactions/%s/parties", transactionId), nil, &http_utils.QueryParams{
		"ids": strings.Join(ids, ","),
	}); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r PartiesResource) List(transactionId string) (*models.PartyList, *glide_errors.ApiError) {
	res := models.PartyList{}
	if err := r.client().Get(&res, true, fmt.Sprintf("/transactions/%s/parties", transactionId), nil, nil); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r PartiesResource) UpdatePartyName(transactionId string, id string, partyUpdateName models.PartyUpdateName) *glide_errors.ApiError {
	if err := r.client().Post(nil, true, fmt.Sprintf("/transactions/%s/parties/%s/update_party_name", transactionId, id), nil, nil, partyUpdateName); err != nil {
		return err
	}
	return nil
}
