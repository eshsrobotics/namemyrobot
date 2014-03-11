package main

import (
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/eshsrobotics/namemyrobot/models"
	"github.com/eshsrobotics/namemyrobot/routes"
)

func main() {
	m := martini.Classic()

	m.Use(martini.Static("public"))
	m.MapTo(models.Dbm, (*gorp.SqlExecutor)(nil))

	m.Get(`/api/firstNames`, routes.GetFirstNames)
	m.Get(`/api/middleNames`, routes.GetLastNames)
	m.Get(`/api/lastNames`, routes.GetLastNames)

	m.Run()
}
