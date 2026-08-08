package askllm

import (
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)



func CheckLanguage(query string, h *models.Server) (string, error){
prompt := `
		----------------------------------------
		SCRIPT AND ROMANIZED LANGUAGE IDENTIFICATION
		----------------------------------------
		1. Identify the actual language, regardless of the writing system used.
		2. The text may be written in:
		- Its native script (e.g., Devanagari, Arabic, Cyrillic, Chinese characters, etc.)
		- A romanized (transliterated) form using the Latin/English alphabet.

		3. Detect both native-script and romanized versions of every language.

		4. Never classify a text as English solely because it is written using the English (Latin) alphabet.

		5. For romanized text, determine the language using:
		- vocabulary
		- grammar
		- sentence structure
		- common expressions
		- language-specific particles and word endings

		6. For native-script text, use both the script and the linguistic characteristics to determine the language.

		7. If two or more languages share the same script (for example Nepali and Hindi both use Devanagari), distinguish them using vocabulary and grammar rather than script alone.

		8. If two or more languages are romanized using the Latin alphabet (for example Nepali and Hindi), distinguish them using vocabulary, grammar, and common expressions rather than the alphabet.


		----------------------------------------
		EXAMPLES
		----------------------------------------

		English

		Text:
		How are you today?
		Return:
		English

		--------------------------------------------------

		Mandarin Chinese

		Text:
		你好，你今天怎么样？
		Return:
		Mandarin Chinese

		Text:
		Ni hao, ni jintian zenmeyang?
		Return:
		Mandarin Chinese

		--------------------------------------------------

		Hindi

		Text:
		आपका नाम क्या है?
		Return:
		Hindi

		Text:
		Aapka naam kya hai?
		Return:
		Hindi

		--------------------------------------------------

		Spanish

		Text:
		¿Cómo estás hoy?
		Return:
		Spanish

		Text:
		Como estas hoy?
		Return:
		Spanish

		--------------------------------------------------

		French

		Text:
		Comment allez-vous aujourd'hui ?
		Return:
		French

		Text:
		Comment allez vous aujourd hui?
		Return:
		French

		--------------------------------------------------

		Arabic

		Text:
		كيف حالك اليوم؟
		Return:
		Arabic

		Text:
		Kayfa haluka alyawm?
		Return:
		Arabic

		--------------------------------------------------

		Bengali

		Text:
		আপনার নাম কী?
		Return:
		Bengali

		Text:
		Apnar nam ki?
		Return:
		Bengali

		--------------------------------------------------

		Portuguese

		Text:
		Como você está hoje?
		Return:
		Portuguese

		Text:
		Como voce esta hoje?
		Return:
		Portuguese

		--------------------------------------------------

		Russian

		Text:
		Как дела?
		Return:
		Russian

		Text:
		Kak dela?
		Return:
		Russian

		--------------------------------------------------

		Urdu

		Text:
		آپ کا نام کیا ہے؟
		Return:
		Urdu

		Text:
		Aap ka naam kya hai?
		Return:
		Urdu

		--------------------------------------------------

		Nepali (Important: Differentiate from Hindi)

		Text:
		तपाईंको नाम के हो?
		Return:
		Nepali

		Text:
		Tapai ko naam ke ho?
		Return:
		Nepali

		Text:
		Mero ghar Kathmandu ma cha.
		Return:
		Nepali

		--------------------------------------------------

		IMPORTANT DISTINCTIONS

		Nepali:
		- Tapai
		- Timi
		- Hajur
		- Mero
		- Hamro
		- Cha
		- Chha
		- Ho
		- Hoina
		- Garnu
		- Bhayo

		Hindi:
		- Aap
		- Tum
		- Mera
		- Hamara
		- Hai
		- Hain
		- Kya
		- Nahi
		- Karna
		- Ho gaya

		Urdu:
		- آپ
		- میرا
		- کیا
		- ہے

		Bengali:
		- Apnar
		- Amar
		- Ki
		- Achhe

		Chinese (Romanized):
		- Ni hao
		- Xiexie
		- Wo
		- Shi
		- Ma

		Arabic (Romanized):
		- Marhaba
		- Kayfa
		- Ana
		- Shukran

		Russian (Romanized):
		- Privet
		- Kak dela
		- Spasibo

		French:
		- Bonjour
		- Merci
		- Comment allez-vous

		Spanish:
		- Hola
		- Gracias
		- Cómo estás

		Portuguese:
		- Olá
		- Obrigado
		- Como você está


		Also implement such for other languages as well.

		------------------------------------------------

		ANSWER RETURNING FORMAT
		-----------------------------------------------
		1. Just return the name of the language don't provide the explaination.
		For example: 
		Text: Mero naam milan ho.
		Your response just: Nepali
		2. Follow this format specifically.

		Here given below is the text now Identify it.

		----------------------------------------
		TEXT
		----------------------------------------
		` + query


	// generate the result based on the given prompt
	answer, err := GenerateLanguageResults(prompt, h)
	if err != nil{
		return "", err
	}

	return answer, nil
	
}