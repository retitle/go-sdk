module github.com/retitle/go-sdk/v3

go 1.18

replace (
	github.com/retitle/go-sdk/v3/core => ./core/
	github.com/retitle/go-sdk/v3/core/mocks => ./core/mocks/
	github.com/retitle/go-sdk/v3/tests_utils => ./tests_utils/
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
