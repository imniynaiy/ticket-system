package main

import (
	"github.com/imniynaiy/ticket-system/internal/config"
	"github.com/imniynaiy/ticket-system/internal/database"
	"github.com/imniynaiy/ticket-system/internal/flag"
	"github.com/imniynaiy/ticket-system/internal/log"
	"github.com/imniynaiy/ticket-system/internal/server"
	"github.com/imniynaiy/ticket-system/internal/verflag"
)

func main() {
	flag.ParseAndHandleHelpFlag()
	verflag.PrintAndExitIfRequested()
	config.ParseConfig()
	log.Init(&config.GlobalConfig.Log)
	defer log.Sync()

	log.Info("config:", log.String("file", *config.ConfigFile))
	log.Info("version:", log.String("build_date", verflag.BuildDate),
		log.String("git_version", verflag.GitVersion), log.String("git_commit", verflag.GitCommit),
		log.String("git_tree_state", verflag.GitTreeState),
	)

	database.InitDB(&config.GlobalConfig.Database)
	database.InitRedis(&config.GlobalConfig.Redis)

	server.Start()
}
