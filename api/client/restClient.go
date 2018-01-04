package client

import (
  "os"
  "fmt"
  "strings"
  "net/http"
  "encoding/json"
  "regexp"
  "github.com/olekukonko/tablewriter"
)

var linkListRegex = regexp.MustCompile("<([A-Za-z0-9\\:\\.\\/\\=\\?\\{\\}]*)>; rel=\"([A-Za-z]*)\"")

func GetRestApiURL(server, restPath string) (string) {
  // Remove leading `/` if provided.
  restPath = strings.TrimPrefix(restPath, "/")
  if server == "github.com" {
    return fmt.Sprintf("https://api.%s/%s", server, restPath)
  } else {
    return fmt.Sprintf("https://%s/api/v3/%s", server, restPath)
  }
}

func RestGet(apiURL, token string, responseHandler RestResponseHandler) {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader(responseHandler.TableHeader())
  paginatedRestGet(apiURL, token, table, responseHandler)
  table.Render()
}

func paginatedRestGet(apiURL, token string, table *tablewriter.Table, responseHandler RestResponseHandler) {
  req, err := http.NewRequest("GET", apiURL, nil)
  if err != nil {
    fmt.Printf("Failed while building the HTTP client: %s", err)
    return
  }

  // Provide authentication
  req.Header.Add("Authorization", fmt.Sprintf("bearer %s", token))

  client := http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    fmt.Printf("Error while querying the server: %s", err)
    return
  } else if resp.StatusCode != http.StatusOK {
    fmt.Printf("Ooops... sorry, server sent a %d HTTP status code: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
  }

  // Close when method returns
  defer resp.Body.Close()

  var jsonObj []map[string]interface{}

  // Decode the JSON array
  decodeError := json.NewDecoder(resp.Body).Decode(&jsonObj)
  if decodeError != nil {
    fmt.Printf("Error while decoding the server response: %s", decodeError)
    return
  } else  {
    table.AppendBulk(responseHandler.TableRows(jsonObj))

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
          paginatedRestGet(linkElement[1], token, table, responseHandler)
          return
        }
	    }
    }
  }
}

func RestPost(apiURL, token, params string) map[string]interface{} {
  req, err := http.NewRequest("POST", apiURL, strings.NewReader(params))
  if err != nil {
    fmt.Printf("Failed while building the HTTP client: %s", err)
    return nil
  }

  // Provide authentication
  req.Header.Add("Authorization", fmt.Sprintf("bearer %s", token))

  client := http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    fmt.Printf("Error while querying the server: %s", err)
    return nil
  } else if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
    fmt.Printf("Ooops... sorry, server sent a %d HTTP status code: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
  }

  // Close when method returns
  defer resp.Body.Close()

  var jsonObj map[string]interface{}

  // Decode the JSON array
  decodeError := json.NewDecoder(resp.Body).Decode(&jsonObj)
  if decodeError != nil {
    fmt.Printf("Error while decoding the server response: %s", decodeError)
    return nil
  } else  {
    return jsonObj
  }

}
