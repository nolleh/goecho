module gobank

go 1.12

replace gobank => ./

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.6
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.2.9 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.4.0
)