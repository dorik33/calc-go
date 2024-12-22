package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dorik33/calc-go/internal/model"
	"github.com/dorik33/calc-go/pkg/calculate"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(model.Response{Error: fmt.Sprintf("Error reading request body: %v", err)})
		w.Write(response)
		log.Printf("Error reading request body: %v", err)
		return
	}

	var req model.Request
	err = json.Unmarshal(data, &req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(model.Response{Error: fmt.Sprintf("Error parsing JSON: %v", err)})
		w.Write(response)
		log.Printf("Error parsing JSON: %v", err)
		return
	}
	strings.ReplaceAll(req.Expression, " ", "")
	result, err := calculate.Calculate(req.Expression)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		response, _ := json.Marshal(model.Response{Error: fmt.Sprintf("%s", err)})
		w.Write(response)
		log.Printf("Calculation error: %v, expression: %s", err, req.Expression)
		return
	}

	log.Printf("Calculation completed successfully: %.4f", result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(model.Response{Result: fmt.Sprintf("%.4f", result)})
	w.Write(response)
}
