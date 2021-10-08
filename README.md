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

    "fmt"
)

func main() {
    clientKey := "12345678-9abc-def0-1234-56789abcdef0"
    key := glide.GetRsa256Key("/keys/private.pem")
    /*
        Also posible to get PEM formatted key from memory using
        either `glide.GetRsa256KeyFromPEMBytes` (recives []byte)
        or `glide.GetRsa256KeyFromPEMString` (recives string)
    */
    glideClient := glide.GetClient(clientKey, key,
        glide.WithServer("api.dev.glide.com"),
    }) // exlcude WithServer option for prod (or use "api.glide.com")

    // This will fail because no user is being impersonated
    if _, err := glideClient.Users().Current(); err != nil {
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

    // From here on out, all requests will have impersonated user's objects, but no scopes were requested
    if err := glideClient.StartImpersonating(aGlideUsername, scopes); err != nil {
        panic(err)
    }

    // This works because there is an impersonated user, and the current user endpoint doesn't require any scope to be accessed
    user, err := glideClient.Users().Current()
    if err != nil {
        panic(err)
    }
    fmt.Println(*user)

    // This will fail because accessed resource (Transactions) requires missing TRANSACTIONS scope
    if _, err := glideClient.Transactions().List(); err != nil {
        fmt.Println("This error is expected: ", err)
    }

    // From here on out, all requests will have access to the TRANSACTIONS scope for impersonated user
    scopes = []string{"TRANSACTIONS"}
    if err := glideClient.StartImpersonating(aGlideUsername, scopes); err != nil {
        panic(err)
    }

    // List all transactions avaialble for impersonated user
    txns, err := glideClient.Transactions().List()
    if err != nil {
        panic(err)
    }
    fmt.Println(*txns)

    // Fetch multiple specific transactions avaialble for impersonated user
    transactionIds := []string{"2246", "1486", "490"} // Make sure all these exist, else a 404 error will occur
    txns, err = glideClient.Transactions().GetMulti(transactionIds)
    if err != nil {
        panic(err)
    }
    fmt.Println(*txns)

    // Fetch a single transactions avaialble for impersonated user
    transactionId := transactionIds[0]
    txn, err := glideClient.Transactions().GetDetail(transactionId) // Make sure this exists, else a 404 error will occur
    if err != nil {
        panic(err)
    }
    fmt.Println(*txn)

    // Fetch all parties on a transaction
    parties, err := glideClient.Transactions().Parties().List(txn.Id)
    if err != nil {
        panic(err)
    }
    fmt.Println(*parties)

    // Use page parameters to control List's methods pagination. Pagination has a limit and a cursor.
    // No custom sorting is yet supported. WithPage(nil) is equivalent to not including the option at all.
    // WithPageParams(5, "1486") is an equivalent way of writing this example
    txns, err = glideClient.Transactions().List(glide.WithPage(&glide.PageParams{Limit: 5, StartingAfter: "1486"}))
    if err != nil {
        panic(err)
    }
    fmt.Println(*txns)

    // Simplified approach to only set the Limit without setting cursor
    txns, err = glideClient.Transactions().List(glide.WithPageLimit(5))
    if err != nil {
        panic(err)
    }
    fmt.Println(*txns)

    // Simplified approach to only set the Starting after cursor while using default limit value
    txns, err = glideClient.Transactions().List(glide.WithPageStartingAfter("1486"))
    if err != nil {
        panic(err)
    }
    fmt.Println(*txns)

    // All list objects have `HasMore bool` and `NextPageParams() *PageParams` to figure out next page's attributes
    // If HasMore is false, then NextPageParams() will be nil. Otherwise, NextPageParams() will return the appropriate
    // PageParams data to request the next page.
    txns, err = glideClient.Transactions().List()
    if err != nil {
        panic(err)
    }
    for {
        fmt.Println("Fetched", len(txns.Data), "txns")
        if !txns.HasMore {
            break
        }
        txns, err = glideClient.Transactions().List(WithPage(txns.NextPageParams()))
        if err != nil {
            panic(err)
        }
    }

    // Most nested objects won't be included in the responses by default, but they can be expanded if the
    // option is included.
    // All response objects have an `IsRef() bool` to check if the content corresponds to the actual object
    // or if it's just a reference to it (i.e. it could have been expanded, but it wasn't).
    txn, err = glideClient.Transactions().GetDetail(transactionId) // Make sure this exists, else a 404 error will occur
    if err != nil {
        panic(err)
    }
    if !txn.Folders.IsRef() || !txn.Parties.IsRef() {
        panic("These should be refs!")
    }

    txn, err = glideClient.Transactions().GetDetail(transactionId,
        glide.WithExpand(
            "folders.transaction_documents",
            "parties",
        ),
    ) // Make sure this exists, else a 404 error will occur
    if err != nil {
        panic(err)
    }
    fmt.Println(txn)
    if txn.Folders.IsRef() || txn.Parties.IsRef() {
        panic("These should NOT be refs!")
    }
    for _, folder := range txn.Folders.Data {
        if folder.TransactionDocuments.IsRef() {
            panic("These should NOT be refs!")
        }
        fmt.Println(folder.Title)
        for _, td := range folder.TransactionDocuments.Data {
            fmt.Println(td)
        }
    }

    // Expansion also works for List and GetMulti methods
    _, err = glideClient.Transactions().Folders().List(txn.Id,
        glide.WithExpand(
            "transaction_documents",
        ),
    )
    if err != nil {
        panic(err)
    }

    folderIds := []string{}
    for _, f := range txn.Folders.Data {
        folderIds = append(folderIds, f.Id)
    }
    _, err = glideClient.Transactions().Folders().GetMulti(txn.Id, folderIds,
        glide.WithExpand(
            "transaction_documents",
        ),
    )
    if err != nil {
        panic(err)
    }

    // Most write operations are performed at the Transaction level. Some examples follow
    // Some operations will include some details about the performed action on its Response object,
    // while others will now
    partyCreateRes, err := glideClient.Transactions().PartyCreates(transactionId, glide.PartyCreates{
        Creates: []glide.PartyCreate{{
            Email:     "some@test.com",
            FirstName: "Test",
            LastName:  "Test",
            Roles:     []string{"LISTING_TC"},
        }},
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(partyCreateRes)
    txn, err = glideClient.Transactions().GetDetail(transactionId, glide.WithExpand("parties"))
    if err != nil {
        panic(err)
    }
    var party glide.Party
    for _, p := range txn.Parties.Data {
        if p.Email == "some@test.com" {
            party = p
            break
        }
    }
    partyRemovesRes, err := glideClient.Transactions().PartyRemoves(transactionId, glide.PartyRemoves{
        Removes: []glide.PartyRemove{{
            PartyId: party.Id,
        }},
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(partyRemovesRes)

    // Transaction Fields is currently the only object not included by default (it ahs to be explictly expanded) that
    // does not correspond to a related entity.
    // To include fields when fetching a transaction, the "fields" attribute has to be exapnded. Expanind "fields"
    // only or "fields[*]" will include all fields. To get a subset of fields, the required fileds can be included
    // like this "fields[property/apn,property/address,parties/seller,fieldx,fieldy...]"

    txn, err = glideClient.Transactions().GetDetail(transactionId)
    if err != nil {
        panic(err)
    }
    fmt.Println(txn.Fields) // Empty

    txn, err = glideClient.Transactions().GetDetail(transactionId,
        glide.WithExpand("fields"), // equivalent to WithExpand("fields[*]")
    )
    if err != nil {
        panic(err)
    }
    fmt.Println(txn.Fields) // All fields included

    txn, err = glideClient.Transactions().GetDetail(transactionId,
        glide.WithExpand("fields[property/apn,property/block,property/lot]"),
    )
    if err != nil {
        panic(err)
    }
    fmt.Println(txn.Fields) // Only property/apn,property/block,property/lot included

    // Writing to fields requires a TransactionFieldsWrite object
    // There, each key corresponds to a field_id, and each value is the field value
    // to write, and optionally the control_version. The controlVersion (if included)
    // will make Glide check if that value corresponds to the version # in which that field
    // was last updated, and will fail in case it's not. If not provided (or provided with
    // value 0), then it will always write. In the same TransactionFieldsWrite there could be
    // a combination of some fields having controlVersion and others not having it
    fieldsWrite := glide.TransactionFieldsWrite{
        "property/apn":   glide.GetFieldWrite("12345", 2247),    // Will fail if current version of property/apn != 2247
        "property/block": glide.GetFieldWriteNoControl("67890"), // Will always write to property/block
        "property/lot":   glide.GetFieldWrite("1233", 0),        // This is equivalent to using GetFieldWriteNoControl
    }
    fieldsRes, err := glideClient.Transactions().Fields(txn.Id, fieldsWrite)
    if err != nil {
        // err's params will detail each failing field in case some of the controled ones found a version mismatch
        fmt.Println("err.Params", err.Params)
    } else {
        // Response will detail all the values and latest versions of alterted fields
        fmt.Println("fieldsRes", fieldsRes)
    }

    // TransactionFieldsWrite can be created from txn objects, so you can get the controlVersion values from there
    txn, err = glideClient.Transactions().GetDetail(transactionId, glide.WithExpand("fields"))
    if err != nil {
        panic(err)
    }
    // This TransactionFieldsWrite will include the controlVersion that match whatever there is stored
    // on the specified fields for txn.Fields. IF the field is not present, no cotnrol will be included for it
    fieldsWrite = txn.GetFieldsWrite(glide.TransactionFieldValues{
        "property/apn":   "12345",
        "property/block": "910233",
    })
    fmt.Println("fieldsWrite", fieldsWrite)

    // Multiple TransactionFieldsWrite objects can be combined into one using CombineFieldsWrites
    fieldsWrite = CombineFieldsWrites(
        txn.GetFieldsWrite(glide.TransactionFieldValues{
            "property/apn":   "12345",
            "property/block": "910233",
        }),
        // If the same field is included in more than one of the combined TransactionFieldsWrite,
        // the last provided value will be kept
        glide.TransactionFieldsWrite{
            "property/block": GetFieldWriteNoControl("67890"),
            "property/lot":   GetFieldWrite("1233", 0),
        },
    )
    fmt.Println("fieldsWrite", fieldsWrite)
}
```

