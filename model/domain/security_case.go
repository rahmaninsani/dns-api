package domain

import "time"

type SecurityCase struct {
	ID        string    `json:"id"`
	Subject   string    `json:"subject"`
	Risk      string    `json:"risk"`
	Analyst   string    `json:"analyst"`
	Timestamp time.Time `json:"timestamp"`
}
