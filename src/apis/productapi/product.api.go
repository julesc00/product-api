package productapi

import (
	"encoding/json"
	"example/to_from_json_api/src/entities"
	"net/http"
)

func Find(response http.ResponseWriter, request *http.Request) {
	product := entities.Product{
		Id:       "p01",
		Name:     "name 1",
		Price:    5.9,
		Quantity: 10,
		Status:   true,
	}
	respondWithJson(response, http.StatusOK, product)
}

func respondWithJson(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	_, err := response.Write(result)
	if err != nil {
		return
	}
}
