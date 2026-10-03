package extractpdf

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/gen2brain/go-fitz"
	"github.com/tmc/langchaingo/textsplitter"
)

type DataExtract struct{
	Server *models.Server
}

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

func(e *DataExtract)  ExtractData(file *multipart.File, filePath string) (int, error) {
	// initialize the fitz document
	document, err := fitz.New(filePath)
	if err != nil{
		log.Println("[EXTRACT DATA]: Error in creating fitz document.")
		return http.StatusInternalServerError, err
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
			return http.StatusInternalServerError, err
		}
		data.WriteString(string(text))
	}
	result := data.String()
	result = strings.TrimSpace(result)
	result = strings.TrimSuffix(result,"\n\n")
	result = strings.Trim(result, "\n\n")
	result = strings.Trim(result, "\t")



	// convert them to splitted text and create embeddings of them
	splittedDocx, err  := splitDocument(data.String())
	if err != nil{
		log.Println("[EXTRACT DATA]: Error in splitting the docx.")
		return http.StatusInternalServerError, errors.New("Error in splitting docx.")
	}

	// create embeddings and store it in the vector store
	err = VectorStore(splittedDocx, e.Server, filePath)
	if err != nil{
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}