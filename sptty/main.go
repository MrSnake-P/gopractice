package main

import (
	"gin/sptty/src/user"
	"github.com/linshenqi/sptty"
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