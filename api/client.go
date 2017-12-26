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

func DoRestApiCall(apiURL, token string, resultDecoder ResponseDecoder) {
  req, err := http.NewRequest("GET", apiURL, nil)
  if err != nil {
    log.Fatal("Failed while building the HTTP client: ", err)
    return
  }

  // Provide authentication
  req.Header.Add("Authorization", fmt.Sprintf("bearer %s", token))

  client := http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    log.Fatal("Error while querying the server.", err)
    return
  } else if resp.StatusCode != http.StatusOK {
    log.Fatalf("Ooops... sorry, server sent a %d HTTP status code: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
  }

  // Close when method returns
  defer resp.Body.Close()

  // Decode the JSON array
  decoder := json.NewDecoder(resp.Body)
  decodeError := resultDecoder.Decode(decoder)
  if decodeError != nil {
    log.Fatal("Error while decoding the server response.", decodeError)
    return
  } else  {
    // Working the pagination as described in https://developer.github.com/guides/traversing-with-pagination/
    // linkHeader is '<https://gheserver.com/api/v3/users?since=35>; rel="next", <https://gheserver.com/api/v3/users{?since}>; rel="first"'
    linkHeader := resp.Header.Get("Link")

    if linkHeader != "" {
      linkArray := linkListRegex.FindAllStringSubmatch(linkHeader, -1)
      /*
      linkArray := [["<https://gheserver.com/api/v3/users?since=35>; rel=next", "https://gheserver.com/api/v3/users?since=35", "next"],
       ["<https://gheserver.com/api/v3/users{?since}>; rel="first", "https://gheserver.com/api/v3/users{?since}, "first"]]
      */

      for _, linkElement := range linkArray {
        if linkElement[2] == "next" {
          DoRestApiCall(linkElement[1], token, resultDecoder)
          return
        }
	    }
    }
  }
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
