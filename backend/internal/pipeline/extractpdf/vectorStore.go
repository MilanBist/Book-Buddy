package extractpdf

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/google/uuid"
	"github.com/qdrant/go-client/qdrant"
	"github.com/tmc/langchaingo/llms/ollama"
)

// creating the embeddings
func CreateEmbeddings(splittedDocx []string, llm *ollama.LLM) ([][]float32, error) {
	startTime := time.Now()
	response, err := llm.CreateEmbedding(context.Background(), 
		splittedDocx,
	)
	if err != nil{
		fmt.Println("[CREATE EMBEDDINGS] : Error is: ", err)
		fmt.Println("Error in creating embeddings.")
		return nil,err
	}
	fmt.Println("For the time to generate embedding. : ", time.Since(startTime))
	return response,nil
}



func createCollectionWithName(splittedDocx []string, embededBook [][]float32, s *models.Server, bookTitle string) error{
	client := s.Store
	var points []*qdrant.PointStruct
	for i:=0; i<len(embededBook); i++{
		point := &qdrant.PointStruct{
			Id:       qdrant.NewIDUUID(uuid.NewString()),
			Vectors:  qdrant.NewVectors(embededBook[i]...),
			Payload:  qdrant.NewValueMap(map[string]any{
				"text": splittedDocx[i],
				"chunkId": strconv.Itoa(i+1),
				"bookTitle": bookTitle,
			}),
		}
		points = append(points, point)
	}


	// insert the data in the certain database
	startTime := time.Now()
	_, err := client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: "AI_Book_summarizer",
		Points: points,
	})
	if err != nil{
		fmt.Println("Error in inserting the data.")
		return err
	}
	fmt.Println("Inserting time into the vector db: ", time.Since(startTime))
	return nil
}

// create the embeddings and store them to the vector storage.
func VectorStore(splittedDocx []string, s *models.Server, fileName string,llm *ollama.LLM, c chan error){

	path := strings.Split(fileName, "/")
	bookTitle := path[len(path)-1]
	book := strings.Split(bookTitle, ".")
	bookTitle = book[0]

	startTime := time.Now()
	response, err := CreateEmbeddings(splittedDocx, llm)
	if err != nil{
		c <- err
		return
	}
	fmt.Println("Creating embedding time: ", time.Since(startTime))
	err = createCollectionWithName(splittedDocx, response, s, bookTitle)
	if err != nil{
		c <- err
		return
	}	
	c <- nil
}