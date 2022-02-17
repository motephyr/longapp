package main

import (
	"flag"
	"log"

	"github.com/motephyr/longcare/app"
	"github.com/motephyr/longcare/migrations"
	"github.com/motephyr/longcare/models"
	"github.com/motephyr/longcare/pkg/modelhelper"
	"github.com/motephyr/longcare/rest/routes"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	configFile := flag.String("config", "config.yml", "User Config file from user")
	migrate := flag.Bool("migrate", false, "Update db structure")
	flag.Parse()
	app.Load(*configFile)
	if app.Http.Profiler.Enabled {
		_, _ = profiler.Start(profiler.Config{
			ApplicationName: app.Http.Server.Name,
			ServerAddress:   app.Http.Profiler.Server,
		})
	}

	app.Http.Server.Version = app.Version
	if *migrate {
		migrations.Migrate()
	} else {
		models.AddOlderHook(boil.AfterUpdateHook, modelhelper.Older.AfterUpdateHook)
		models.AddUserOlderHook(boil.AfterDeleteHook, modelhelper.UserOlder.AfterDeleteHook)
		models.AddUserOlderHook(boil.AfterInsertHook, modelhelper.UserOlder.AfterInsertHook)
		routes.LoadRoutes(app.Http.Server.App)

		app.Http.Route404()
		log.Fatal(app.Http.Server.ServeWithGraceFullShutdown())
	}

}
