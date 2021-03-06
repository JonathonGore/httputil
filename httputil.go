package httputil

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Success writes a JSON response object with a single message key indicating
// success. No status code is written as this gives more flexibility, due to
// the fact that Golang http response defaults to 200 if no status code is written.
func Success(w http.ResponseWriter) {
	w.Write(JSON(SuccessResponse{Message: "success", Code: http.StatusOK}))
}

// Error is a utility function to responsd to a response writer with
// the given message and error code.
func Error(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	w.Write(JSON(ErrorResponse{message, code}))
}

// JSONString consumes an interface and marshals it into a JSON representation
// in string format.
func JSONString(e interface{}) string {
	return string(JSON(e))
}

// JSON consumes an interface and marshals it into a JSON representation
// stored in a byte slice.
func JSON(e interface{}) []byte {
	contents, err := json.Marshal(e)
	if err != nil {
		return []byte{}
	}

	return contents
}

// UnmarshalRequestBody consumes an http request and marshals the body into the
// provided interface.
func UnmarshalRequestBody(r *http.Request, v interface{}) error {
	if r == nil {
		return errors.New("http request must be non-nil")
	}

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(contents, v)
	if err != nil {
		return err
	}

	return nil
}
