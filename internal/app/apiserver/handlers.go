package apiserver

import (
	"encoding/json"
	"github.com/evd1ser/go-homework-2/internal/app/models"
	"net/http"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *APIServer) PostGrab(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get All Articles POST /grab")

	var equation models.Equation

	err := json.NewDecoder(req.Body).Decode(&equation)

	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	api.store.Equation().Set(&equation)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(equation)
}

func (api *APIServer) GetSolve(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get solve GET /solve")
	equation := api.store.Equation().GetWithCalc()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(*equation)
}
