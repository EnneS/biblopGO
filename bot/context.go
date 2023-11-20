package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ennes/disgo"
)

type Context struct {
	Discord       *discordgo.Session
	GuildID       string
	TextChannelID string
	VoiceChannel  *discordgo.Channel
	User          *discordgo.User
	Message       *discordgo.MessageCreate
	Args          []string

	CommandHandler *CommandManager
	Player         *disgo.Player
}

func NewContext(discord *discordgo.Session, guild string, textChannel string,
	user *discordgo.User, message *discordgo.MessageCreate, commandHandler *CommandManager) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.GuildID = guild
	ctx.TextChannelID = textChannel
	ctx.User = user
	ctx.Message = message
	ctx.CommandHandler = commandHandler
	ctx.Player = commandHandler.Player
	return ctx
}

func (ctx Context) RichReply(content *discordgo.MessageEmbed) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannelID, content)
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return nil
	}
	return msg
}

func (ctx Context) Reply(content string, color int) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannelID, &discordgo.MessageEmbed{
		Color:       color,
		Description: content,
	})
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return nil
	}
	return msg
}

func (ctx *Context) GetVoiceChannel() *discordgo.Channel {
	if ctx.VoiceChannel != nil {
		return ctx.VoiceChannel
	}
	guild, err := ctx.Discord.State.Guild(ctx.GuildID)
	if err != nil {
		fmt.Println("Error whilst getting guild,", err)
		return nil
	}
	for _, channel := range guild.VoiceStates {
		if channel.UserID == ctx.User.ID {
			ctx.VoiceChannel, err = ctx.Discord.State.Channel(channel.ChannelID)
			if err != nil {
				fmt.Println("Error whilst getting channel,", err)
				return nil
			}
			return ctx.VoiceChannel
		}
	}
	return nil
}
