package api

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)


type FakeStoreFileToLocation struct{
	FullPath 	string
	Error 		error
}

func (f *FakeStoreFileToLocation) StoreFile(fileName string, file multipart.File) (string, error){
	return f.FullPath, f.Error
}

type FakeStoreToDb 	struct{
	Status 		int
	UserId 		int
	Error 		error
}
func (f *FakeStoreToDb) AddBookIdentity(bookTitle string, userId int) (int, int, error){
	return f.Status, f.UserId, f.Error
}

type FakeExtract struct{
	Status 	int
	Error 	error
}
func (f *FakeExtract) ExtractData(file *multipart.File, filePath string) (int, error){
	return f.Status, f.Error
}

type HandlePdf struct{	
	Success bool        	`json:"success"`
    Message string      	`json:"message"`
    Data    models.BookData `json:"data"`
}

func TestHandleRawPdf(t *testing.T){
	fakeStoreFile := &FakeStoreFileToLocation{
		FullPath: "location1/",
		Error: nil,
	}

	fakeStoreDb := &FakeStoreToDb{
		Status: 200,
		UserId: 1,
		Error: nil,
	}

	fakeStoreExtract := &FakeExtract{
		Status: 200,
		Error: nil,
	}

	pdfHandler := &RawPdfHandler{
		GetFileLocation: fakeStoreFile,
		StoreBook: fakeStoreDb,
		ExtractPdf: fakeStoreExtract,
	}

	// for the pdf file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("document", "book.pdf")
	if err != nil {
		t.Fatal(err)
	}

	fakedPdfData := []byte("fake raw file contents")

	_, err = part.Write(fakedPdfData)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/handlePdf", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	ctx := context.WithValue(req.Context(), "userId", 1)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()
	pdfHandler.HandleRawPdfFile(res, req)

	if res.Code != 200{
		t.Fatalf("Required 200 and got %d", res.Code)
	}
	var responseData HandlePdf
	json.NewDecoder(res.Body).Decode(&responseData)

	bookId := responseData.Data.BookId
	if bookId != 1{
		t.Fatalf("Required 1 and got book id as : %d", bookId)
	}
}