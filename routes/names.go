package routes

import (
	"encoding/json"
	"log"
	"net/http"

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
