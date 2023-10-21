package main

//noinspection GoUnusedConst
const (
	baritoneServer   = "826950651690745876"
	prettyembedcolor = 0x81C3FF

	trash = "🗑"
	check = "✅"

	announcements = "826950736574414859"
	general       = "826950651711979536"
	offtopic      = "826950651711979537"
	gaming        = "826950652160114708"
	bots          = "826950652160114709"
	help          = "826950652160114710"
	dev           = "826950652160114711"
	bariBotLog    = "826950652516106245"
)

var (
	Developer = Role{"826950651698610204", "Developer"}
	Admin     = Role{"826950651698610204", "Admin"}
	Moderator = Role{"826950651698610201", "Moderator"}
	Helper    = Role{"826950651698610200", "Helper"}
	Voice     = Role{"826954171961704458", "Voice"}
	Ignored   = Role{"826954212059381810", "Ignored"}
)

type Role struct {
	// Discord id
	ID   string
	Name string
}

func RolesToIDs(roles []Role) []string {
	var ids []string
	for _, role := range roles {
		ids = append(ids, role.ID)
	}
	return ids
}

var muteRoles = map[string]string{
	"": "826950651690745885",
	/*	general:       "352144990606196749",
		offtopic:         "669632558551793666",
		gaming:        "624971877424693288",
		bots:           "640263788985188362",
		help:          "230803433752363020",
		dev:   "669632988686057472", */ //Add per channel mute roles
}

func (r Role) Mention() string {
	return "<@&" + r.ID + ">"
}
