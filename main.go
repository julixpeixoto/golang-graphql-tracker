package main

import (
	"tracker/data"
	"tracker/database"
	"tracker/email"
	"tracker/graphql"
	"tracker/model"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	database.CreateDatabaseIfNotExists()
	responseInBytes := graphql.GetData()
	responseData := data.ConvertResponseData(responseInBytes)

	if checkNewEvent(responseData) {
		println("Nova movimentação localizada!")
		body := data.FormatToBody(responseData)
		email.SendEmail(string(body))
	} else {
		println("Nenhuma nova movimentação localizada!")
	}
}

func checkNewEvent(response model.Response) bool {
	events := len(response.Data.TrackingStatus.Tracking.History)
	countDatabase := database.GetCount()

	if int32(events) > countDatabase {
		database.SaveCount(int32(events))
		return true
	}

	return false
}
