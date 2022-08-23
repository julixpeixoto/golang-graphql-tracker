package main

import (
	"webscraping/email"
	"webscraping/graphql"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	email.SendEmail("body here")
	graphql.GetData()
}
