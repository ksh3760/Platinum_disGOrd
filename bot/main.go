package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"bot/common/constant"
	"bot/module/lotto"
)

var (
	G_TOKEN  string = constant.G_DISCORD_BOT_TOKEN
	G_PREFIX string = constant.G_BOT_CALL_PREFIX
)

func main() {
	var (
		sDg         *discordgo.Session
		sErr        error
		sSystemCall chan os.Signal
		sToken      []byte
	)

	if G_TOKEN == "" {
		G_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
	}

	sToken, sErr = ioutil.ReadFile(G_TOKEN)
	if sErr != nil {
		fmt.Println("Can not read token")
		return
	}

	sDg, sErr = discordgo.New("Bot " + string(sToken))
	if sErr != nil {
		fmt.Println("Error creating Discord session,", sErr)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	// MessageCreate 이벤트에 대한 콜백으로 messageCreate func를 등록합니다.
	sDg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	// 이 예제는 메세지 이벤트 수신에서만 신경쓴다.
	sDg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	// 디스코드에 대한 웹 소켓 연결을 열고 수신을 시작한다.
	sErr = sDg.Open()
	if sErr != nil {
		fmt.Println("Error opening connection,", sErr)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	// CTRL-C 또는 기타 용어 신호가 수신될 때까지 여기서 기다린다.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sSystemCall = make(chan os.Signal, 1)
	signal.Notify(sSystemCall, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sSystemCall

	// Cleanly close down the Discord session.
	sDg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == G_PREFIX+"ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == G_PREFIX+"pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	// 예제) 로또 번호 생성
	if m.Content == G_PREFIX+"로또번호" {
		s.ChannelMessageSend(m.ChannelID, lotto.GenLottoNum())
	}

	// 명령어 목록을 출력한다.
	if m.Content == G_PREFIX+"help" {
		s.ChannelMessageSend(m.ChannelID, "명령어 모음")
	}
}
