package discordgo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// Command type for: [commandPrefix][command] (arguments) => <value>
//
// (...) are not always asked for.
type Command struct {
	// The name of the command used in the chat with the API.
	name string

	// The value returned from the API to the user.
	value string

	// The description of the command.
	description string
}

// Log logs the command to indicate when its executed. This is used with error tracking.
func (cmd *Command) Log(s string) {

	if cmd.name == Kickcmd.name {
		return
	}

	if cmd.name == Closecmd.name {
		log.Println("Alex offline.")
		Alex.Close()
		time.Sleep(time.Second * 2)
		log.Printf("%s Executed. %s\n", cmd.name, s)
		os.Exit(0)
	}

	log.Printf("%s Executed. %s\n", cmd.name, s)
}

// Discord things variables.
var (

	// The application token from https://discord.com/developers/applications.
	Token string

	// The command prefix.
	CommandPrefix string = "$"

	// The id of BFF's server. Which's my server :)
	MyServerID = "863540139032707082"
)

// Errors.
var (
	errEmptyArg          = errors.New("error: Not enough arguments?")
	errMissingKeyLiteral = errors.New("error: Missing key literal (< or >)?")
	errInvalidCommand    = errors.New("error: invalid command")
	errEmptySlice        = errors.New("error: Empty slice (no users or roles mentioned)?")
)

// List of available commands.
var (
	Closecmd = &Command{ // 1
		name:        CommandPrefix + "close",
		value:       "Alex offline.",
		description: "Returns Alex to sleep.",
	}

	Hicmd = &Command{ // 2
		name:        CommandPrefix + "hi",
		value:       "Hello",
		description: "Greeting.",
	}

	Breakfastcmd = &Command{ // 3
		name:        CommandPrefix + "breakfast",
		value:       "Let's eat some fried eggs 🍳🥚 and cheese 🧀!",
		description: "Breakfast of cheese and fried eggs!",
	}

	Lunchcmd = &Command{ // 4
		name:        CommandPrefix + "lunch",
		value:       "Let's eat pizza 🍕!",
		description: "Pizza time!",
	}

	Dinnercmd = &Command{ //5
		name:        CommandPrefix + "dinner",
		value:       "Let's eat some meat 🥩 and drink something 🍸!",
		description: "Dinner time!",
	}

	Partycmd = &Command{ // 6
		name:        CommandPrefix + "party",
		value:       "Let's enjoy some gâteau 🎂 and live our life!",
		description: "Party time!",
	}

	Pingcmd = &Command{ // 7
		name:        CommandPrefix + "ping",
		value:       "pang",
		description: "Ping pang pong!",
	}

	Pangcmd = &Command{ // 8
		name:        CommandPrefix + "pang",
		value:       "pong",
		description: "Ping pang pong!",
	}

	Pongcmd = &Command{ // 9
		name:        CommandPrefix + "pong",
		value:       "ping",
		description: "Ping pang pong!",
	}

	Shrugcmd = &Command{ // 10
		name:        CommandPrefix + "shrug",
		value:       `¯\_(ツ)_/¯`,
		description: `Adds ¯\_(ツ)_/¯ to the end of the message.`,
	}

	Comparecmd = &Command{ // 11
		name:        CommandPrefix + "compare",
		value:       "true or false",
		description: `Compare two arguments like: $compare <"arg1"> <"arg2">`,
	}

	Kickcmd = &Command{ // 12
		name:        CommandPrefix + "exe_kick",
		value:       "Alex kicked",
		description: "Kicks all users mentioned. \n$exe_kick <@user[1]>...<@user[n]>",
	}

	UnBancmd = &Command{ // 13
		name:        CommandPrefix + "exe_unban",
		value:       "Alex removed ban for",
		description: "Removes a user ban from a server. \n$exe_unban <username>",
	}

	Bancmd = &Command{ // 14
		name:        CommandPrefix + "exe_ban",
		value:       "Alex banned",
		description: "Ban all users mentioned. \n$exe_ban <@user[1]>...<@user[n]>",
	}

	AddRolecmd = &Command{
		name:        CommandPrefix + "role_add",
		value:       "Alex added role to",
		description: "Adds the rule to mentioned users. \n$role_add <@role> <@user[1]>...<@user[n]>",
	}

	DeleteRolecmd = &Command{
		name:        CommandPrefix + "role_del",
		value:       "Alex removed role from",
		description: "Removes the role from mentioned users. \n$role_del <@role> <@user[1]>...<@user[n]",
	}

	MuteTextcmd = &Command{ // 15
		name:        CommandPrefix + "mute",
		value:       ` \(°×°)/ `,
		description: "Mutes all users mentioned of a guild. \n$mute <@user[1]>...<@user[n]>",
	}

	UnMuteTextcmd = &Command{ // 16
		name:        CommandPrefix + "unmute",
		value:       ` \(°o°)/ `,
		description: "Unmutes all users mentioned of a guild. \n$unmute <@user[1]>...<@user[n]>",
	}

	SoftBancmd = &Command{ // 17
		name:        CommandPrefix + "exe_softban",
		value:       "Alex soft-banned",
		description: "Soft-banes all users mentioned of a guild (Kick the users and delete all thier messages). \n$exe_softban <@user[1]>...<@user[n]>",
	}

	Testcmd = &Command{ // 18
		name:        CommandPrefix + "test",
		value:       fmt.Sprint(randomBool()),
		description: "$test <sentence>?",
	}

	Helpcmd = &Command{ // 19
		name:        CommandPrefix + "help",
		value:       help(),
		description: "Lists all commands and theirs descriptions.",
	}
)

// All commands stored in a map by their names except for $help command.
var commands = map[string]*Command{
	Closecmd.name: Closecmd,

	Hicmd.name: Hicmd,

	Breakfastcmd.name: Breakfastcmd,

	Lunchcmd.name: Lunchcmd,

	Dinnercmd.name: Dinnercmd,

	Partycmd.name: Partycmd,

	Shrugcmd.name: Shrugcmd,

	Kickcmd.name: Kickcmd,

	SoftBancmd.name: SoftBancmd,

	Bancmd.name: Bancmd,

	UnBancmd.name: UnBancmd,

	AddRolecmd.name: AddRolecmd,

	DeleteRolecmd.name: DeleteRolecmd,

	MuteTextcmd.name: MuteTextcmd,

	UnMuteTextcmd.name: UnMuteTextcmd,

	Pingcmd.name: Pingcmd,

	Pangcmd.name: Pingcmd,

	Pongcmd.name: Pongcmd,

	Comparecmd.name: Comparecmd,

	Testcmd.name: Testcmd,
}

// Used to send the $help command value to the user on private to not distribute the server.
// usersWithPrvtChannelsIDS[name#xxx] = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
var usersWithPrvtChannelsIDS map[string]string
