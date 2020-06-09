package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Error ErrorMsg `json:"error"`
}
type ErrorMsg struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// JSON . . .
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR . . .
func ERROR(w http.ResponseWriter, statusCode int, err string) {
	if statusCode != http.StatusOK {
		statusText := http.StatusText(statusCode)
		res := Error{
			Error: ErrorMsg{
				Message: err,
				Code:    statusCode,
				Title:   statusText,
			},
		}
		JSON(w, statusCode, res)
		return
	}
	JSON(w, statusCode, nil)
}
