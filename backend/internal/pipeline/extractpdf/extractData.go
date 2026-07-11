package extractpdf

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	askllm "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/askLLm"
	"github.com/gen2brain/go-fitz"
	"github.com/tmc/langchaingo/textsplitter"
)

func splitDocument(document string)([]string, error){
	splitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(1000),
		textsplitter.WithChunkOverlap(200),
	)
	chunks, err := splitter.SplitText(document)
	if err != nil{
		return nil, err
	}

	return chunks, nil
}

func ExtractData(file *multipart.File, filePath string, s *models.Server) error {

	// initialize the fitz document
	document, err := fitz.New(filePath)
	if err != nil{
		log.Println("Error in creating fitz document.")
		return err
	}
	defer document.Close()

	var data strings.Builder

	// extract the text from the given document
	for n:=0; n<document.NumPage(); n++{
		// each time add the data as per text
		var text string
		text, err = document.Text(n)
		if err != nil{
			log.Println("Error in getting text.")
			return err
		}
		data.WriteString(string(text))
	}
	result := data.String()
	result = strings.TrimSpace(result)
	result = strings.TrimSuffix(result,"\n\n")
	result = strings.Trim(result, "\n\n")
	result = strings.Trim(result, "\t")


	// convert whole of the result in string remove unnecessay numbers and other things
	response, err := askllm.ConvertToEnglish(result, s)
	if err != nil{
		log.Println("Error in generating the response of the docx.")
		return err
	}

	// convert them to splitted text and create embeddings of them
	splittedDocx, err  := splitDocument(response)
	if err != nil{
		return errors.New("Error in splitting docx.")
	}

	// create embeddings and store it in the vector store
	err = VectorStore(splittedDocx, s, filePath)

	return nil
}