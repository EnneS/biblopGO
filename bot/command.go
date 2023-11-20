package bot

import (
	"fmt"
	"strings"

	"github.com/ennes/disgo"

	"github.com/bwmarrin/discordgo"
)

var (
	LastChannel string
)

type (
	CommandFn func(Context)
	Command   struct {
		command CommandFn
		help    string
	}
	Commands       map[string]Command
	CommandManager struct {
		Prefix string
		cmds   map[string]Command
		Player *disgo.Player
	}
)

func (handler *CommandManager) HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if len(m.Content) < len(handler.Prefix) || m.Content[:len(handler.Prefix)] != handler.Prefix {
		return
	}
	content := m.Content[len(handler.Prefix):]
	args := strings.Split(content, " ")
	name := args[0]
	if cmd, found := handler.Get(name); found {
		LastChannel = m.ChannelID
		ctx := NewContext(s, m.GuildID, m.ChannelID, m.Author, m, handler)
		ctx.Args = args[1:]
		(*cmd)(*ctx)
	}
}

func (handler *CommandManager) HandleInteraction(s *discordgo.Session, m *discordgo.InteractionCreate) {
	if m.Member.User.ID == s.State.User.ID {
		return
	}
	// Handle the interaction
	if cmd, found := handler.Get(m.Data.(discordgo.MessageComponentInteractionData).CustomID); found {
		LastChannel = m.ChannelID
		ctx := NewContext(s, m.GuildID, m.ChannelID, m.Member.User, nil, handler)
		(*cmd)(*ctx)
	}
	// Respond to the interaction
	err := s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	if err != nil {
		fmt.Println("Error whilst responding to interaction,", err)
	}
}

func NewCommandManager(prefix string, player *disgo.Player) *CommandManager {
	return &CommandManager{
		cmds:   make(Commands),
		Prefix: prefix,
		Player: player,
	}
}

func (handler CommandManager) GetCmds() Commands {
	return handler.cmds
}

func (handler CommandManager) Get(name string) (*CommandFn, bool) {
	cmd, found := handler.cmds[name]
	return &cmd.command, found
}

func (handler CommandManager) Register(name string, command CommandFn, helpmsg string) {
	// Massage the arguments into a "Full command"
	cmdstruct := Command{command: command, help: helpmsg}
	handler.cmds[name] = cmdstruct
	if len(name) > 1 {
		handler.cmds[name[:1]] = cmdstruct
	}
}

func (command Command) GetHelp() string {
	return command.help
}
