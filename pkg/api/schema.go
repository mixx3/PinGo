package api

import "time"

type LogPostSchema struct {
	StatusCode     int    `json:"status_code"`
	ResponseTimeMs int    `json:"response_time_ms"`
	Address        string `json:"address" example:"https://github.com/mixx3"`
	Name           string `json:"name"`
	ReceiverID     int    `json:"receiver_id"`
}

type ReceiverPostSchema struct {
	SocialID string `json:"social_id"`
	Name     string `json:"name"`
}

type RequestPostSchema struct {
	StatusExpected         []int  `json:"status_expected"`
	Body                   string `json:"body"`
	ExpectedResponseTimeMs int    `json:"expected_response_time_ms"`
	ReceiverID             int    `json:"receiver_id"`
}

type LogGetSchema struct {
	ID             int       `json:"id"`
	StatusCode     int       `json:"status_code"`
	ResponseTimeMs int       `json:"response_time_ms"`
	Address        string    `json:"address" example:"https://github.com/mixx3"`
	Name           string    `json:"name"`
	ReceiverID     int       `json:"receiver_id"`
	DtCreated      time.Time `json:"dt_created"`
}

type ReceiverGetSchema struct {
	ID       int    `json:"id"`
	SocialID string `json:"social_id"`
	Name     string `json:"name"`
}

type RequestGetSchema struct {
	ID                     int    `json:"id"`
	StatusExpected         []int  `json:"status_expected"`
	Body                   string `json:"body"`
	ExpectedResponseTimeMs int    `json:"expected_response_time_ms"`
	ReceiverID             int    `json:"receiver_id"`
}
