package api

import (
  "fmt"
  "log"
  "bytes"
  "net/http"
  "encoding/json"
)

func GetGraphQLApiURL(server string) (string) {
  if server == "github.com" {
    return fmt.Sprintf("https://api.%s/graphql", server)
  } else {
    return fmt.Sprintf("https://%s/api/graphql", server)
  }
}

func GetRestApiURL(server, restPath string) (string) {
  if server == "github.com" {
    return fmt.Sprintf("https://api.%s/%s", server, restPath)
  } else {
    return fmt.Sprintf("https://%s/api/v3/%s", server, restPath)
  }
}

func DoRestApiCall(apiURL, token string) (response *http.Response, err error) {
  req, err := http.NewRequest("GET", apiURL, nil)
  if err != nil {
    log.Fatal("Failed while building the HTTP client: ", err)
    return
  }

  // Provide authentication
  req.Header.Add("Authorization", fmt.Sprintf("bearer %s", token))

  client := http.Client{}
  return client.Do(req)
}


func DoGraphQLApiCall(server, token string, query GraphQLQuery) (response *http.Response, err error) {
  jsonValue, _ := json.Marshal(query)
  apiURL := GetGraphQLApiURL(server)
  req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonValue))
  if err != nil {
    log.Fatal("Failed while building the HTTP client: ", err)
    return
  }

  // Provide authentication
  req.Header.Add("Authorization", fmt.Sprintf("bearer %s", token))

  client := http.Client{}
  return client.Do(req)
}
