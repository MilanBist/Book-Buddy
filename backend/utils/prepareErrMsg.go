package utils

import (
	"net/http"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)


func PrepareErrorMessage(msg string) models.ErrorResponse{
	var message models.ErrorResponse

	message.Success = false
	message.ErrMsg.Code = http.StatusInternalServerError
	message.ErrMsg.Message = msg


	return message
}