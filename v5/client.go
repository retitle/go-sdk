package glide

import (
	"fmt"
	"strings"
	"time"

	"github.com/retitle/go-sdk/v5/core"
)

type Client interface {
	Get(res core.Response, authRequired bool, path string, opts ...core.RequestOption) error
	Post(res core.Response, authRequired bool, path string, payload core.Request, opts ...core.RequestOption) error
	PostWithFiles(res core.Response, authRequired bool, path string, payload core.Request, files []core.File, opts ...core.RequestOption) error
	GetOptions() core.ClientOptions
	SetHttpClient(httpClient core.HttpClient)
	StartImpersonating(sub string, scopes []string) error
	IsImpersonating() bool
	ImpersonatingSub() string
	ImpersonatingScopes() []string
	StopImpersonating()
	// DO NOT remove these comments since they serve as anchors for code autogeneration
	/* Autogenerated-root-resource-interface-defs begins */
	Contacts() ContactsResource
	Documents() DocumentsResource
	Listings() ListingsResource
	Notifications() NotificationsResource
	Transactions() TransactionsResource
	UserManagement() UserManagementResource
	Users() UsersResource
	/* Autogenerated-root-resource-interface-defs ends */
}

type ClientImpl struct {
	httpClient    core.HttpClient
	clientKey     string
	key           core.Key
	options       core.ClientOptions
	impersonating core.Impersonation

	// DO NOT remove these comments since they serve as anchors for code autogeneration
	/* Autogenerated-root-resource-defs begins */
	contacts       ContactsResource
	documents      DocumentsResource
	listings       ListingsResource
	notifications  NotificationsResource
	transactions   TransactionsResource
	userManagement UserManagementResource
	users          UsersResource
	/* Autogenerated-root-resource-defs ends */
}

const JWT_EXPIRES = 60

func GetClient(clientKey string, key core.Key, opts ...core.ClientOption) Client {
	defaultOptions := core.NewClientOptions().SetProtocol("https").
		SetHost("api.glide.com").
		SetUrl("").
		SetBasePath("").
		SetAudience("")

	c := ClientImpl{
		httpClient: core.NewHttpClient(),
		clientKey:  clientKey,
		key:        key,
		options:    defaultOptions,
	}
	for _, option := range opts {
		option(c.GetOptions())
	}
	if c.GetOptions().GetUrl() == "" {
		c.GetOptions().SetUrl(c.GetOptions().GetHost())
	}
	if c.GetOptions().GetAudience() == "" {
		c.GetOptions().SetAudience(c.GetOptions().GetHost())
	}
	// DO NOT remove these comments since they serve as anchors for code autogeneration
	/* Autogenerated-root-resource-init begins */
	c.contacts = GetContactsResource(&c)
	c.documents = GetDocumentsResource(&c)
	c.listings = GetListingsResource(&c)
	c.notifications = GetNotificationsResource(&c)
	c.transactions = GetTransactionsResource(&c)
	c.userManagement = GetUserManagementResource(&c)
	c.users = GetUsersResource(&c)
	/* Autogenerated-root-resource-init ends */
	return &c
}

func (c *ClientImpl) GetOptions() core.ClientOptions {
	return c.options
}

// DO NOT remove these comments since they serve as anchors for code autogeneration
/* Autogenerated-root-resource-getters begins */
func (c *ClientImpl) Contacts() ContactsResource {
	return c.contacts
}
func (c *ClientImpl) Documents() DocumentsResource {
	return c.documents
}
func (c *ClientImpl) Listings() ListingsResource {
	return c.listings
}
func (c *ClientImpl) Notifications() NotificationsResource {
	return c.notifications
}
func (c *ClientImpl) Transactions() TransactionsResource {
	return c.transactions
}
func (c *ClientImpl) UserManagement() UserManagementResource {
	return c.userManagement
}
func (c *ClientImpl) Users() UsersResource {
	return c.users
}

/* Autogenerated-root-resource-getters ends */

func (c *ClientImpl) getUrl(path string) string {
	return fmt.Sprintf("%s://%s%s%s", strings.ToLower(c.GetOptions().GetProtocol()),
		strings.ToLower(c.GetOptions().GetUrl()),
		strings.ToLower(c.GetOptions().GetBasePath()),
		strings.ToLower(path),
	)
}

