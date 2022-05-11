package glide

import (
  "fmt"
  "testing"

  // "github.com/dgrijalva/jwt-go"
)

const (
  JwtExpirationInMinutes = 60
  TransactionScope       = "TRANSACTIONS"
  ContactScope           = "CONTACTS"
  clientKey              = "1c21de69-64b7-4f2e-8695-31681c60cd26"
  glideHostName          = "http://api.localhost"
  userEmail              = "joel3@joel.com"
  privatekey             = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDbw5MCC+pohI0q
B3JSc3PFMRpSf1h6TAFWinPreiU8ruwSXBhjA6vwmJqv8IWvkFED0u16/KGOWzzZ
uJqsnMgQvInbc3PllVZLJH0cjvhM8pLNiGPRmksk1QLZSP1zKBh/wKZFQaPTRNTg
p0HVZomvXGt0lo0Ou/JjQb4dS733TbjY8pTupFdCsRvfblvb+Xt6bMugA6trY3dv
59REiIYhL3At+OpO1rvB8rR4liIufRvD5G7tFJv9r85LtwdTdmK4gEljwwAicG1h
3gK5vudjzsbBpMfpG//IaimctHMKoHTxRfWli2+r4KrJ9E3gpwFfPA94N/gw1etr
ePxV/aC1AgMBAAECggEBAIg6HrnJCyCI8jyYaxyDJsOY4Zk8CWueC8JKhvr0N3r7
b6kd35wZHY9B+bmQXj5rNl0pdj6jtb3Z/slzrqXLdhUx2j7nvXMWdGyWDNwUIAUK
5Ud1AXNfsq79QAvTPBETTMR8dlU+EECZTWCJLb0MF9NVdeLKpFv605EFosTkO7i9
Ff1HYNSLjHjfE7X3aefZ6RHA3x/qwhZMOsHqjM+WAyI8BAqeRrwqO68YCWitPjT+
E7e6ienSgLe7AHzW975EEavhGcttlUyUgtYlQtGHr4jFgTgUj7vGlFHw3+MePgwF
pgKR9NXnP6Hjr5lgPOrjrhw1xm6yO/YCduGZujHZzI0CgYEA8KvwTsdLQ0ueFVAa
5I7OhT2+SM9jqU7NB9FoDA2Um9mvOb0S5Sp+tuGC1bDFAijzgZ8wt8UoOLwB6RLS
Kr64v65i47orO1Dufg3azdlHQqN2SI4T2NB7vUUr/lRJHma6vLza9QlTX+Pnf71y
rc/+gbgTa+zHW7+pGsnyZPjd+IcCgYEA6cK+ZKdUbwNvOd9+UTnSwOCy0NKxUyll
jF9Ef7oqH2zHoFJfMkoqdbqZF3vMHbhXhPI7O3eSfI/jslHFBCwPkEYbXr1z8kBf
KQ1sDUEKo57YsQhBzoV7rcSTcumbAcxsqRYDZW+L/qiz/J5k8FVQ5QhmBKEoNwLQ
Q9Q3wQag9+MCgYA/Z+0NtC+98QQa9VnAcWczb0rcf9bv2hTmRGM3GbTXQoiJm6iI
u6NapxsDFWkx4nwU4E3DfKvWFqIiN5UrMcgWp+jhukB9hhrvFtNYfC/r/IjDILtV
2cdf3AN1I9uHqOT1qIO/Hs/aaX6qBs1ZwXx5zHdBbcYwA9SXfDDiLTa9rQKBgHaa
3MdtLyX0dCbFAu5rjEdYuuHBRT+QrXl/jN3RszWml9L8eFin0MtTFgIYSgR6V82Y
qf0OOkEBMOJ1IqVvRaZKK/Dx5zZu+tl1efFvotpJ6mBIdDs49vu/1aBkbsWG10sj
ZC2/XySirE/sfgKDBzxt+nU235Sp1MnVj0Rc/KdJAoGBAObglEa/AJdrB8O912z6
9AVeBbdz7kDzYkE1Yplvf0jxeI+aw2ve7fVanbJjuViVvkpFg1HSHdxrKwC4IjwJ
jh/Qffk+RdbUHkOyOcVQnzkkWFApzlktylRC3bVmxUHusfE2PB4wRXi4iqWf0cfn
s2zwZX+YXr64UQmYnkOnsEzk
-----END PRIVATE KEY-----`
)

func GetTransactionsClient() (TransactionsResource, error) {
  key := GetRsa256KeyFromPEMString(privatekey)
  glideClient := GetClient(clientKey, key, WithProtocol("http"), WithHost("api.localhost"))
  if err := glideClient.StartImpersonating(userEmail, []string{"TRANSACTIONS"}); err != nil {
    fmt.Println(err)
    return nil, err
  }
  return glideClient.Transactions(), nil
}

func TestCreateTransaction(t *testing.T) {
  transactionsClient, err := GetTransactionsClient()
  if err != nil {
    fmt.Println(err)
  }

  txn, err := transactionsClient.GetDetail("1075")

  if err != nil {
    panic(err)
  }
  wantId := "1075"
  if gotId := txn.Id; gotId != wantId {
    t.Errorf("GetClient().options.audience = %q, want %q", gotId, wantId)
  }
}

func TestCreateTransactionWithNewParty(t *testing.T) {
  transactionsClient, err := GetTransactionsClient()
  if err != nil {
    fmt.Println(err)
  }

  txn, err := transactionsClient.GetDetail("1075")
  roles := []string{"BUYER"}
  contactRequest := ContactRequest{}
  contactSource := ContactSourceRequest{}
  contactSource.Origin = "COMPASS_PERSON"
	contactRequest.Email = "new@email.com"
	contactRequest.FirstName = "New"
	contactRequest.LastName = "Email"

	partyCreate := PartyCreate{
		Contact:             contactRequest,
		Roles:               roles,
		Invite:              true,
		SuppressInviteEmail: true,
    UserContactSource: contactSource,
    UserContactId: "1",
	}
	partyCreates := PartyCreates{
		Creates: []PartyCreate{partyCreate},
	}

  fmt.Println(partyCreates)

  if err != nil {
    panic(err)
  }
  wantId := "1075"
  if gotId := txn.Id; gotId != wantId {
    t.Errorf("GetClient().options.audience = %q, want %q", gotId, wantId)
  }
}