package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ennes/disgo"
)

var virtualPlayer *discordgo.Message

func SetupEventHandlers(p *disgo.Player) {
	p.On("play", func(e disgo.Event) {
		pe := e.(*disgo.PlayEvent)

		if virtualPlayer != nil {
			p.DiscordSession.ChannelMessageDelete(LastChannel, virtualPlayer.ID)
		}

		msg, err := p.DiscordSession.ChannelMessageSendComplex(LastChannel, &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				{
					Color: 0x0099FF,
					Author: &discordgo.MessageEmbedAuthor{
						Name: fmt.Sprintf("%s - %s | [%s]", pe.Song.Title(), pe.Song.Author(), pe.Song.Duration()),
					},
					Description: "🎶",
				},
			},
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.Button{
							Label:    "⏭",
							Style:    discordgo.PrimaryButton,
							CustomID: "skip",
						},
						&discordgo.Button{
							Label:    "⏹",
							Style:    discordgo.DangerButton,
							CustomID: "stop",
						},
					},
				},
			},
		})
		if err != nil {
			fmt.Println("Error whilst sending message,", err)
		}
		virtualPlayer = msg
	})

	p.On("stop", func(e disgo.Event) {
		if virtualPlayer != nil {
			p.DiscordSession.ChannelMessageDelete(LastChannel, virtualPlayer.ID)
		}
	})
}
