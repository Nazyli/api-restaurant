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
func ERROR(w http.ResponseWriter, statusCode int, err ...string) {
	if statusCode != http.StatusOK {
		statusText := http.StatusText(statusCode)

		res := Error{
			Error: ErrorMsg{
				Message: MessageError(statusCode, err...),
				Code:    statusCode,
				Title:   statusText,
			},
		}
		JSON(w, statusCode, res)
		return
	}
	JSON(w, statusCode, nil)
}

func OK(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, data)
}

func MessageError(code int, err ...string) string {
	if len(err) == 1 {
		return err[0]
	}
	switch code {
	case http.StatusNotFound:
		return err[0] + ", " + err[1] + " Data Not Found"
	case http.StatusInternalServerError:
		return err[0] + ", Unsuccessful Commuicating " + err[1] + " Data With Repository"
	default:
		return err[0] + ", " + err[1]
	}
}
