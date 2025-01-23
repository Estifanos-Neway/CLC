package gemini

type Gemini struct {
	Url    string
	ApiKey string
}

type Chat struct {
	Gemini
	SystemInstruction Content          `json:"systemInstruction"`
	GenerationConfig  GenerationConfig `json:"generationConfig"`
}

type Message struct {
	Chat
	Contents []Content `json:"contents"`
}

type GenerationConfig struct {
	Temperature      float32 `json:"temperature"`
	TopK             int     `json:"topK"`
	TopP             float32 `json:"topP"`
	MaxOutputTokens  int     `json:"maxOutputTokens"`
	ResponseMimeType string  `json:"responseMimeType"`
}

type Content struct {
	Role  Role   `json:"role"`
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type Role string

const (
	User  Role = "user"
	Model Role = "model"
)
