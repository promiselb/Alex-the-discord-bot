package discordgo

import (
	"fmt"
	"log"
	"strings"

	GG "github.com/bwmarrin/discordgo"
)

func MuteUnmuteMemberHandler(s *GG.Session, m *GG.MessageCreate) {

	if strings.Contains(m.Content, "unmute") {
		unmuteProcess(s, m)
		return
	}

	if strings.Contains(m.Content, "mute") {
		muteProcess(s, m)
		return
	}
}

// muteProcess adds the Muted role to the mentioned users and adds the Bro role to prevent them from sending messages unitl they get unmuted manually.
func muteProcess(s *GG.Session, m *GG.MessageCreate) {
	for _, v := range m.Mentions {
		err = s.GuildMemberRoleAdd(m.GuildID, v.ID, "866627068226043954") // the id of "Muted" role.
		if err != nil {
			log.Println(err)
			continue
		}

		err = s.GuildMemberRoleRemove(m.GuildID, v.ID, "863551384371724308") // the id of "Bro" role.
		if err != nil {
			log.Println(err)
			continue
		}

		s.ChannelMessageSend(m.ChannelID, MuteTextcmd.value+v.Mention())
		MuteTextcmd.Log(v.Username)
	}
}

// unmuteProcess adds the Bro role to the mentioned users and adds the Muted role to allow them a new from sending messages.
func unmuteProcess(s *GG.Session, m *GG.MessageCreate) {
	for _, v := range m.Mentions {
		err = s.GuildMemberRoleRemove(m.GuildID, v.ID, "866627068226043954") // the id of "Muted" role.
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = s.GuildMemberRoleAdd(m.GuildID, v.ID, "863551384371724308") // the id of "Bro" role.
		if err != nil {
			fmt.Println(err)
			continue
		}
		s.ChannelMessageSend(m.ChannelID, UnMuteTextcmd.value+v.Mention())
		UnMuteTextcmd.Log("")
	}
}
