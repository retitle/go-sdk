package interfaces

import (
	"net/http"

	"github.com/retitle/go-sdk/glide_errors"
	"github.com/retitle/go-sdk/http_utils"
	"github.com/retitle/go-sdk/models"
)

type Client interface {
	Get(res models.Response, authRequired bool, path string, headers *http.Header, qParams *http_utils.QueryParams) *glide_errors.ApiError
	Post(res models.Response, authRequired bool, path string, headers *http.Header, qParams *http_utils.QueryParams, payload models.Request) *glide_errors.ApiError
}
