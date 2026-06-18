package extractanswer

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func uint64Ptr(v uint64) *uint64 {
    return &v
}


func FindBestEmbeddings(client *qdrant.Client, point []float32)([]string, error){
	searchResult, err := client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: "AI_Book_summarizer",
		Query: qdrant.NewQuery(point...),
		Limit: uint64Ptr(10),
		WithPayload: qdrant.NewWithPayload(true),
	})
	if err != nil{
		fmt.Println("Error in getting the result. ")
		return nil, err
	}

	var bestResults []string

	for _, value := range searchResult {
    	if textValue, ok := value.Payload["text"]; ok {
        	bestResults = append(bestResults, textValue.GetStringValue())
    	}
	}

	// return thi one to the api folder

	return bestResults, nil

}