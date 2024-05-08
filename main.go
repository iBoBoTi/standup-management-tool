package main

import (
	"expvar"
	"io"
	"log"
	"os"

	"github.com/iBoBoTi/standup-management-tool/internal/config"
	logger "github.com/iBoBoTi/standup-management-tool/internal/log"
	"github.com/iBoBoTi/standup-management-tool/internal/models"
	"github.com/iBoBoTi/standup-management-tool/routers"
	"github.com/iBoBoTi/standup-management-tool/server"
)

// build time variables
var (
	buildTime string
	version   string
)

func main() {

	expvar.NewString("version").Set(version)

	cfg, err := config.Load(".")
	if err != nil {
		log.Println("unable to read env")
		log.Fatal(err)
	}

	var logWriter io.Writer = os.Stdout

	db := models.GetDB(cfg)

	zeroLogger := logger.NewZeroLogger(logWriter, logger.LevelInfo)
	srv, err := server.NewServer(cfg, db, zeroLogger)
	if err != nil {
		log.Fatal(err)
	}

	srv.BuildTime = buildTime
	srv.Version = version

	routers.SetupRouter(srv)
	server.RunGinServer(srv)

}
