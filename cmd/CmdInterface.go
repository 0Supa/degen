package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

type Command struct {
	Name           string
	Guilds         []discord.GuildID
	DiscordData    api.CreateCommandData
	DiscordHandler cmdroute.CommandHandlerFunc
}

var CommandMap = make(map[string]Command)

func RegisterCommand(command Command) {
	CommandMap[command.Name] = command
}

func ErrorResponse(err error) *api.InteractionResponseData {
	str := err.Error()
	var data json.RawMessage
	json.Unmarshal([]byte(str), &data)

	var lang string
	if data != nil {
		m, _ := json.MarshalIndent(data, "", "  ")
		str = string(m)
		lang = "json"
	}

	return &api.InteractionResponseData{
		Embeds: &[]discord.Embed{{
			Title:       "😳 OOPSIE WOOPSIE!!",
			Footer:      &discord.EmbedFooter{Text: "Uwu We made a fucky wucky!!"},
			Description: "Error\n" + CodeBlock(lang, str),
		}},
		// Flags:           discord.EphemeralMessage,
		AllowedMentions: &api.AllowedMentions{},
	}
}

func Response(format string, a ...any) *api.InteractionResponseData {
	return &api.InteractionResponseData{
		AllowedMentions: &api.AllowedMentions{},
		Content:         option.NewNullableString(fmt.Sprintf(format, a...)),
	}
}
