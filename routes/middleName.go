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

func GetMiddleNames(db gorp.SqlExecutor) (int, string) {
	var middleNames []models.FirstName
	_, err := db.Select(&middleNames, "select * from middle_names")
	if err != nil {
		log.Println(err, "Error selecting middle names from database")
		return http.StatusInternalServerError, ""
	}

	json, err := json.Marshal(middleNames)
	if err != nil {
		log.Println(err, "Error marshaling middle names JSON")
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, string(json)
}

func AddMiddleName(middleName models.MiddleName, w http.ResponseWriter,
	db gorp.SqlExecutor) (int, string) {

	middleName.CreatedAt = time.Now()

	if err := db.Insert(&middleName); err != nil {
		log.Println(err, "Insert failed")
		return http.StatusConflict, ""
	}

	w.Header().Set("Location", fmt.Sprintf("/api/middleNames/%d", middleName.Id))

	json, err := json.Marshal(middleName)
	if err != nil {
		log.Println(err, "Error marshaling inserted middle name")
		return http.StatusInternalServerError, ""
	}

	return http.StatusCreated, string(json)
}
