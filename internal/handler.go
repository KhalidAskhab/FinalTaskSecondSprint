package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed) //405
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest) //400
		return
	}

	result, err := Calc(req.Expression)
	if err != nil {
		if strings.Contains(err.Error(), "некорректное выражение") {
			http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity) //422
		} else {
			http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError) //500
		}
		return
	}

	var resultStr string
	if result == float64(int(result)) {
		resultStr = fmt.Sprintf("%d", int(result))
	} else {
		resultStr = strconv.FormatFloat(result, 'f', -1, 64)
	}

	response := Response{Result: resultStr}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
