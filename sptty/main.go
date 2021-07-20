package main

import (
	"github.com/linshenqi/sptty"
	"sptty/src/user"
)

func main() {

	app := sptty.GetApp()
	app.LoadConfFromFile()

	app.AddServices(sptty.Services{
         &user.Service{},
	})

	app.AddConfigs(sptty.Configs{

	})

	app.Sptting()
}