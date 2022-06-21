package discordgo

import (
	"fmt"
	"log"
	"strings"

	GG "github.com/bwmarrin/discordgo"
)

// In order to the moderators to work with the bot and use his dangerous commands. They must have roles with "Moderator" or "Mod" or "Admin" or "The Owner" as their names.
func GetRidOfUserHandler(s *GG.Session, m *GG.MessageCreate) {
	var existingRolesIDInGuild []string

	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		log.Println(err.Error())
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	for _, v := range roles {
		if v.Name == "Moderator" || v.Name == "Mod" || v.Name == "Admin" || v.Name == "The Owner" {
			existingRolesIDInGuild = append(existingRolesIDInGuild, v.ID)
		}
	}

outer:
	for i, v := range m.Member.Roles {
	inner:
		for _, vv := range existingRolesIDInGuild {
			if v == vv {
				break outer
			} else {
				continue inner
			}
		}
		if i == len(m.Member.Roles)-1 {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s' roles have no administration permissions!", m.Author.Username))
			return
		}
	}

	if strings.Contains(m.Content, Kickcmd.name) {
		kickProcess(s, m)
		return
	}

	if strings.Contains(m.Content, SoftBancmd.name) {
		softBanProcess(s, m)
		return
	}

	if strings.Contains(m.Content, UnBancmd.name) {
		unBanProcess(s, m)
		return
	}

	if strings.Contains(m.Content, Bancmd.name) {
		banProcess(s, m)
		return
	}

	if m.Content == Closecmd.name {
		Closecmd.Log("")
		return
	}
}

// kickProcess kicks the mentioned users from the guild (server).
func kickProcess(s *GG.Session, m *GG.MessageCreate) {
	if len(m.Mentions) == 0 {
		s.ChannelMessageSend(m.ChannelID, errEmptySlice.Error())
		return
	}

	for _, v := range m.Mentions {
		err = s.GuildMemberDelete(m.GuildID, v.ID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %s", Kickcmd.value, v.Username))
		Kickcmd.Log("")
	}
}

// softBanProcess soft-banes (kick them and delet all messages since 7 days) the mentioned users from the guild (server).
func softBanProcess(s *GG.Session, m *GG.MessageCreate) {
	if len(m.Mentions) == 0 {
		s.ChannelMessageSend(m.ChannelID, errEmptySlice.Error())
		return
	}

	for _, v := range m.Mentions {
		err = s.GuildBanCreate(m.GuildID, v.ID, 7)
		if err != nil {
			log.Println(err)
			continue
		}

		err = s.GuildBanDelete(m.GuildID, v.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %s", SoftBancmd.value, v.Username))
		SoftBancmd.Log("")
	}
}

// unBanProcess unban the mentioned users (@name) from the guild (server).
func unBanProcess(s *GG.Session, m *GG.MessageCreate) {
	bans, err := s.GuildBans(m.GuildID, 100, "", "")
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	slice := strings.Split(m.Content, "@")

	if len(slice) <= 1 {
		s.ChannelMessageSend(m.ChannelID, errInvalidCommand.Error())
		log.Printf("GuildID: %s | ChannelID: %s | MessageID: %s", m.GuildID, m.ChannelID, m.ID)
		return
	}

	if len(bans) == 0 {
		s.ChannelMessageSend(m.ChannelID, "no one is banned here.")
		return
	}

	for _, v := range bans {
		if v.User.Username == slice[1] {
			err = s.GuildBanDelete(m.GuildID, v.User.ID)
			if err != nil {
				log.Println(err)
				continue
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %s", UnBancmd.value, slice[1]))
		}
		UnBancmd.Log("")
	}
}

// banProcess bans the mentioned users from the guild (server).
func banProcess(s *GG.Session, m *GG.MessageCreate) {
	if len(m.Mentions) == 0 {
		s.ChannelMessageSend(m.ChannelID, errEmptySlice.Error())
		return
	}

	for _, v := range m.Mentions {
		err = s.GuildBanCreate(m.GuildID, v.ID, 0)
		if err != nil {
			fmt.Println(err)
			continue
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %s", Bancmd.value, v.Username))
		Bancmd.Log("")
	}
}