func (c *ClientImpl) request(authRequired bool, doRequest func(extraOpts ...core.RequestOption) error) error {
	if authRequired && c.IsImpersonating() {
		c.refreshAccessToken()
	}

	extraOpts := []core.RequestOption{
		core.WithRequestHost(c.GetOptions().GetHost()),
	}

	if c.IsImpersonating() {
		extraOpts = append(extraOpts, core.WithHeader("Authorization", fmt.Sprintf("Bearer %s", c.impersonating.GetAccessToken())))
	}

	return doRequest(extraOpts...)
}

func (c *ClientImpl) Get(res core.Response, authRequired bool, path string, opts ...core.RequestOption) error {
	return c.request(authRequired, func(extraOpts ...core.RequestOption) error {
		return c.httpClient.Get(res, c.getUrl(path), append(opts, extraOpts...)...)
	})
}

func (c *ClientImpl) Post(res core.Response, authRequired bool, path string, payload core.Request, opts ...core.RequestOption) error {
	return c.request(authRequired, func(extraOpts ...core.RequestOption) error {
		return c.httpClient.Post(res, c.getUrl(path), payload, append(opts, extraOpts...)...)
	})
}

func (c *ClientImpl) PostWithFiles(res core.Response, authRequired bool, path string, payload core.Request, files []core.File, opts ...core.RequestOption) error {
	return c.request(authRequired, func(extraOpts ...core.RequestOption) error {
		return c.httpClient.PostWithFiles(res, c.getUrl(path), payload, files, append(opts, extraOpts...)...)
	})
}

func (c *ClientImpl) getJwt(sub string, scopes []string) (string, error) {
	return core.GetJwt(c.key, c.clientKey, sub, c.GetOptions().GetAudience(), scopes, JWT_EXPIRES)
}

func (c *ClientImpl) getAccessToken(sub string, scopes []string) (string, error) {
	jwt, err := c.getJwt(sub, scopes)
	if err != nil {
		return "", &core.ApiErrorImpl{
			Params: map[string]interface{}{
				"desc": "Error issuing assertions JWT",
			},
			Err: err,
		}
	}

	return jwt, nil
}

func (c *ClientImpl) StartImpersonating(sub string, scopes []string) error {
	jwt, err := c.getAccessToken(sub, scopes)
	if err != nil {
		return err
	}
	c.impersonating = core.NewImpersonation().SetSub(sub).
		SetScopes(scopes).SetAccessToken(jwt).
		SetAccessTokenExpires(time.Now().
			Add(time.Second * time.Duration(JWT_EXPIRES)))

	return nil
}

func (c *ClientImpl) IsImpersonating() bool {
	return c.impersonating != nil
}

func (c *ClientImpl) ImpersonatingSub() string {
	if c.IsImpersonating() {
		return c.impersonating.GetSub()
	}
	return ""
}

func (c *ClientImpl) ImpersonatingScopes() []string {
	if c.IsImpersonating() {
		return c.impersonating.GetScopes()
	}
	return []string{}
}

func (c *ClientImpl) StopImpersonating() {
	c.impersonating = nil
}

func (c *ClientImpl) isTokenExpired() bool {
	return time.Now().After(c.impersonating.GetAccessTokenExpires())
}

func (c *ClientImpl) isTokenAboutToExpire() bool {
	// Consider an access token "about to expire" if it's set to expire in less than 15m
	return !c.isTokenExpired() && time.Now().After(c.impersonating.GetAccessTokenExpires().Add(time.Minute*time.Duration(-15)))
}

func (c *ClientImpl) shouldRefreshToken() bool {
	return c.IsImpersonating() && (c.isTokenAboutToExpire() || c.isTokenExpired())
}

func (c *ClientImpl) refreshAccessToken() error {
	if c.shouldRefreshToken() {
		jwt, err := c.getAccessToken(c.impersonating.GetSub(), c.impersonating.GetScopes())

		if err != nil {
			// fail silently if token is not yet expired
			if c.isTokenExpired() {
				return err
			}
		} else {
			c.impersonating.SetAccessToken(jwt)
			c.impersonating.SetAccessTokenExpires(time.Now().Add(time.Second * time.Duration(JWT_EXPIRES)))
		}
	}

	return nil
}

func (c *ClientImpl) SetHttpClient(httpClient core.HttpClient) {
	c.httpClient = httpClient
}