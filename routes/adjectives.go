package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/eshsrobotics/namemyrobot/models"
)

func GetAdjectives(db gorp.SqlExecutor) (int, string) {
	var adjectives []models.Adjective
	_, err := db.Select(&adjectives, "select * from adjective")
	if err != nil {
		log.Println(err, "Error selecting adjectives from database")
		return http.StatusInternalServerError, ""
	}

	json, err := json.Marshal(adjectives)
	if err != nil {
		log.Println(err, "Error marshaling adjectives JSON")
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, string(json)
}

func AddAdjective(adjective models.Adjective, w http.ResponseWriter,
	db gorp.SqlExecutor) (int, string) {

	adjective.CreatedAt = time.Now()

	if err := db.Insert(&adjective); err != nil {
		log.Println(err, "Insert failed")
		return http.StatusConflict, ""
	}

	w.Header().Set("Location", fmt.Sprintf("/api/adjectives/%d", adjective.Id))

	json, err := json.Marshal(adjective)
	if err != nil {
		log.Println(err, "Error marshaling inserted first name")
		return http.StatusInternalServerError, ""
	}

	return http.StatusCreated, string(json)
}

func VoteAdjective(db gorp.SqlExecutor, params martini.Params) (int, string) {
	obj, err := db.Get(models.Adjective{}, params["id"])
	if err != nil {
		log.Println(err, "Adjective not found")
		return http.StatusNotFound, ""
	}

	adjective := obj.(*models.Adjective)

	adjective.Votes++

	_, err = db.Update(adjective)
	if err != nil {
		log.Println(err, "Error updating adjective")
		return http.StatusInternalServerError, ""
	}

	json, err := json.Marshal(adjective)
	if err != nil {
		log.Println(err, "Error marshalling adjective to JSON")
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, string(json)
}
