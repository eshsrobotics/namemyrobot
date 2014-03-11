package main

import (
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/eshsrobotics/namemyrobot/models"
	"github.com/eshsrobotics/namemyrobot/routes"
	"github.com/martini-contrib/binding"
)

func main() {
	m := martini.Classic()

	m.Use(martini.Static("public"))
	m.MapTo(models.Dbm, (*gorp.SqlExecutor)(nil))

	m.Get(`/api/firstNames`, routes.GetFirstNames)
	m.Post(`/api/firstNames`, binding.Bind(models.FirstName{}),
		routes.AddFirstName)

	m.Get(`/api/middleNames`, routes.GetLastNames)
	m.Post(`/api/middleNames`, binding.Bind(models.MiddleName{}),
		routes.AddMiddleName)

	m.Get(`/api/lastNames`, routes.GetLastNames)
	m.Post(`/api/lastNames`, binding.Bind(models.LastName{}),
		routes.AddLastName)

	m.Run()
}
