package storeextracteddata

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type StoreDataToLocation struct{
	BasePath string
}

func(s *StoreDataToLocation) StoreFile(fileName string, file multipart.File)(string, error){
	entries, err := os.ReadDir(s.BasePath)
	fmt.Println(entries)
	info, err := os.Stat(s.BasePath)
	if err != nil{
		if info == nil{
		err = os.Mkdir("./uploadedFiles", 0755)
		if err != nil{
			fmt.Println("[PDF HANDLER] Error in creating folder: ", err)
			}
		}
	}
	// create the file
	fullPath := s.BasePath + "/" + fileName
	copiedFile, err := os.Create(fullPath)
	if err != nil{
		fmt.Println("[storeFile] Error: ", err)
		return "", errors.New("Error in creating file.")
	}

	// make a copy of the file inside the given folder made above
	_, err = io.Copy(copiedFile, file)
	if err != nil{
		fmt.Println("[storeFile] Error: ", err)
		return "", errors.New("Error in storing file.")
	}

	return fullPath, nil
}