package model

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
