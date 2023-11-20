package bot

func Stop(ctx Context) {
	ctx.Player.Stop(ctx.GuildID)
	if ctx.Message != nil {
		ctx.Discord.MessageReactionAdd(ctx.Message.ChannelID, ctx.Message.ID, "⏹")
	}
}
