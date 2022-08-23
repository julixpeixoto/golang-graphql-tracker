package data

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	TrackingStatus TrackingStatus `json:"trackingStatus"`
}

type TrackingStatus struct {
	Tracking Tracking `json:"tracking"`
}

type Tracking struct {
	History []History `json:"history`
}

type History struct {
	EventDate   string `json:"event_date"`
	StatusLabel string `json:"status_label"`
}

func convertResponseData(response []byte) Response {
	data := Response{}
	json.Unmarshal(response, &data)
	return data
}

func FormatToBody(responseData []byte) string {
	newData := convertResponseData(responseData)
	var body string

	for _, h := range newData.Data.TrackingStatus.Tracking.History {
		body = body + fmt.Sprintf("%s %s\n", h.EventDate, h.StatusLabel)
	}

	res, _ := regexp.Compile(`[^\w]`)
	return res.ReplaceAllString(body, "")
}
