package models

import (
	"net/http"
	"time"

	"github.com/martini-contrib/binding"
)

type Name struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content" binding:"required"`
	Votes     int       `json:"votes"`
}

func (n Name) Validate(errors *binding.Errors, req *http.Request) {
	if n.Votes != 0 {
		errors.Fields["votes"] = "Votes must not be anything but 0"
	}
}
