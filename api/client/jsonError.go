package client

import (
  "fmt"
  "net/http"
  "encoding/json"
)


type JSONErrorDetails struct {
  Code string `json:"code"`
  Field string `json:"field"`
  Resource string `json:"resource"`
}

type JSONError struct {
  Message string `json:"message"`
  Errors []JSONErrorDetails `json:"errors"`
}


func GetJSONError(resp *http.Response) (JSONError, error) {
  var jsonObj JSONError

  // Decode the JSON array
  decodeError := json.NewDecoder(resp.Body).Decode(&jsonObj)
  if decodeError != nil {
    fmt.Printf("Error while decoding the error object: %s\n", decodeError)
    return JSONError{}, decodeError
  } else  {
    return jsonObj, nil
  }
}
