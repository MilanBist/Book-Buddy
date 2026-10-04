package storeextracteddata

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
)

func TestStoreFile(t *testing.T) {
	basePath := filepath.Join(t.TempDir(), "uploadedFiles")
	err := os.MkdirAll(basePath, 0755)
	if err != nil {
		t.Fatal("Error from portion first as : ",err)
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	pdfContent := []byte("%PDF-1.4\nFake PDF content\n%%EOF")
	part, err := writer.CreateFormFile("file", "test.pdf")
	if err != nil {
		t.Fatal(err)
	}

	_, err = part.Write(pdfContent)
	if err != nil {
		t.Fatal("Error from second part: ",err)
	}
	writer.Close()


	// convert to multipart.filee
	reader := multipart.NewReader(&body, writer.Boundary())
	form, err := reader.ReadForm(1 << 20) // 1 MB kept in memory
	if err != nil {
		t.Fatal(err)
	}

	fileHeader := form.File["file"][0]   // *multipart.FileHeader
	file, err := fileHeader.Open()       // multipart.File
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	data := &StoreDataToLocation{
		BasePath: basePath,
	}

	path, err := data.StoreFile(fileHeader.Filename, file)
	fmt.Println(path)
}