package data

import (
	"encoding/json"
	"fmt"
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

func ConvertResponseData(response []byte) Response {
	data := Response{}
	json.Unmarshal(response, &data)
	return data
}

func FormatToBody(newData Response) string {
	var body string

	for _, h := range newData.Data.TrackingStatus.Tracking.History {
		body = body + fmt.Sprintf("%s %v\n", h.EventDate, h.StatusLabel)
	}

	return body
}
