package main

import (
	"github.com/bwmarrin/discordgo"
)

func ignore(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildMemberRoleAdd(baritoneServer, caller.User.ID, Ignored.ID)
	if err != nil {
		return err // fuckups get caught here
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, check)
}

func unignore(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	err := discord.GuildMemberRoleRemove(baritoneServer, caller.User.ID, Ignored.ID)
	if err != nil {
		return err // fuckups get caught here
	}
	return discord.MessageReactionAdd(msg.ChannelID, msg.ID, check)
}
