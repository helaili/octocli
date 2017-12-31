package api

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "regexp"
)

var linkListRegex = regexp.MustCompile("<([A-Za-z0-9\\:\\.\\/\\=\\?\\{\\}]*)>; rel=\"([A-Za-z]*)\"")

func GetRestApiURL(server, restPath string) (string) {
  if server == "github.com" {
    return fmt.Sprintf("https://api.%s/%s", server, restPath)
  } else {
    return fmt.Sprintf("https://%s/api/v3/%s", server, restPath)
  }
}

func DoRestApiCall(apiURL, token string, responseHandler RestResponseHandler) {
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

  var jsonObj []map[string]interface{}

  // Decode the JSON array
  decodeError := json.NewDecoder(resp.Body).Decode(&jsonObj)
  if decodeError != nil {
    log.Fatal("Error while decoding the server response.", decodeError)
    return
  } else  {
    responseHandler.Print(jsonObj)
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
          DoRestApiCall(linkElement[1], token, responseHandler)
          return
        }
	    }
    }
  }
}
