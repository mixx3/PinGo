package repo

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	StatusCode     int
	ResponseTimeMs int
	Address        string
	Name           string
	ReceiverID     int
	Receiver       Receiver
}

type Request struct {
	gorm.Model
	StatusExpected         []int
	Body                   string
	ExpectedResponseTimeMs int
	ReceiverID             int
	Receiver               Receiver
}

type Receiver struct {
	gorm.Model
	SocialID string
	Name     string
}
