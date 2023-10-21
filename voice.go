package main

import (
	"github.com/bwmarrin/discordgo"
)

func onVoiceStateUpdate(session *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	if m.GuildID != baritoneServer {
		return
	}
	if m.ChannelID == "" || m.Deaf || m.SelfDeaf {
		_ = session.GuildMemberRoleRemove(baritoneServer, m.UserID, Voice.ID)
	} else {
		_ = session.GuildMemberRoleAdd(baritoneServer, m.UserID, Voice.ID)
	}
}

func checkDeservesInVoiceRole(userid string) bool {
	for _, guild := range discord.State.Guilds {
		for _, vs := range guild.VoiceStates {
			if vs.UserID == userid && !vs.Deaf && !vs.SelfDeaf {
				return true
			}
		}
	}
	return false
}
