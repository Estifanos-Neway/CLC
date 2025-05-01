package clc

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/estifanos-neway/CLC/config"
	"github.com/estifanos-neway/CLC/internal/api/gemini"
)

func (c *CLC) GetResponse() error {
	dirContents, err := filepath.Glob("." + "/*")
	if err != nil {
		return err
	}
	msgContent := request{
		Prompt: c.Prompt,
		Context: Context{
			HostMachine:       runtime.GOOS, // TODO Add better host machine details
			CurrentDirContent: dirContents,
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

func (c *CLC) Go(keep bool, skip bool) error {
	slog.Debug("Creating script file...")
	scriptName := fmt.Sprintf("./clc-script%s", c.Response.Commands.FileExtension)
	f, err := os.Create(scriptName)
	if err != nil {
		slog.Error("Error creating script file", "Error", err)
		return err
	}
	defer f.Close()
	slog.Debug("Script file created.")

	slog.Debug("Writing script file...")
	if _, err := f.WriteString(c.Response.Commands.ScriptContent); err != nil {
		slog.Error("Error writing script file", "Error", err)
		return err
	}
	slog.Debug("Wrote script file.")

	if !skip {
		out, err := exec.Command(c.Response.Commands.Runner, scriptName).Output()
		if err != nil {
			slog.Error("Error running exec command", "Error", err)
			return err
		}
		slog.Debug("Script file executed.")
		fmt.Println(string(out))
	}
	slog.Debug("Script file executed.")

	if !keep {
		slog.Debug("Removing script file...")
		if err := os.Remove(f.Name()); err != nil {
			slog.Error("Error removing script file", "Error", err)
			return err
		}
		slog.Debug("Script file removed.")
	} else {
		slog.Debug("Script file kept.")
	}
	return nil
}

func (r *Response) String() string {
	// TODO Implement better formatted string
	text, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		panic(err)
	}
	return string(text)
}
