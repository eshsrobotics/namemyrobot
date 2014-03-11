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

	m.Get(`/api/adjectives`, routes.GetAdjectives)
	m.Post(`/api/adjectives`, binding.Bind(models.Adjective{}),
		routes.AddAdjective)
	m.Put(`/api/adjectives/:id/vote`, routes.VoteAdjective)

	m.Get(`/api/names`, routes.GetNames)
	m.Post(`/api/names`, binding.Bind(models.Name{}), routes.AddName)
	m.Put(`/api/names/:id/vote`, routes.VoteName)

	m.Run()
}
