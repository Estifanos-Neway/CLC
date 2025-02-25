package clc

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"text/tabwriter"

	"github.com/estifanos-neway/CLC/config"
	"github.com/estifanos-neway/CLC/internal/api/gemini"
	"github.com/estifanos-neway/CLC/internal/pkg/dir"
)

func (c *CLC) GetResponse() error {
	dirContents, err := dir.GetDirectoryContents(".")
	if err != nil {
		return err
	}
	msgContent := request{
		Prompt: c.Prompt,
		Context: Context{
			HostMachine:       runtime.GOOS, // TODO Add better host machine details
			CurrentDirContent: *dirContents,
		},
	}

	msgContentText, err := json.Marshal(msgContent)
	if err != nil {
		return err
	}

	geminiGeneralConfig := gemini.GenerationConfig(config.AppConfig.Gemini.GenerationConfig)
	chat := &gemini.Chat{
		Gemini:            c.Gemini,
		SystemInstruction: gemini.CreateContent(gemini.User, config.AppConfig.Gemini.SystemInstruction),
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
		return err
	}

	if len(res.Candidates) == 0 || len(res.Candidates[0].Content.Parts) == 0 {
		resString, err := json.Marshal(res)
		if err != nil {
			return err
		}
		return fmt.Errorf("invalid response from the ai api %s", string(resString))
	}

	resText := res.Candidates[0].Content.Parts[0].Text
	c.Response = &Response{}
	if err := json.Unmarshal([]byte(resText), c.Response); err != nil {
		return err
	}

	return nil
}

func (c *CLC) Print(out io.Writer) error {
	w := tabwriter.NewWriter(out, 5, 1, 3, ' ', 0)
	fmt.Fprintln(w, "\033[5;32m1. Check for required tools 🧪\033[0m")
	fmt.Fprint(w, "| Tool\t| Command\t| Description\t| If Not Found\t|\t\n")
	for _, t := range c.Response.Commands.ToolCheck {
		fmt.Fprintf(w, "| %s\t| %s\t| %s\t| %s\t|\t\n", t.Tool, t.Cmd, t.Description, t.OnFail.InstructionToInstall)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "\033[5;32m2. Execute these commands 🚀\033[0m")
	fmt.Fprint(w, "| Command\t| Description\t|\t\n")
	for _, cmd := range c.Response.Commands.Cmds {
		fmt.Fprintf(w, "| %s\t| %s \t|\t\n", cmd.Cmd, cmd.Description)
	}

	return w.Flush()
}

func (c *CLC) Exec() error {
	fmt.Println("Exec: not implemented!")
	return nil
}
