package discordgo

import (
	"fmt"
	"log"
	"strings"

	GG "github.com/bwmarrin/discordgo"
)

// In order to the moderators to work with the bot and use his dangerous commands. They must have roles with "Moderator" or "Mod" or "Admin" or "The Owner" as their name.
func RoleAddRemoveHandler(s *GG.Session, m *GG.MessageCreate) {
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

	if strings.Contains(m.Content, AddRolecmd.name) {
		addRolePorcess(s, m)
		return
	}

	if strings.Contains(m.Content, DeleteRolecmd.name) {
		delRoleProcess(s, m)
		return
	}
}

// addRolePorcess adds the mentioned role to the mentioneds users.
func addRolePorcess(s *GG.Session, m *GG.MessageCreate) {

	if len(m.MentionRoles) == 0 || len(m.Mentions) == 0 {
		s.ChannelMessageSend(m.ChannelID, errEmptySlice.Error())
		log.Println(errEmptySlice)
		return
	}

	for _, v := range m.Mentions {
		err = s.GuildMemberRoleAdd(m.GuildID, v.ID, m.MentionRoles[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			log.Println(err)
			continue
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %s", AddRolecmd.value, v.Username))
		AddRolecmd.Log("")
	}
}

// delRoleProcess removes the mentioned role from the mentioneds users.
func delRoleProcess(s *GG.Session, m *GG.MessageCreate) {

	if len(m.MentionRoles) == 0 || len(m.Mentions) == 0 {
		s.ChannelMessageSend(m.ChannelID, errEmptySlice.Error())
		log.Println(errEmptySlice)
		return
	}

	for _, v := range m.Mentions {
		err = s.GuildMemberRoleRemove(m.GuildID, v.ID, m.MentionRoles[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			log.Println(err)
			continue
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %s", DeleteRolecmd.value, v.Username))
		DeleteRolecmd.Log("")
	}
}
