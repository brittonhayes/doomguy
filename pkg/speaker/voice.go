package speaker

import (
	"fmt"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/arikawa/v2/voice"
	"github.com/diamondburned/arikawa/v2/voice/voicegateway"
	log "github.com/sirupsen/logrus"
)

type Speaker struct {
	Ctx     *bot.Context
	Voice   *voice.Session
	Token   string
	Channel discord.ChannelID
}

// Setup demonstrates the CanSetup interface. This function will never be parsed
// as a callback of any event.
func (s *Speaker) Setup(sub *bot.Subcommand) {
	sub.Command = "speaker"
	sub.Description = "Join voice channels, play music, or youtube videos"
	sub.ChangeCommandInfo("Join", "join", "Join the voice channel")
	sub.ChangeCommandInfo("Leave", "leave", "Leave the voice channel")
	s.Ctx.AddIntents(gateway.IntentGuildVoiceStates)
	s.Ctx.AddIntents(gateway.IntentGuildMessages)
}

// NewSpeaker generates a new Speaker instance
func NewSpeaker() *Speaker {
	return &Speaker{}
}

// Join allows the bot to join a voice channel from an user's request
func (s *Speaker) Join(m *gateway.MessageCreateEvent) error {
	v, err := voice.NewSession(s.Ctx.State)
	if err != nil {
		log.Fatal("Failed to create a new voice session:", err)
	}

	// Fetch voice state from store
	err = v.JoinChannel(m.GuildID, m.ChannelID, false, false)
	if err != nil {
		log.Error("Failed to join channel:", err)
		return fmt.Errorf("**failed to join channel**")
	}

	// Indicate speaking
	if err := v.Speaking(voicegateway.Microphone); err != nil {
		log.Errorf("failed to indicate speaking, %s", err.Error())
		return nil
	}

	// Send success message
	_, err = s.Ctx.SendMessage(m.ChannelID, fmt.Sprintf("**Joined**: %s", m.ChannelID.Mention()), nil)
	if err != nil {
		log.Errorf("Failed to send joined channel message")
		return nil
	}

	log.Infof("Joined voice channel: %s", m.ChannelID.String())

	return nil
}

// Leave allows the bot to leave a voice channel from an user's request
func (s *Speaker) Leave(m *gateway.MessageCreateEvent) error {
	if !m.GuildID.IsValid() {
		return fmt.Errorf("invalid guild ID")
	}
	s.Voice.ErrorLog = func(err error) {
		log.Error(err)
	}

	return s.Voice.Leave()
}

// Queue up and play a track in the queue
func (s *Speaker) Play(m *gateway.MessageCreateEvent, args bot.ArgumentParts) {
	panic("implement me")
}
