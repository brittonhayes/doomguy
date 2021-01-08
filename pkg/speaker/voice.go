package speaker

import (
	"fmt"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/arikawa/v2/state"
	"github.com/diamondburned/arikawa/v2/voice"
	"github.com/diamondburned/arikawa/v2/voice/voicegateway"
	log "github.com/sirupsen/logrus"
)

type Speaker struct {
	Ctx   *bot.Context
	State *state.State
	Voice *voice.Session
}

// Setup demonstrates the CanSetup interface. This function will never be parsed
// as a callback of any event.
func (s *Speaker) Setup(sub *bot.Subcommand) {
	sub.Command = "speaker"
	sub.Description = "Join voice channels, play music, or youtube videos"
	sub.ChangeCommandInfo("Join", "join", "Join the voice channel")
	sub.ChangeCommandInfo("Leave", "leave", "Leave the voice channel")
}

// NewSpeaker generates a new Speaker instance
func NewSpeaker(state *state.State) *Speaker {
	if err := state.Open(); err != nil {
		log.Fatalln("failed to open gateway:", err)
	}

	v, err := voice.NewSession(state)
	if err != nil {
		log.Fatalln("failed to create voice session:", err)
	}

	return &Speaker{
		State: state,
		Voice: v,
	}
}

// Join allows the bot to join a voice channel from an user's request
func (s *Speaker) Join(m *gateway.MessageCreateEvent) error {
	vs, err := s.State.VoiceState(m.GuildID, m.Author.ID)
	if err != nil {
		log.Error("Failed to get voice state:", err)
		return nil
	}

	if !vs.ChannelID.IsValid() {
		log.Errorf("voice state channel is invalid, %#v", vs)
		return fmt.Errorf("**user is not in a channel**")
	}

	c, err := s.State.Channel(vs.ChannelID)
	if err != nil {
		log.Fatalln("failed to get channel:", err)
	}

	// Fetch voice state from store
	err = s.Voice.JoinChannel(m.GuildID, c.ID, false, false)
	if err != nil {
		//log.Errorf("MESSAGE: %#v\nSPEAKER: %#v", m, s)
		//panic(err)
		log.Error("Failed to join channel:", err)
		return fmt.Errorf("**failed to join channel**")
	}

	// Indicate speaking
	if err := s.Voice.Speaking(voicegateway.Microphone); err != nil {
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
