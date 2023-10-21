package main

import (
	"regexp"
)

type Reply struct {
	pattern         string
	unless          string
	message         string
	excludeChannels []string
	excludeRoles    []Role
	onlyChannels    []string
	onlyRoles       []Role
	regex           *regexp.Regexp
	notRegex        *regexp.Regexp
}

type currencyinfo struct {
	Amount      int64  `json:"premium_amount"`
	DisplayName string `json:"display_name"`
	Symbol      string `json:"symbol"`
}

var replies = []Reply{
	{
		pattern: `faq|question|tutorial`,
		message: "[Setup/Install FAQ](https://github.com/impactdevelopment/impactclient/wiki/Setup-FAQ)\n[Usage FAQ](https://github.com/impactdevelopment/impactclient/wiki/Usage-FAQ)",
	},
	{
		pattern: `baritone\s*setting`,
		message: "[Baritone settings list and documentation](https://baritone.leijurv.com/baritone/api/Settings.html#field.detail)",
	},
	{
		pattern: `screenshot`,
		message: "[How to take a screenshot in Minecraft](https://www.minecraft.net/en-us/article/screenshotting-guide)",
	},
	{
		pattern: `use\s*baritone|baritone\s*(usage|command)|[^u]\.b|goal|goto|path`,
		message: "Please read the [Baritone usage guide](https://github.com/cabaletta/baritone/blob/master/USAGE.md)",
	},
	{ // GitHub repo for issues
		pattern: `issue|bug|crash|error|suggest(ion)?s?|feature|enhancement`,
		message: "Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactIssues/issues) to report issues/suggestions!",
	},
	{ // Non-donator help redirect
		pattern:         `help|support`,
		message:         "Switch to the <#" + help + "> channel!",
		excludeChannels: []string{help},
	},
	{ // Information on macros
		pattern: `macros?`,
		message: "Macros are in-game chat commands, they can be accessed in-game by clicking on the Impact button, then Macros.",
	},
	{ // Notice about "hacking"
		pattern: `hack(s|ing|er|client)?`,
		message: "**Impact is not a hacked client**, it is designed as a utility mod (e.g. for anarchy servers).\nSupport will not be provided to users who utilise Impact on servers that do not allow it.\nPlease also note that the discussion of \"hacks\" in this Discord server is prohibited to comply with the [Discord Community Guidelines](https://discord.com/guidelines)",
	},
	{ // Info on using schematics
		pattern: `schematics?`,
		message: "Schematic file **MUST** be made in a 1.12.2 world or prior.\n1) Place the .schematic file into `.minecraft/schematics`.\n2) Ensure all the blocks are in your hotbar.\n3) Type `#build name.schematic`.",
	},
	{ // Info on using cracked launchers
		pattern: `((crack|cracked) (launcher|account|game|minecraft))|(terramining|shiginima|(t(-|)launcher))`,
		message: "Impact does not support cracked launchers. You can attempt to use the unstable Forge version, but no further support will be provided.",
	},
	{ // Downloads for JRE
		pattern: `java.*(download|runtime|environment)`,
		message: "[Downloads for Java Runtime Environment](https://www.java.com/download/)",
	},
	{ // How to use Impact's automine function
		pattern: `how.+(mine|auto\s*mine)`,
		message: "You can mine a specific type of block(s) by typing `#mine [number of blocks to mine] <ID> [<ID>]` in chat.\nYou can find a list of block ID names [here](https://www.digminecraft.com/lists/)",
	},
	{ // Information on using Impact with modpacks
		pattern: `(modpack|\bftb\b|rlcraft|skyfactory|valhelsia|pixelmon|sevtech)`,
		message: "Impact is generally incompatible with modpacks and support will not be provided if you encounter bugs with them. It's likely your game will just crash on startup.",
	},
	{
		pattern: `good\s*bot`,
		message: "tnyak yow *nuwzzwes yoww necky wecky*",
	},
	{
		pattern: `((anti(-|\s*)(kb|knockback))|velocity)`,
		message: "**Velocity**, also known as **Anti-knockback**, is a module under \"Movement\" that prevents the player from taking knockback.",
	},
	{
		pattern: `(gui|r(-|\s)shift|module|(open|close|show|hide)\s*impact)`,
		message: "To open or close the Impact GUI, press the `rshift` key, located below `enter`.",
	},
}
