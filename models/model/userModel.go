package model

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Age       int16     `json:"age"`
	JobTitle  string    `json:"job_title"`
	Company   string    `json:"company"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
