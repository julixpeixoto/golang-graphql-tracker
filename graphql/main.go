package graphql

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func GetData() {
	query := getQuery()
	print(query)
	request, err := http.NewRequest("POST", viper.GetString("API_URL"), bytes.NewBuffer([]byte(query)))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		println("Error than mount request:", err)
	}

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)

	if err != nil {
		println("Error than request:", err)
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	print(string(data))
}

func getQuery() string {
	return fmt.Sprintf(`{"query":" {\r\n      trackingStatus(clientId: \"%s\", orderNumber: \"%s\", orderHash:\"%s\") {\r\n        order {\r\n          order_number\r\n        }\r\n        tracking {\r\n          status\r\n          status_label\r\n          estimated_delivery_date\r\n          estimated_delivery_date_lp\r\n          history {\r\n            event_date\r\n            status_label\r\n            is_warning\r\n          }\r\n        }\r\n        sender {\r\n          address {\r\n            city\r\n            zip_code\r\n            state_code\r\n          }\r\n        }\r\n        end_customer {\r\n          address {\r\n            city\r\n            zip_code\r\n            state\r\n          }\r\n        }\r\n      }\r\n    }\r\n\r\n\r\n","variables":{}}`, viper.Get("CLIENT_ID"), viper.Get("ORDER_NUMBER"), viper.Get("ORDER_HASH"))
}
