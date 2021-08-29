package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zeroidentidad/gopherbot/botservice/messages"
	"github.com/zeroidentidad/gopherbot/global"
	"github.com/zeroidentidad/gopherbot/webservice"
	"github.com/zeroidentidad/gopherbot/webservice/storage"
)

func main() {
	dg, err := discordgo.New("Bot " + global.TOKEN)
	if err != nil {
		log.Fatal("Session error:", err.Error())
		return
	}
	// dg.Debug = true

	// Register bot handlers.
	dg.AddHandler(messages.MessageCmd)
	dg.AddHandler(messages.SetStatus)

	err = dg.Open()
	if err != nil {
		log.Fatal("Websocket error,", err.Error())
		return
	}

	// Register web handlers.
	go func() {
		storage.Migrate()
		webservice.Start(global.Port())
	}()

	// Until CTRL-C or another term signal.
	log.Println("Bot running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
