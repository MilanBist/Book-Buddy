package extractpdf

import (
	"context"
	"fmt"
	"strings"
	"os"
	"strconv"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/qdrant/go-client/qdrant"
	"github.com/tmc/langchaingo/llms/ollama"
)

func CreateEmbeddings(splittedDocx []string, s *models.Server) ([][]float32, error) {
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

func createCollectionWithName(splittedDocx []string, embededBook [][]float32, s *models.Server, bookTitle string) error{
	// now just create the collection and add the given one

	client := s.Store
	collectionName := "AI_Book_summarizer"
	// check if the collection exists

	exists, err := client.CollectionExists(context.Background(), collectionName)
	if err != nil{
		// collection doesn't exist
		fmt.Println("Collection doesn't exist so creating a new one.")
	}

	fmt.Println("[COLLECTION EXISTENCE]: ", exists)

	if !exists{
		err = client.CreateCollection(context.Background(), &qdrant.CreateCollection{
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
	}

	// length of the splittedDocx
	fmt.Println("The length is : ", len(splittedDocx))

	


	// store the given data inside of the given collection
	// create the points for the given embeddings
	var points []*qdrant.PointStruct

	// make the points by adding each of the above embeddings
	// point is going to contain -> id, vector, payload
	// payload is simply going to be the message 

	// store the id in the .txt file so that it can be reused
	// create the new file
	// first check the file if it exists
	

	// check for the .txt file in the output
	var stringedData string
	var numId int
	path := "./output"
	data, err := os.ReadFile(path+"/output.txt")
	if err != nil || len(data)==0{
		file, err := os.Create(path+"/output.txt")
		if err != nil{
			return err
		}
		file.Close()
		numId = 1
	} else{
		stringedData = string(data)
		numId, _ = strconv.Atoi(stringedData)
	}	


	for i:=0; i<len(embededBook); i++{
		point := &qdrant.PointStruct{
			Id:       qdrant.NewIDNum(uint64(numId)),
			Vectors:  qdrant.NewVectors(embededBook[i]...),
			Payload:  qdrant.NewValueMap(map[string]any{
				"text": splittedDocx[i],
				"chunkId": strconv.Itoa(i+1),
				"bookTitle": bookTitle,
			}),
		}
		numId += 1
		points = append(points, point)
	}

	//insert into the output.txt the value of numId
	file, err := os.OpenFile(path+"/output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil{
		fmt.Println("Error in opening the file.")
		return err
	}
	defer file.Close()
	file.WriteString(strconv.Itoa(numId))



	// insert the data in the certain database
	operationInfo, err := client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: "AI_Book_summarizer",
		Points: points,
	})

	if err != nil{
		fmt.Println("Error in inserting the data.")
		return err
	}
	fmt.Println(operationInfo)
	return nil
}

func VectorStore(splittedDocx []string, s *models.Server, fileName string) error{
	newPath := "./output"
	_, err := os.ReadDir(newPath)

	if err != nil{
		err := os.MkdirAll(newPath, 0755)
		if err != nil{
			return err
		}
	}
	
	// get the book name
	path := strings.Split(fileName, "/")
	// // set the last name to be the book title
	bookTitle := path[len(path)-1]
	// remove the extension
	book := strings.Split(bookTitle, ".")
	bookTitle = book[0]
	fmt.Println("Book title saved: ", bookTitle)

	// first create the embedding of each of the given docx
	response, err := CreateEmbeddings(splittedDocx, s)
	if err != nil{
		fmt.Println("Error in creating the embeddings.")
		return err
	}

	fmt.Println(bookTitle)
	// create the collection and store the given response in it\
	err = createCollectionWithName(splittedDocx, response, s, bookTitle)
	if err != nil{
		return err
	}


	fmt.Println("File name with the giiven is stored. ", bookTitle)

	return nil
}