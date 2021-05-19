package models

import "time"

type Odd struct {
	Id          string
	Name        string
	Match       string
	Coefficient float64
	Timestamp   time.Time
}
