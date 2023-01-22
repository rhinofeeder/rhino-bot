package bot

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/textproto"
	"os"
	"regexp"
	"rhino-bot/command"
	"strings"
	"time"
)

const (
	Format  = "Jan 2 15:04:05"
	PingMsg = "PING :tmi.twitch.tv"
	PongMsg = "PONG :tmi.twitch.tv\r\n"
)

var (
	MsgRegex = regexp.MustCompile(`^(.+):(\w+)!\w+@\w+\.tmi\.twitch\.tv (PRIVMSG) #\w+(?: :(.*))?$`)
	CmdRegex = regexp.MustCompile(`^!(\w+)\s?(.*)`)
)

type RhinoBot struct {
	Channel     string
	MsgRate     time.Duration
	Name        string
	Port        string
	PrivatePath string
	Server      string
	Token       string
	commands    map[string]command.Command
	conn        net.Conn
	startTime   time.Time
}

func (rb *RhinoBot) RegisterCommand(commands ...command.Command) {
	if rb.commands == nil {
		rb.commands = make(map[string]command.Command, len(commands))
	}
	for _, c := range commands {
		rb.commands[c.Name()] = c
	}
}

func (rb *RhinoBot) Connect() {
	fmt.Printf("[%s] Connecting to %s...\n", timeStamp(), rb.Server)

	// makes connection to Twitch IRC server
	connection, err := net.Dial("tcp", rb.Server+":"+rb.Port)
	count := 1
	for err != nil {
		count++
		fmt.Printf("[%s] Cannot connect to %s, retrying (Attempt %v).\n", timeStamp(), rb.Server, count)
		connection, err = net.Dial("tcp", rb.Server+":"+rb.Port)
	}

	rb.startTime = time.Now()
	rb.conn = connection

	fmt.Printf("[%s] Connected to %s!\n", timeStamp(), rb.Server)
}

func (rb *RhinoBot) Disconnect() {
	_ = rb.conn.Close()
	upTime := time.Now().Sub(rb.startTime).Seconds()
	fmt.Printf("[%s] Closed connection from %s! | Live for: %fs\n", timeStamp(), rb.Server, upTime)
}

func (rb *RhinoBot) HandleChat() error {
	fmt.Printf("[%s] Watching #%s...\n", timeStamp(), rb.Channel)

	tp := textproto.NewReader(bufio.NewReader(rb.conn))
	for {
		line, err := tp.ReadLine()
		if err != nil {
			rb.Disconnect()
			return errors.New("failed to read line from channel, disconnected")
		}

		fmt.Printf("[%s] %s\n", timeStamp(), line)

		if line == PingMsg {
			_, _ = rb.conn.Write([]byte(PongMsg))
		} else {
			matches := MsgRegex.FindStringSubmatch(line)
			if matches != nil {
				msgType := matches[3]
				badges := strings.Split(matches[1], ";")[1]

				if msgType == "PRIVMSG" {
					msg := matches[4]

					// parse commands from user message
					cmdMatches := CmdRegex.FindStringSubmatch(msg)
					if cmdMatches != nil {
						cmd := cmdMatches[1]

						if registeredCommand := rb.commands[cmd]; registeredCommand != nil {
							if registeredCommand.RequiresMod() {
								if strings.Contains(badges, "broadcaster") || strings.Contains(badges, "moderator") {
									rb.handleCommand(registeredCommand, cmdMatches[2])
								}
							} else {
								rb.handleCommand(registeredCommand, cmdMatches[2])
							}
						}
					}
				}
			}
		}
	}
}

func (rb *RhinoBot) handleCommand(registeredCommand command.Command, message string) {
	var sayErr error
	if response, err := registeredCommand.Handler(message); err != nil {
		sayErr = rb.Say("Oops, there was an issue!")
		fmt.Printf("Error: %v\n", err)
	} else {
		sayErr = rb.Say(response)
	}

	if sayErr != nil {
		fmt.Printf("Error in Say(): %v\n", sayErr)
	}

	time.Sleep(rb.MsgRate)
}

func (rb *RhinoBot) JoinChannel() {
	fmt.Printf("[%s] Joining #%s...\n", timeStamp(), rb.Channel)
	_, _ = rb.conn.Write([]byte("PASS " + rb.Token + "\r\n"))
	_, _ = rb.conn.Write([]byte("NICK " + rb.Name + "\r\n"))
	_, _ = rb.conn.Write([]byte("JOIN #" + rb.Channel + "\r\n"))
	_, _ = rb.conn.Write([]byte("CAP REQ :twitch.tv/tags\r\n"))

	fmt.Printf("[%s] Joined #%s as @%s!\n", timeStamp(), rb.Channel, rb.Name)
}

func (rb *RhinoBot) ReadCredentials() error {
	token, err := os.ReadFile(rb.PrivatePath)
	if err != nil {
		return err
	}

	rb.Token = string(token)

	return nil
}

func (rb *RhinoBot) Say(msg string) error {
	if msg == "" {
		return errors.New("msg was empty")
	}

	// check if message is too large for IRC
	if len(msg) > 512 {
		return errors.New("msg exceeded 512 bytes")
	}

	_, err := rb.conn.Write([]byte(fmt.Sprintf("PRIVMSG #%s :%s\r\n", rb.Channel, msg)))
	if nil != err {
		return err
	}
	return nil
}

func (rb *RhinoBot) Start() {
	err := rb.ReadCredentials()
	if nil != err {
		fmt.Println(err)
		fmt.Println("Aborting...")
		return
	}

	for {
		rb.Connect()
		rb.JoinChannel()
		err = rb.HandleChat()
		if nil != err {

			// attempts to reconnect upon unexpected chat error
			time.Sleep(1000 * time.Millisecond)
			fmt.Println(err)
			fmt.Println("Starting bot again...")
		} else {
			return
		}
	}
}

func timeStamp() string {
	return time.Now().Format(Format)
}
