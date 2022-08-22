package main

import (
	"webscraping/graphql"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	graphql.GetData()
}
