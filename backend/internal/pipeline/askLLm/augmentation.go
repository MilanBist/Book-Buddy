package askllm

import (
	"strings"
)

func GeneratePromptAugmentation(docx []string, query string)(string, error){

	prompt := `You are a helpful assistant, who is going to guide a person on the question asked by him/her on basis of the 
	context provided below.

	Question you need to answer on the basis of the above context:
	`+query+`

	Use the context below. But answer like you are saying the answer like an expert of this domain.
	Don't mention like according to the context or such say like in my view insted or other.
	If possible decorate the answer like in points or bullets.
	
	Context:
	`+strings.Join(docx, "\n") +`

	If the question is out of the context return the answer if possible or else return with the statement of "Out of context."

	Return the answer in most simple way possible and must cover all technical terms mentioned in the context.
	`
	return prompt, nil
}