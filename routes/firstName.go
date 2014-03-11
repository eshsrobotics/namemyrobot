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

func GetFirstNames(db gorp.SqlExecutor) (int, string) {
	var firstNames []models.FirstName
	_, err := db.Select(&firstNames, "select * from first_names")
	if err != nil {
		log.Println(err, "Error selecting first names from database")
		return http.StatusInternalServerError, ""
	}

	json, err := json.Marshal(firstNames)
	if err != nil {
		log.Println(err, "Error marshaling first names JSON")
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, string(json)
}

func AddFirstName(firstName models.FirstName, w http.ResponseWriter,
	db gorp.SqlExecutor) (int, string) {

	firstName.CreatedAt = time.Now()

	if err := db.Insert(&firstName); err != nil {
		log.Println(err, "Insert failed")
		return http.StatusConflict, ""
	}

	w.Header().Set("Location", fmt.Sprintf("/api/firstNames/%d", firstName.Id))

	json, err := json.Marshal(firstName)
	if err != nil {
		log.Println(err, "Error marshaling inserted first name")
		return http.StatusInternalServerError, ""
	}

	return http.StatusCreated, string(json)
}
