package main

import (
	"webscraping/data"
	"webscraping/email"
	"webscraping/graphql"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	responseData := graphql.GetData()
	body := data.FormatToBody(responseData)
	email.SendEmail(string(body))
}
