package resources

import (
	"fmt"
	"strings"

	"github.com/retitle/go-sdk/glide_errors"
	"github.com/retitle/go-sdk/http_utils"
	"github.com/retitle/go-sdk/interfaces"
	"github.com/retitle/go-sdk/models"
)

type TransactionsResource struct {
	GlideClient interface{}
	Parties     PartiesResource
}

func GetTransactionsResource(glideClient interface{}) TransactionsResource {
	return TransactionsResource{
		GlideClient: glideClient,
		Parties:     GetPartiesResource(glideClient),
	}
}

func (r TransactionsResource) client() interfaces.Client {
	return r.GlideClient.(interfaces.Client)
}

func (r TransactionsResource) GetDetail(id string) (*models.Transaction, *glide_errors.ApiError) {
	res := models.Transaction{}
	if err := r.client().Get(&res, true, fmt.Sprintf("/transactions/%s", id), nil, nil); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) GetMulti(ids []string) (*models.TransactionList, *glide_errors.ApiError) {
	res := models.TransactionList{}
	if err := r.client().Get(&res, true, fmt.Sprintf("/transactions"), nil, &http_utils.QueryParams{
		"ids": strings.Join(ids, ","),
	}); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) List() (*models.TransactionList, *glide_errors.ApiError) {
	res := models.TransactionList{}
	if err := r.client().Get(&res, true, fmt.Sprintf("/transactions"), nil, nil); err != nil {
		return nil, err
	}
	return &res, nil
}
