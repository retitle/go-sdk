# Go Glide API SDK

> :warning: **The SDK and underlaying API are still being developed and current APIs are bound to change, not respecting backward compatibility conventions on versioning**

This SDK is meant to be used to more easily consume Glide's external APIs while using Go after an integration app has been set up for your Brokerage at Glide.

The underlying API's spec is available for both development and production environments on the links below:
* API root url: `https://api.dev.glide.com/` (development), & `https://api.glide.com/` (production)
* [Swagger docs](https://api.dev.glide.com/apidocs/)
* [Open API Spec](https://api.dev.glide.com/apispec_1.json)

## Integration App

If you don't have an integration app setup with Glide, please reach out to [Glide Engineering (engineering@glide.com)](mailto:engineering@glide.com) for guidance.
The integration app you use will dictate the client key and RS256 key pair values, and the associated brokerage's members are the Glide users you are allowed to impersonate while using these keys.

## Instalation

```bash
go get github.com/retitle/go-sdk
```

## Example usage

```go
package my_package

import (
	"github.com/retitle/go-sdk"
	"github.com/retitle/go-sdk/security"

	"fmt"
)

func main() {
    clientKey := "12345678-9abc-def0-1234-56789abcdef0"
    key := security.GetRsa256Key("/keys/private.pem")
    /*
        Also posible to get PEM formatted key from memory using
        either `security.GetRsa256KeyFromPEMBytes` (recives []byte)
        or `security.GetRsa256KeyFromPEMString` (recives string)
    */
    glideClient := glide.GetClient(clientKey, key, &glide.Options{
        Server: "api.dev.glide.com",
    }) // can send `nil` or empty `&glide.Options{}` for production config, or set Server to "api.glide.com"

    // This will fail because no user is being impersonated
    if _, err := glideClient.Users.Current(); err != nil {
        fmt.Println("This error is expected: ", err)
    }

    aGlideUsername := "your.user@domain.com"
    scopes := []string{}
    /*
        While impersonating a user, the SDK will seamlessly keep the access token refreshed.
        To stop impersonating you can use `glideClient.StopImpersonating()`, or you can use
        `glideClient.StartImpersonating(...)` with different parameters to just change the current impersonated user/scopes.
        At any point in time
        * `glideClient.IsImpersonating()`
        * `glideClient.ImpersonatingSub()`
        * `glideClient.ImpersonatingScopes()`
        can be used to find out whether or not an impersonation session is active, and find out details about it.
    */
    if err := glideClient.StartImpersonating(aGlideUsername, scopes); err != nil {
        panic(err)
    }

    user, err := glideClient.Users.Current();
    if err != nil {
        panic(err)
    }
    fmt.Println(*user)

    // This will fail because accessed resource (Transactions) requires missing TRANSACTIONS scope
    if _, err := glideClient.Transactions.List(); err != nil {
        fmt.Println("This error is expected: ", err)
    }

    scopes = []string{"TRANSACTIONS"}
    if err := glideClient.StartImpersonating(aGlideUsername, scopes); err != nil {
        panic(err)
    }

    txns, err := glideClient.Transactions.List()
	if err != nil {
		panic(err)
	}
	fmt.Println(*txns)

    transactionIds := []string{"2246", "1486", "490"} // Make sure all these exist, else a 404 error will occur
	txns, err = glideClient.Transactions.GetMulti(transactionIds)
	if err != nil {
		panic(err)
	}
	fmt.Println(*txns)

    transactionId := transactionIds[0]
	txn, err := glideClient.Transactions.GetDetail(transactionId) // Make sure this exists, else a 404 error will occur
	if err != nil {
		panic(err)
	}
	fmt.Println(*txn)
}
```
