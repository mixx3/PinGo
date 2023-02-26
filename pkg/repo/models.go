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
	Receiver       Receiver `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Request struct {
	gorm.Model
	StatusExpected         int
	Name                   string
	Address                string
	Body                   string
	ExpectedResponseTimeMs int
	RepeatTimeMs           int
	ReceiverID             int
	Receiver               Receiver `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Receiver struct {
	gorm.Model
	SocialID string
	Name     string
}
