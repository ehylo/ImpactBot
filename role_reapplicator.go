package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// note: this is not for mutes. rekt does that
// this is for roles determined by the central db
func onUserJoin3(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	if m.GuildID != baritoneServer {
		return
	}
}

func isYes(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	return string(data) == "yes"
}
