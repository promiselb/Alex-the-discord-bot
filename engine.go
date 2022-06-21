package discordgo

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This where the whole story beggins.
var Alex, err = discordgo.New("Bot " + Token)
// The "Token" variable won't be provided. Check https://discord.com/developers/applications to make your own discord bot and get its token to work.

// The Handlers.
func Handlers() {
	Alex.AddHandler(MessageCreate)
	Alex.AddHandler(ShrugfHandler)
}

// The engine.
func ConnectToDiscord() {

	if err != nil {
		log.Fatal(err)
	}

	Handlers()

	Alex.Identify.Intents = discordgo.IntentGuildMessages

	err = Alex.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer Alex.Close()

	fmt.Println("Alex online.")

	sc := make(chan struct{})
	<-sc
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if !strings.HasPrefix(m.Content, CommandPrefix) {
		return
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := ingoreWhiteSpaces(m.Content)

	if strings.HasPrefix(m.Content, CommandPrefix+"exe") {
		GetRidOfUserHandler(s, m)
		return
	}

	if strings.HasPrefix(m.Content, CommandPrefix+"role") {
		RoleAddRemoveHandler(s, m)
		return
	}

	if strings.Contains(m.Content, "mute") {
		MuteUnmuteMemberHandler(s, m)
		return
	}

	if strings.Contains(msg, "help") {
		helpProcess(msg, s, m)
		return
	}

	if strings.Contains(msg, "compare") {
		compareProcess(s, m)
		return
	}

	if m.Content == Hicmd.name {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %s!", Hicmd.value, m.Author.Username))
		Hicmd.Log(m.Author.Username)
		return
	}

	if strings.Contains(msg, Testcmd.name) {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprint(randomBool()))
		Testcmd.Log("")
		return
	}

	if _, ok := commands[msg]; !ok {
		s.ChannelMessageSend(m.ChannelID, errInvalidCommand.Error()+`. Call `+"`$help`")
		log.Println("invalid command")
		return
	}

	s.ChannelMessageSend(m.ChannelID, commands[msg].value)
	commands[msg].Log("")
}

func helpProcess(msg string, s *discordgo.Session, m *discordgo.MessageCreate) {
	slice := strings.Split(msg, " ")

	ch, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		log.Printf("Failed to create a private channel with %s\n %v", m.Author.ID, err)
		return
	}

	usersWithPrvtChannelsIDS[m.Author.Username+m.Author.Discriminator] = ch.ID

	if len(slice) == 1 {
		s.ChannelMessageSend(ch.ID, sendItAsCode(help(), "txt"))
		Helpcmd.Log("")
		return
	}

	if slice[1] == Helpcmd.name {
		s.ChannelMessageSend(ch.ID, sendItAsCode(Helpcmd.description, "txt"))
		return
	}

	if _, ok := commands[slice[1]]; !ok {
		s.ChannelMessageSend(ch.ID, sendItAsCode(help(), "txt"))
		Helpcmd.Log("")
		return
	} else {
		s.ChannelMessageSend(ch.ID, sendItAsCode(commands[slice[1]].description, "txt"))
		log.Printf("%s %s Executed.", Helpcmd.name, commands[slice[1]].name)
	}
}

func compareProcess(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Count(m.Content, "<") != 2 || strings.Count(m.Content, ">") != 2 {
		s.ChannelMessageSend(m.ChannelID, errMissingKeyLiteral.Error())
		return
	}
	msg := splitArgs(m.Content)

	if msg[2] == "" {
		s.ChannelMessageSend(m.ChannelID, errEmptyArg.Error())
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprint(msg[1] == msg[2]))
	Comparecmd.Log("")
}
