package gemini

type Gemini struct {
	Url    string `json:"-"`
	ApiKey string `json:"-"`
}

type Chat struct {
	*Gemini
	SystemInstruction *Content          `json:"systemInstruction"`
	GenerationConfig  *GenerationConfig `json:"generationConfig"`
}

type Message struct {
	*Chat
	Contents []*Content `json:"contents"`
}

type GenerationConfig struct {
	Temperature      float32 `json:"temperature"`
	TopK             int     `json:"topK"`
	TopP             float32 `json:"topP"`
	MaxOutputTokens  int     `json:"maxOutputTokens"`
	ResponseMimeType string  `json:"responseMimeType"`
}

type Content struct {
	Role  Role    `json:"role"`
	Parts []*Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type Response struct {
	Candidates    []*Candidate   `json:"candidates"`
	UsageMetadata *UsageMetadata `json:"usageMetadata"`
	ModelVersion  string         `json:"modelVersion"`
}

type Candidate struct {
	Content       *Content        `json:"content"`
	FinishReason  string          `json:"finishReason"`
	SafetyRatings []*SafetyRating `json:"safetyRatings"`
	AvgLogprobs   float64         `json:"avgLogprobs"`
}

type UsageMetadata struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}

type SafetyRating struct {
	Category    SafetyRatingCategory    `json:"category"`
	Probability SafetyRatingProbability `json:"probability"`
}

type Role string
type SafetyRatingCategory string
type SafetyRatingProbability string
