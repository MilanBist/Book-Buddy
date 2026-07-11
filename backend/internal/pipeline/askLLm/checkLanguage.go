package askllm

import (
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)



func CheckLanguage(query string, h *models.Server) (string, error){
	prompt := `
			You are an expert language recognizer. Based on the given text provided tell me which language is it.

			------------------------------------
			ANSWER RETURNING RULES
			------------------------------------
			1. Return the answer in form of the string.
			2. Just return the language name starting with the capital letter like: English.


			------------------------------------
			EXAMPLE
			------------------------------------
			query : Hello, what is new today.
			Your response : English.

			------------------------------------
			Now Based on the text below return me the language it is in.
			--------------------------------------------------------
			`+query

	// generate the result based on the given prompt
	answer, err := GenerateResult(prompt, h)
	if err != nil{
		return "", err
	}

	return answer, nil
	
}