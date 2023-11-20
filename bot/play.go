package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Play(ctx Context) {
	ctx.Discord.MessageReactionAdd(ctx.Message.ChannelID, ctx.Message.ID, "⏳")
	voiceChannel := ctx.GetVoiceChannel()
	if voiceChannel == nil {
		ctx.Reply("Vous devez être dans un salon vocal.", 0xFF0000)
		return
	}
	query := strings.Join(ctx.Args, " ")
	songs, err := ctx.Player.Play(ctx.GuildID, voiceChannel.ID, query)
	if err != nil {
		ctx.Reply("🙅‍♂️ Impossible de jouer la musique (peut-être que le lien est invalide ?)", 0xFF0000)
		return
	}
	if len(songs) == 0 {
		ctx.Reply("🙅‍♂️ Aucune musique trouvée.", 0xFF0000)
		return
	}
	queue := ctx.Player.GetQueue(ctx.GuildID)
	var description string
	if len(songs) == 1 {
		description = fmt.Sprintf("**%s** par **%s** [%s]", songs[0].Title(), songs[0].Author(), songs[0].Duration())
	} else {
		for i, song := range songs {
			description += fmt.Sprintf("%d. **%s** par **%s** [%s]\n", i+1, song.Title(), song.Author(), song.Duration())
		}
	}
	ctx.RichReply(&discordgo.MessageEmbed{
		Color: 0x0099FF,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    fmt.Sprintf("%s | Ajouté en #%d", ctx.User.Username, queue.Len()),
			IconURL: ctx.User.AvatarURL("32"),
		},
		Description: description,
	})
	ctx.Discord.MessageReactionsRemoveAll(ctx.Message.ChannelID, ctx.Message.ID)
	ctx.Discord.MessageReactionAdd(ctx.Message.ChannelID, ctx.Message.ID, "🎶")
}
