package main

import (
	"biblopGO/bot"
	"fmt"

	"github.com/ennes/disgo"

	"github.com/bwmarrin/discordgo"
)

var (
	commandHandler *bot.CommandManager
)

var LastChannel string

func main() {
	dg, err := discordgo.New("Bot " + "MTE3NDg0MzU4Mjk4MTYxMTU2MA.G4ggaL.ym36lciEBQBo2DJK4tQZ44LrXY2oHkbB1cO63E")
	if err != nil {
		panic(err)
	}
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildPresences | discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates | discordgo.IntentMessageContent | discordgo.IntentsGuilds)

	// query := "u1V8YRJnr4Q"
	// guild := "myGuild"
	// chanId := "myChannel"
	// query = "https://www.youtube.com/watch?v=-GEHyAfV4OI&list=PLDBVggJyVNfvUfBgBlrRTnjON0BXcAjpf"
	p := &disgo.Player{
		DiscordSession: dg,
	}

	p.Init()
	bot.SetupEventHandlers(p)

	commandHandler = bot.NewCommandManager("&", p)
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
