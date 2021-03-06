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

func GetNames(db gorp.SqlExecutor) (int, string) {
	var names []models.Name
	_, err := db.Select(&names, "select * from name order by id")
	if err != nil {
		log.Println(err, "Error selecting names from database")
		return http.StatusInternalServerError, ""
	}

	json, err := json.Marshal(names)
	if err != nil {
		log.Println(err, "Error marshaling names JSON")
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, string(json)
}

func AddName(name models.Name, w http.ResponseWriter,
	db gorp.SqlExecutor) (int, string) {

	name.CreatedAt = time.Now()

	if err := db.Insert(&name); err != nil {
		log.Println(err, "Insert failed")
		return http.StatusConflict, ""
	}

	w.Header().Set("Location", fmt.Sprintf("/api/names/%d", name.Id))

	json, err := json.Marshal(name)
	if err != nil {
		log.Println(err, "Error marshaling inserted last name")
		return http.StatusInternalServerError, ""
	}

	return http.StatusCreated, string(json)
}

func VoteName(db gorp.SqlExecutor, params martini.Params) (int, string) {
	obj, err := db.Get(models.Name{}, params["id"])
	if err != nil {
		log.Println(err, "Name not found")
		return http.StatusNotFound, ""
	}

	name := obj.(*models.Name)

	name.Votes++

	_, err = db.Update(name)
	if err != nil {
		log.Println(err, "Error updating name")
		return http.StatusInternalServerError, ""
	}

	json, err := json.Marshal(name)
	if err != nil {
		log.Println(err, "Error marshalling name to JSON")
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, string(json)
}
