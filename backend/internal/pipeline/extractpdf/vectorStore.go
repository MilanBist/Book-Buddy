package extractpdf

import (
	"context"
	"fmt"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/qdrant/go-client/qdrant"
	"github.com/tmc/langchaingo/llms/ollama"
)

func createEmbeddings(splittedDocx []string, s *models.Server) ([][]float32, error) {
	// create the ollama client
	llm, err := ollama.New(
		ollama.WithModel(s.Config.OllamaModel),
	)

	if err != nil{
		fmt.Println("Error in splitting docx.")
		return nil,err
	}

	// now generate result with the given llm
	response, err := llm.CreateEmbedding(context.Background(), 
		splittedDocx,
	)
	if err != nil{
		fmt.Println("Error in creating embeddings.")
		return nil,err
	}

	// if not just print the response
	return response,nil
}

func createCollectionWithName(splittedDocx []string, embededBook [][]float32, s *models.Server) error{
	// now just create the collection and add the given one

	client := s.Store
	err := client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: "AI_Book_summarizer",
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
		Size:     uint64(len(embededBook[0])),
		Distance: qdrant.Distance_Cosine,
		}),
	})

	if err != nil{
		fmt.Println("Error in creating the collection.")
		return err
	}


	// store the given data inside of the given collection
	// create the points for the given embeddings
	var points []*qdrant.PointStruct

	// make the points by adding each of the above embeddings
	// point is going to contain -> id, vector, payload
	// payload is simply going to be the message 


	for i:=0; i<len(embededBook); i++{
		stringValue := splittedDocx[i]
		
		point := &qdrant.PointStruct{
			Id:       qdrant.NewIDNum(uint64(i+1)),
			Vectors:  qdrant.NewVectors(embededBook[i]...),
			Payload:  qdrant.NewValueMap(map[string]any{
				"textFile": stringValue,
			}),
		}

		points = append(points, point)
	}




	return nil
}

func vectorStore(splittedDocx []string, s *models.Server) error{
	// first create the embedding of each of the given docx
	response, err := createEmbeddings(splittedDocx, s)
	if err != nil{
		fmt.Println("Error in creating the embeddings.")
		return err
	}

	// create the collection and store the given response in it\
	err = createCollectionWithName(splittedDocx, response, s)
	if err != nil{
		return err
	}

	return nil
}