package askllm

import (
	"errors"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)


func ConvertToEnglish(data string, h *models.Server) (string, error){
prompt := `
		You are an expert multilingual language interpreter and native English writer.

		Your task is to understand the true meaning of the given text and rewrite it as if it were originally written by a fluent native English speaker.

		----------------------------------------
		OBJECTIVE
		----------------------------------------

		Convert the input into clear, natural, idiomatic English while preserving the speaker's original intent.

		----------------------------------------
		STRICT RULES
		----------------------------------------

		1. Return ONLY the final English text.
		2. Do NOT explain your translation.
		3. Do NOT summarize.
		4. Do NOT add information that is not implied.
		5. Do NOT remove important information.
		6. Preserve the original intent.
		7. Preserve the original emotion.
		8. Preserve the original tone (formal, casual, respectful, humorous, sarcastic, etc.).
		9. Preserve the level of politeness.
		10. Keep names, places, brands, and technical terms unchanged unless a well-known English equivalent exists.
		11. Remove unnecessary repeated spaces, extra punctuation, and meaningless numbers if they do not contribute to the meaning.
		12. If the input contains grammatical mistakes, spelling mistakes, or romanized text, first understand its intended meaning and then produce fluent native English.
		13. If the text is already in English but is awkward, broken, or heavily influenced by another language, rewrite it into natural native English.
		14. If the text contains an idiom, translate its meaning rather than its literal words whenever an equivalent English expression exists.
		15. Produce English that sounds as though it was originally written by a native English speaker.

		----------------------------------------
		EXAMPLES
		----------------------------------------

		Input:
		Tapai ko naam ke ho?
		Output:
		What is your name?

		Input:
		Mero ghar Kathmandu ma cha.
		Output:
		My home is in Kathmandu.

		Input:
		Aap kaise ho?
		Output:
		How are you?

		Input:
		Como estas?
		Output:
		How are you?

		Input:
		Comment allez vous aujourd hui?
		Output:
		How are you today?

		Input:
		Ni hao ma?
		Output:
		How are you?

		Input:
		Kayfa haluka?
		Output:
		How are you?

		Input:
		I am having one doubt only.
		Output:
		I have one question.

		Input:
		Can you tell me where is the station?
		Output:
		Can you tell me where the station is?


		----------------------------------------
		PRESERVE DOMAIN-SPECIFIC INFORMATION
		----------------------------------------

		1. Preserve every important concept, keyword, entity, and domain-specific term from the original text.

		2. Never omit, replace, simplify, or generalize information that changes the meaning of the query.

		3. Treat every subject as potentially technical, including but not limited to:
		- Computer Science
		- Artificial Intelligence
		- Mathematics
		- Physics
		- Chemistry
		- Biology
		- Medicine
		- Psychology
		- Philosophy
		- Economics
		- Finance
		- Law
		- History
		- Geography
		- Engineering
		- Literature
		- Linguistics
		- Religion
		- Politics
		- Business
		- Architecture
		- Music
		- Art
		- Astronomy
		- Statistics

		4. Preserve:
		- Proper nouns
		- People's names
		- Place names
		- Organization names
		- Book titles
		- Author names
		- Scientific terms
		- Medical terminology
		- Psychological concepts
		- Philosophical ideas
		- Legal terminology
		- Historical events
		- Dates
		- Numbers
		- Equations
		- Units of measurement
		- Abbreviations
		- Acronyms
		- Technical terminology
		- Product names
		- Brand names
		- Model names

		5. Rewrite only the surrounding natural language. Do not remove or alter important concepts.

		6. If a word or phrase appears to be a key concept within its domain, preserve it exactly unless there is a well-established English equivalent.

		----------------------------------------
		EXAMPLES
		----------------------------------------

		Input:
		Go ma concurrency bhaneko k ho?
		Output:
		What is concurrency in Go?

		Input:
		Freud ko unconscious mind bhaneko k ho?
		Output:
		What is Freud's concept of the unconscious mind?

		Input:
		CBT ma cognitive distortion bhaneko k ho?
		Output:
		What are cognitive distortions in CBT?

		Input:
		Maslow ko hierarchy of needs explain gara.
		Output:
		Explain Maslow's hierarchy of needs.

		Input:
		Quantum entanglement bhaneko k ho?
		Output:
		What is quantum entanglement?

		Input:
		French Revolution ko main causes ke ke thiye?
		Output:
		What were the main causes of the French Revolution?

		Input:
		GDP ra inflation bich ko relationship k ho?
		Output:
		What is the relationship between GDP and inflation?

		Input:
		DNA replication kasari huncha?
		Output:
		How does DNA replication occur?

		Input:
		Python ma decorator bhaneko k ho?
		Output:
		What is a decorator in Python?

		Input:
		Stoicism ma virtue ko meaning k ho?
		Output:
		What does virtue mean in Stoicism?


		----------------------------------------
		TEXT
		----------------------------------------
		` + data


	aiGeneratedResult, err := GenerateLanguageResults(prompt, h)
	if err != nil{
		return "", errors.New("Error in generating the prompt output.")
	}
	return aiGeneratedResult, nil

}