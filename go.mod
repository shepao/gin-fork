module github.com/gin-gonic/gin

go 1.12

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7
	github.com/golang/protobuf v0.0.0-20170601230230-5a0f697c9ed9
	github.com/json-iterator/go v0.0.0-20170829155851-36b14963da70
	github.com/mattn/go-isatty v0.0.0-20170307163044-57fdcb988a5c
	github.com/shepao/valid v1.0.0
	github.com/stretchr/testify v0.0.0-20160925220609-976c720a22c8
	github.com/ugorji/go v0.0.0-20170215201144-c88ee250d022
	golang.org/x/net v0.0.0-20161018194804-d4c55e66d8c3
	golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.1
	gopkg.in/yaml.v2 v2.0.0-20160928153709-a5b47d31c556
)

replace golang.org/x/net => github.com/golang/net v0.0.0-20161018194804-d4c55e66d8c3

replace github.com/shepao/valid => github.com/shepao/validator v1.0.0
