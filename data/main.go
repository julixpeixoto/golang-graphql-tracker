package data

import (
	"encoding/json"
	"fmt"
	"tracker/model"
)

func ConvertResponseData(response []byte) model.Response {
	data := model.Response{}
	json.Unmarshal(response, &data)
	return data
}

func FormatToBody(newData model.Response) string {
	var body string

	for _, h := range newData.Data.TrackingStatus.Tracking.History {
		body = body + fmt.Sprintf("%s %v\n", h.EventDate, h.StatusLabel)
	}

	return body
}
