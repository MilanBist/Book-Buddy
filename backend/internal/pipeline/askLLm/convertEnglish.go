package askllm

import (
	"errors"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)


func ConvertToEnglish(data string, h *models.Server) (string, error){
	prompt := `
		You are an expert multilingual translator specializing in emotionally faithful translations.

		Your task is to translate the given text into natural, fluent English while preserving:

		1. The original meaning.
		2. The speaker's emotion and tone.
		3. The writing style (formal, casual, poetic, humorous, sarcastic, respectful, etc.).
		4. Cultural nuances whenever possible.
		5. The level of politeness and honorifics.
		6. Any implied meaning, not just the literal words.

		Translation Guidelines:
		- Do NOT summarize.
		- Do NOT explain the translation.
		- Do NOT add information.
		- Do NOT remove information.
		- Keep the same emotional intensity.
		- If an idiom exists in English with the same meaning and emotional impact, use it instead of translating literally.
		- If no equivalent exists, translate naturally while preserving the intended feeling.
		- Preserve punctuation and emphasis where appropriate.
		- Keep names, places, and technical terms unchanged unless there is a well-known English equivalent.
		-If unnecessary numbers or spaces are there just remove them in the output.
		-If possible keep the data in same length.

		Output only the translated English text.

		Text:
		`+data+`
	`

	aiGeneratedResult, err := GenerateResult(prompt, h)
	if err != nil{
		return "", errors.New("Error in generating the prompt output.")
	}

	return aiGeneratedResult, nil

}