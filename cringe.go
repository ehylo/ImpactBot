package main

import (
	"fmt"
	"net/url"

	"github.com/bwmarrin/discordgo"
)

// this file is the story of my life lol

func handleCringe(_ *discordgo.Member, msg *discordgo.Message, _ []string) error {
	var rngCringe string
	err := DB.QueryRow("SELECT image FROM cringe ORDER BY RANDOM() LIMIT 1").Scan(&rngCringe)
	if err != nil {
		return err
	}
	reply := discordgo.MessageEmbed{
		Title: ":camera_with_flash:",
		Image: &discordgo.MessageEmbedImage{
			URL: rngCringe,
		},
		Color: prettyembedcolor,
	}
	_, err = discord.ChannelMessageSendEmbed(msg.ChannelID, &reply)
	go cringeReact(msg.ChannelID, msg.ID)
	return err
}

func handleAddCringe(caller *discordgo.Member, msg *discordgo.Message, args []string) error {

	if len(args) < 2 {
		if len(msg.Attachments) > 0 {
			return cringe(msg.Attachments[0].URL, msg.ChannelID, msg.ID)
		}
		return fmt.Errorf("error : no attachments / links found to add")
	}
	_, err := url.ParseRequestURI(args[1])
	if err != nil {
		return fmt.Errorf("invalid url scheme")
	}
	return cringe(args[1], msg.ChannelID, msg.ID)
}

func handleDelCringe(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("give url")
	}
	_, err := url.ParseRequestURI(args[1])
	if err != nil {
		return fmt.Errorf("invalid url scheme")
	}
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM cringe WHERE image = $1", args[1]).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("that is not cringe")
	}
	_, err = DB.Exec("DELETE FROM cringe WHERE image = $1", args[1])
	if err != nil {
		return err
	}
	return resp(msg.ChannelID, "cringe deleted successfully")
}

func cringe(url string, channelID string, messageID string) error {
	_, err := DB.Exec("INSERT INTO cringe(image) VALUES($1)", url)
	if err == nil {
		go cringeReact(channelID, messageID)
	}
	return err
}

func cringeReact(channelID string, messageID string) {
	_ = discord.MessageReactionAdd(channelID, messageID, "troll_crazy:826958873280774195")
}
