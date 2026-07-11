// generate the response of user in required language
package askllm

import (
	"errors"
	"net/http"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

func GenerateInRequiredLanguage(englishResponse string, language string, h *models.Server, w http.ResponseWriter, r *http.Request)(error){
	prompt := `
		You are an expert multilingual translator.
		Your task is to translate the following text into the target language.

		Requirements:
		- Translate ONLY into the target language.
		- Preserve the original meaning exactly.
		- Preserve the tone, emotion, and intent of the original text.
		- Do NOT add, remove, summarize, or explain anything.
		- Do NOT change the motive or message.
		- Keep the same level of formality (formal/informal).
		- Preserve any technical terms, names, code snippets, URLs, or numbers unless they should naturally be translated.
		- If the original contains markdown, bullet points, or formatting, preserve them exactly.
		- The translation should read naturally to a native speaker while conveying the exact same feeling as the original.
		- Never mention the note of your own just respond with the response only in given language.

		Target Language:
		`+language+`
		---------------------------------------
		FORMATTING & STYLE GUIDELINES
		---------------------------------------
		1. Use Markdown formatting:
		- # for main headings
		- ## for subheadings
		2. Use emojis where they add clarity or friendliness 🚀📌💡
		3. Use:
		i. ii. iii. for ordered points
		- for unordered points
		4. Separate major sections clearly, for example:


		Text:
		`+englishResponse+`
	`


	// generate the result in required result
	err := GenerateResultAndSendToFrontend(prompt, h, w, r)
	if err != nil{
		return errors.New("Error in generating the prompt output.")
	}

	return nil
}