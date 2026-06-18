package askllm

import (
	"strings"
)

func GeneratePromptAugmentation(docx []string, query string)(string, error){

	// prompt := `You are a helpful assistant, who is going to guide a person on the question asked by him/her on basis of the 
	// context provided below.

	// Question you need to answer on the basis of the above context:
	// `+query+`

	// Use the context below. But answer like you are saying the answer like an expert of this domain.
	// Don't mention like according to the context or such say like in my view insted or other.
	// If possible decorate the answer like in points or bullets.
	
	// Context:
	// `+strings.Join(docx, "\n") +`

	// If the question is out of the context return the answer if possible or else return with the statement of "Out of context."

	// Return the answer in most simple way possible and must cover all technical terms mentioned in the context.
	// `

	prompt := `
		You are an expert assistant helping users understand information based strictly on the provided context.

		----------------------------
		RULES:
		----------------------------
		1. Use ONLY the provided context to answer.
		2. If the answer is not found in the context, reply exactly:
		"Out of context."
		3. Do NOT mention "context" or "documents" in your answer.
		4. Do NOT say "according to the context".
		5. If possible, explain in simple and clear language.
		6. Structure the answer in bullet points when helpful.
		7. Keep the explanation accurate and easy to understand.
		8. Preserve and explain all important technical terms from the context.

		----------------------------
		CONTEXT:
		----------------------------
		` + strings.Join(docx, "\n\n") + `

		----------------------------
		QUESTION:
		----------------------------
		` + query + `

		----------------------------
		ANSWER:
	`
	return prompt, nil
}