package discordgo

import (
	"strings"

	GG "github.com/bwmarrin/discordgo"
)

func ShrugfHandler(s *GG.Session, m *GG.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, Shrugcmd.name) {
		shrugProcess(s, m)
		return
	}
}

// used send a new message to the same channel wtih ¯\_(ツ)_/¯ at the end.
func shrugProcess(s *GG.Session, m *GG.MessageCreate) {
	var ss string

	for i := 0; i < len(m.Content); i++ {
	slice := strings.Split(m.Content, Shrugcmd.name)
	if len(slice) <= 1 {
		return
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintln(slice[0]+slice[1], Shrugcmd.value))
	Shrugcmd.Log("")
}
