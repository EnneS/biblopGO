package main

import (
	"biblopGO/bot"
	"fmt"
	"os"

	"github.com/ennes/disgo"

	"github.com/bwmarrin/discordgo"
)

var (
	commandHandler *bot.CommandManager
)

var LastChannel string

func main() {
	token := os.Getenv("TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildPresences | discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates | discordgo.IntentMessageContent | discordgo.IntentsGuilds)

	p := &disgo.Player{
		DiscordSession: dg,
	}

	p.Init()
	bot.SetupEventHandlers(p)

	commandHandler = bot.NewCommandManager("!", p)
	registerCommands()
	dg.AddHandler(commandHandler.HandleMessage)
	dg.AddHandler(commandHandler.HandleInteraction)
	dg.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		discord.UpdateGameStatus(0, "!h | blo-blo-blo-blop")
		fmt.Println("Ready!")
	})
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	select {}
}

func registerCommands() {
	commandHandler.Register("help", bot.Help, "Donne la liste des commandes disponibles")
	commandHandler.Register("play", bot.Play, "Plays a song")
	commandHandler.Register("skip", bot.Next, "Skips the current song")
	commandHandler.Register("stop", bot.Stop, "Stops the player")
}
