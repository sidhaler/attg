module github.com/sidhaler/attg

go 1.13

require (
	github.com/c-bata/go-prompt v0.2.6
	github.com/go-sql-driver/mysql v1.6.0
	github.com/spf13/viper v1.8.1
)

replace github.com/sidhaler/attg/Util => ./Util

replace github.com/sidhaler/attg/dbUtil => ./dbUtil

replace github.com/sidhaler/attg/attconf => ./attconf
