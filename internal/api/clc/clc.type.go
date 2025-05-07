package clc

import "github.com/estifanos-neway/CLC/internal/api/gemini"

type CLC struct {
	Gemini   *gemini.Gemini
	Prompt   string
	Response *Response
}

type request struct {
	Prompt  string  `json:"prompt"`
	Context Context `json:"context"`
}

type Context struct {
	HostMachine       string   `json:"hostMachine"`
	CurrentDirContent []string `json:"currentDirContent"`
}

type Response struct {
	Status   int       `json:"status"` // 1|0
	Reason   Reason    `json:"reason"`
	Message  string    `json:"message"`
	Commands *Commands `json:"commands"`
}

type Commands struct {
	ScriptContent string `json:"scriptContent"`
	FileExtension string `json:"fileExtension"`
	Runner string `json:"runner"`
}

type Reason string
