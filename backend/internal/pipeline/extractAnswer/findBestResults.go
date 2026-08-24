package extractanswer

import (
	"context"
	"errors"
	"fmt"
	"github.com/qdrant/go-client/qdrant"
)

func uint64Ptr(v uint64) *uint64 {
    return &v
}


func FindBestEmbeddings(client *qdrant.Client, point []float32, bookName string)([]string, error){
	searchResult, err := client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: "AI_Book_summarizer",
		Query: qdrant.NewQuery(point...),
		Limit: uint64Ptr(10),
		WithPayload: qdrant.NewWithPayload(true),
		Filter: &qdrant.Filter{
			Must: []*qdrant.Condition{
				qdrant.NewMatch("bookTitle",bookName),
			},
		},
	})
	if err != nil{
		fmt.Println(err)
		return nil, errors.New("No table")
	}

	var bestResults []string

	// get first 5 emebeddings and if the rating of the emebedding is greater than 0.9 append them
	for _, value := range searchResult {
			if textValue, ok := value.Payload["text"]; ok {
				if len(bestResults) > 5{
					// check for the embedding rating
					if value.Score > 0.9{
						bestResults = append(bestResults, value.Payload["text"].GetStringValue())
					}
				}
				bestResults = append(bestResults, textValue.GetStringValue())
			}
	}
	return bestResults, nil
}