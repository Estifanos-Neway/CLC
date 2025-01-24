package gemini

const (
	User  Role = "user"
	Model Role = "model"
)

const (
	SafetyRatingCategoryHarassment     SafetyRatingCategory = "HARM_CATEGORY_HARASSMENT"
	SafetyRatingCategoryHeatSpeech     SafetyRatingCategory = "HARM_CATEGORY_HATE_SPEECH"
	SafetyRatingCategorySexualExplicit SafetyRatingCategory = "HARM_CATEGORY_SEXUALLY_EXPLICIT"
	SafetyRatingCategoryDangerous      SafetyRatingCategory = "HARM_CATEGORY_DANGEROUS_CONTENT"
	SafetyRatingCategoryCivilIntegrity SafetyRatingCategory = "HARM_CATEGORY_CIVIC_INTEGRITY"
)

const (
	SafetyRatingProbabilityHigh       SafetyRatingProbability = "HIGH"
	SafetyRatingProbabilityMedium     SafetyRatingProbability = "MEDIUM"
	SafetyRatingProbabilityLow        SafetyRatingProbability = "LOW"
	SafetyRatingProbabilityNegligible SafetyRatingProbability = "NEGLIGIBLE"
)
