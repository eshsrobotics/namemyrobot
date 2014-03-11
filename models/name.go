package models

import "time"

type Name struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name" binding:"required"`
}

type FirstName Name
type MiddleName Name
type LastName Name
