package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/abdukahhor/streamer/app"
	"github.com/abdukahhor/streamer/handlers/rpc"
	"github.com/abdukahhor/streamer/models"
	"github.com/abdukahhor/streamer/storages/rdb"
)

func main() {
	var cfg models.Config
	err := cfg.Get("./config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	cfg.Ln = len(cfg.URLs)
	cache, err := rdb.Connect(":6363", 20)
	if err != nil {
		log.Fatalln(err)
	}
	c := app.New(cache, cfg)
	rpc.Register(c, cfg.NumberOfRequests)
	go func() {
		sigint := make(chan os.Signal)
		signal.Notify(sigint, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
		log.Println("server received signal ", <-sigint)

		rpc.Close()
		os.Exit(0)
	}()

	log.Println("server started at ", ":9090")
	//serve rpc
	err = rpc.Run(":9090")
	if err != nil {
		log.Fatalln(err)
	}
}
