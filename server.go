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

	m.Get(`/namemyrobot/first_names`, routes.GetFirstNames)
	m.Get(`/namemyrobot/middle_names`, routes.GetLastNames)
	m.Get(`/namemyrobot/last_names`, routes.GetLastNames)

	m.Run()
}
