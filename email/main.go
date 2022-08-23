package email

import (
	"fmt"
	"net/smtp"

	"github.com/spf13/viper"
)

func SendEmail(body string) {
	user := viper.GetString("EMAIL_USER")
	password := viper.GetString("EMAIL_PASSWORD")
	host := viper.GetString("EMAIL_HOST")
	port := viper.GetString("EMAIL_PORT")
	address := host + ":" + port
	emailAddress := viper.GetString("EMAIL_ADDRESS")
	orderNumber := viper.GetString("ORDER_NUMBER")
	to := []string{emailAddress}
	subject := fmt.Sprintf("Subject: Status do pedido %s atualizado\n", orderNumber)
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", user, password, host)
	err := smtp.SendMail(address, auth, user, to, message)

	if err != nil {
		println("Erro ao enviar o e-mail: ", err)
	}

	println("E-mail enviado com sucesso!")
}
