package main

import (
	"embed"
	"flag"
	"log"

	"github.com/motephyr/longapp/app"
	config "github.com/motephyr/longapp/config"
	"github.com/motephyr/longapp/migrations"
	"github.com/motephyr/longapp/models"
	"github.com/motephyr/longapp/pkg/modelhelper"
	"github.com/motephyr/longapp/rest/routes"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/webview/webview"
)

// Embed a directory
//go:embed resources/views/*
var embedDirStaticView embed.FS

func main() {
	go func() {
		config.EmbedDirStaticView = embedDirStaticView

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
	}()
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:34115")
	w.Run()
}
