package extractpdf

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/gen2brain/go-fitz"
	"github.com/tmc/langchaingo/textsplitter"
)

func splitDocument(document string)([]string, error){
	splitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(1000),
		textsplitter.WithChunkOverlap(150),
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

	// convert them to splitted text and create embeddings of them
	splittedDocx, err  := splitDocument(result)
	if err != nil{
		return errors.New("Error in splitting docx.")
	}


	// create embeddings and store it in the vector store
	// fmt.Println(len(splittedDocx))
	// fmt.Println(splittedDocx)


	// store in the vector store
	err = vectorStore(splittedDocx, s)


	return nil
}