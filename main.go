package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zeroidentidad/gopherbot/messages"
	"github.com/zeroidentidad/gopherbot/status"
)

var (
	Token string = "NzY3OTc5MDk4NzY4NDA4NTg2.X45yRQ.U-ue6pQabGEFui8pFqIqtFvSJ94"
)

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating session,", err)
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

	go func() {
		log.Fatal("Error starting http server", http.ListenAndServe(":3000", nil))
		_, _ = exec.Command("ping", "-t", "discord.com").Output()
	}()

	// Until CTRL-C or other term signal.
	fmt.Println("Bot running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	// Close session.
	dg.Close()
}
