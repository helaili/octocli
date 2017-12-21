package api

import (
  "fmt"
  "log"
  "bytes"
  "net/http"
  "encoding/json"
)

func GetApiURL(server string) (string) {
  if server == "github.com" {
    return fmt.Sprintf("https://api.%s/graphql", server)
  } else {
    return fmt.Sprintf("https://%s/api/graphql", server)
  }
}

func DoApiCall(server, token string, query GraphQLQuery) (response *http.Response, err error) {
  jsonValue, _ := json.Marshal(query)
  req, err := http.NewRequest("POST", GetApiURL(server), bytes.NewBuffer(jsonValue))
  if err != nil {
    log.Fatal("Failed while building the HTTP client: ", err)
    return
  }

  // Provide authentication
  req.Header.Add("Authorization", fmt.Sprintf("bearer %s", token))

  client := http.Client{}
  return client.Do(req)
}
