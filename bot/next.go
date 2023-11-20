package bot

func Next(ctx Context) {
	ctx.Player.Next(ctx.GuildID)
	if ctx.Message != nil {
		ctx.Discord.MessageReactionAdd(ctx.Message.ChannelID, ctx.Message.ID, "⏭")
	}
}
