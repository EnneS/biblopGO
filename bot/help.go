package bot

import (
	"bytes"
	"fmt"
)

func Help(ctx Context) {
	cmds := ctx.CommandHandler.GetCmds()
	buffer := bytes.NewBufferString("Commands: \n")
	for cmdName, cmdStruct := range cmds {
		if len(cmdName) == 1 {
			continue
		}
		msg := fmt.Sprintf("\t %s%s - %s\n", ctx.CommandHandler.Prefix, cmdName, cmdStruct.GetHelp())
		buffer.WriteString(msg)
	}
	str := buffer.String()
	ctx.Reply(str[:len(str)-2], 0x0099FF)
}
