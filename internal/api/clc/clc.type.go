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
	ToolCheck []*ToolCheck `json:"toolCheck"`
	Cmds      []*Cmd       `json:"cmds"`
}

type ToolCheck struct {
	Tool        string   `json:"tool"`
	Description string   `json:"description"`
	Cmd         string   `json:"cmd"`
	Args        []string `json:"args"`
	OkIf        *OkIf    `json:"okIf"`
	OnFail      *OnFail  `json:"onFail"`
}

type OkIf struct {
	OkExistCode bool   `json:"okExistCode"` // true|false
	RegExp      string `json:"regExp"`
}

type OnFail struct {
	ToolNotFoundMessage  string `json:"toolNotFoundMessage"`
	InstructionToInstall string `json:"instructionToInstall"`
}

type Cmd struct {
	Cmd         string   `json:"cmd"`
	Args        []string `json:"args"`
	Description string   `json:"description"`
}

type Reason string
