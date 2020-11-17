package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zeroidentidad/gopherbot/config"
	"github.com/zeroidentidad/gopherbot/messages"
	"github.com/zeroidentidad/gopherbot/status"
)

func main() {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Session error:", err)
		return
	}

	dg.Debug = true

	// Register handlers.
	dg.AddHandler(messages.MessageCreate)
	dg.AddHandler(status.SetStatus)

	err = dg.Open()
	if err != nil {
		fmt.Println("Websocket connection error,", err)
		return
	}

	/*go func() {
		log.Fatal("Error http server", http.ListenAndServe(":"+config.Port(), nil))
	}()*/

	// Until CTRL-C or other term signal.
	fmt.Println("Bot running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	// Close session.
	dg.Close()
}
