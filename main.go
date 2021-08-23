package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zeroidentidad/gopherbot/config"
	"github.com/zeroidentidad/gopherbot/messages"
	"github.com/zeroidentidad/gopherbot/status"
	"github.com/zeroidentidad/gopherbot/webservice"
	"github.com/zeroidentidad/gopherbot/webservice/storage"
)

func main() {
	dg, err := discordgo.New("Bot " + config.TOKEN)
	if err != nil {
		log.Fatal("Session error:", err.Error())
		return
	}
	dg.Debug = true

	// Register bot handlers.
	dg.AddHandler(messages.MessageCreate)
	dg.AddHandler(status.SetStatus)

	err = dg.Open()
	if err != nil {
		log.Fatal("Websocket error,", err.Error())
		return
	}

	// Register web svc handlers.
	go func() {
		storage.Migrate()
		webservice.Start(config.Port())
	}()

	// Until CTRL-C or another term signal.
	log.Println("Bot running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
