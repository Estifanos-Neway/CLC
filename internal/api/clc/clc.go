package CLC

import (
	"encoding/json"
	"runtime"

	"github.com/estifanos-neway/CLC/config"
	"github.com/estifanos-neway/CLC/internal/api/gemini"
)

func (c *CLC) GetResponse() error {
	msgContent := request{
		HostMachine: runtime.GOOS,
		Prompt:      c.Prompt,
	}

	msgContentText, err := json.Marshal(msgContent)
	if err != nil {
		return nil
	}

	geminiGeneralConfig := gemini.GenerationConfig(config.AppConfig.Gemini.GenerationConfig)
	chat := &gemini.Chat{
		Gemini:            c.Gemini,
		SystemInstruction: gemini.CreateContent(gemini.User, ""),
		GenerationConfig:  &geminiGeneralConfig,
	}

	msg := gemini.Message{
		Chat: chat,
		Contents: []*gemini.Content{
			gemini.CreateContent(gemini.User, string(msgContentText)),
		},
	}

	res, err := msg.Send()
	if err != nil {
		return nil
	}

	// TODO Check if there actually is a response text first.
	resText := res.Candidates[0].Content.Parts[0].Text
	if err := json.Unmarshal([]byte(resText), c.Response); err != nil {
		return nil
	}

	return nil
}
