package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/coopernurse/gorp"
	"github.com/eshsrobotics/namemyrobot/models"
)

func GetLastNames(db gorp.SqlExecutor) (int, string) {
	var lastNames []models.FirstName
	_, err := db.Select(&lastNames, "select * from last_names")
	if err != nil {
		log.Println(err, "Error selecting last names from database")
		return http.StatusInternalServerError, ""
	}

	json, err := json.Marshal(lastNames)
	if err != nil {
		log.Println(err, "Error marshaling last names JSON")
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, string(json)
}

func AddLastName(lastName models.LastName, w http.ResponseWriter,
	db gorp.SqlExecutor) (int, string) {

	lastName.CreatedAt = time.Now()

	if err := db.Insert(&lastName); err != nil {
		log.Println(err, "Insert failed")
		return http.StatusConflict, ""
	}

	w.Header().Set("Location", fmt.Sprintf("/api/lastNames/%d", lastName.Id))

	json, err := json.Marshal(lastName)
	if err != nil {
		log.Println(err, "Error marshaling inserted last name")
		return http.StatusInternalServerError, ""
	}

	return http.StatusCreated, string(json)
}
