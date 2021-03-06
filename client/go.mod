module github.com/divilla/eop09/client

go 1.16

require (
	github.com/divilla/eop09/entityproto v0.0.0
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/labstack/echo/v4 v4.5.0
	github.com/labstack/gommon v0.3.0
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shopspring/decimal v1.2.0
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/tidwall/gjson v1.8.1
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/sjson v1.1.7 // indirect
	github.com/valyala/fastjson v1.6.3 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
	golang.org/x/sys v0.0.0-20210820121016-41cdb8703e55 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
)

replace github.com/divilla/eop09/entityproto => ./../entityproto
