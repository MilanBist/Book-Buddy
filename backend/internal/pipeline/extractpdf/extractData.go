package extractpdf

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/gen2brain/go-fitz"
	"github.com/qdrant/go-client/qdrant"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/textsplitter"
)



type DataExtract struct{
	Server *models.Server
}


func createCollection( s *models.Server, size int64) error{
	client := s.Store
	collectionName := "AI_Book_summarizer"
	exists, err := client.CollectionExists(context.Background(), collectionName)
	if err != nil{
		fmt.Println("Collection doesn't exist so creating a new one.")
	}


	if !exists{
		err = client.CreateCollection(context.Background(), &qdrant.CreateCollection{
			CollectionName: "AI_Book_summarizer",
			VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size: uint64(size),
			Distance: qdrant.Distance_Cosine,
			}),
		})

		if err != nil{
			fmt.Println("Error in creating the collection.")
			return err
		}

		_,err = client.CreateFieldIndex(
			context.Background(), 
			&qdrant.CreateFieldIndexCollection{
				CollectionName: "AI_Book_summarizer",
				FieldName: "Title",
				FieldType: qdrant.FieldType_FieldTypeBool.Enum(),
			},
		)
	}
	return nil
}

func splitDocument(document string)([]string, error){
	splitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(1000),
		textsplitter.WithChunkOverlap(100),
	)
	chunks, err := splitter.SplitText(document)
	if err != nil{
		return nil, err
	}

	return chunks, nil
}

func(e *DataExtract)  ExtractData(file *multipart.File, filePath string) (int, error) {
	// initialize the fitz document
	llm, err := ollama.New(
		ollama.WithModel(e.Server.Config.OllamaEmbeddingModel),
	)

	if err != nil{
		fmt.Println("Error in splitting docx.")
		return http.StatusInternalServerError,err
	}
	document, err := fitz.New(filePath)
	if err != nil{
		log.Println("[EXTRACT DATA]: Error in creating fitz document.")
		return http.StatusInternalServerError, err
	}
	defer document.Close()
	var data strings.Builder

	// extract the text from the given document and store using the batching
	var chunkCount int = 0
	var batch int = 0

	// create the collection

	start := time.Now()
	resultChan := make(chan error)


	err = createCollection(e.Server, 1024)

	// var wg sync.WaitGroup
	for n:=0; n<document.NumPage(); n++{
		var text string
		text, err = document.Text(n)
		if err != nil{
			log.Println("Error in getting text.")
			return http.StatusInternalServerError, err
		}
		chunkCount += 1
		data.WriteString(string(text))
		// check for the count of the chunk if its 10 just store to vector db

		if chunkCount >= 4 || n + 1 == document.NumPage(){
			fmt.Println("For batch ",batch+1)
			// split the data 
			splittedData, err := splitDocument(data.String())
			if err != nil{
				fmt.Println("Error in splitting")
				return http.StatusInternalServerError, err
			}

			// store the data from this section
			
			batch += 1 
			go VectorStore(splittedData, e.Server, filePath,llm, resultChan)
			// let data data be 0 and count be also 0
			chunkCount = 0
			data.Reset()
		}
	}

	var resultedError error
	for i := range batch{
		fmt.Println(i+1)
		resultedError = <- resultChan
		if resultedError != nil{
			return http.StatusInternalServerError, resultedError
		}
	}
	fmt.Println("Time for extracting splitting and adding to vector db: ", time.Since(start))
	return http.StatusAccepted, nil
}