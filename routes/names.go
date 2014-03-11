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
