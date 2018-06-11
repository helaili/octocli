package client

import (
  "os"
  "fmt"
  "strings"
  "net/http"
  "encoding/json"
  "regexp"
	"github.com/spf13/viper"
  "github.com/olekukonko/tablewriter"
)

var linkListRegex = regexp.MustCompile("<([A-Za-z0-9\\:\\.\\/\\=\\?\\{\\}]*)>; rel=\"([A-Za-z]*)\"")

func GetRestApiURL(restPath string) (string) {
  // Remove leading `/` if provided.
  restPath = strings.TrimPrefix(restPath, "/")
  if viper.GetString("server") == "github.com" {
    return fmt.Sprintf("https://api.%s/%s", viper.GetString("server"), restPath)
  } else {
    return fmt.Sprintf("https://%s/api/v3/%s", viper.GetString("server"), restPath)
  }
}

func RestGetAndPrintTable(apiURL string, responseHandler RestResponseHandler) {
  table := tablewriter.NewWriter(os.Stdout)

  if viper.GetBool("markdown") {
    table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
    table.SetCenterSeparator("|")
  }

  table.SetHeader(responseHandler.TableHeader())
  paginatedRestGetAndPrintTable(apiURL, table, responseHandler)
  table.Render()
}

func paginatedRestGetAndPrintTable(apiURL string, table *tablewriter.Table, responseHandler RestResponseHandler) {
  resp, err := RestQuery(apiURL, "GET", "", nil)

  if err != nil {
    fmt.Printf("Error while querying the server: %s\n", err)
    return
  } else if resp.StatusCode != http.StatusOK {
    fmt.Printf("Ooops... sorry, server sent a %d HTTP status code: %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
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
          paginatedRestGetAndPrintTable(linkElement[1], table, responseHandler)
          return
        }
	    }
    }
  }
}

func RestPostForObject(apiURL, params string) map[string]interface{} {
  return RestQueryForObject(apiURL, "POST", params)
}

func RestGetForObject(apiURL string) map[string]interface{} {
  return RestQueryForObject(apiURL, "GET", "")
}

func RestGetForArray(apiURL string) []map[string]interface{} {
  return RestQueryForArray(apiURL, "GET", "")
}


func RestQueryForObject(apiURL, verb, params string) map[string]interface{} {
  resp, err := RestQuery(apiURL, verb, params, nil)
  if err != nil {
    fmt.Printf("Error while querying the server: %s\n", err)
    return nil
  } else if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
    fmt.Printf("Ooops... sorry, server sent a %d HTTP status code: %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
    return nil
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


func RestQueryForArray(apiURL, verb, params string) []map[string]interface{} {
  resp, err := RestQuery(apiURL, verb, params, nil)
  if err != nil {
    fmt.Printf("Error while querying the server: %s\n", err)
    return nil
  } else if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
    fmt.Printf("Ooops... sorry, server sent a %d HTTP status code: %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
    return nil
  }

  // Close when method returns
  defer resp.Body.Close()

  var jsonObj []map[string]interface{}

  // Decode the JSON array
  decodeError := json.NewDecoder(resp.Body).Decode(&jsonObj)
  if decodeError != nil {
    fmt.Printf("Error while decoding the server response: %s", decodeError)
    return nil
  } else  {
    return jsonObj
  }
}

func RestQuery(apiURL, verb, params string, headers map[string]string) (resp *http.Response, err error) {
  req, err := http.NewRequest(verb, apiURL, strings.NewReader(params))
  if err != nil {
    return nil, err
  }

  // Provide authentication
  req.Header.Add("Authorization", fmt.Sprintf("bearer %s", viper.GetString("token")))
  // Add other headers
  for header, value := range headers {
    req.Header.Add(header, value)
  }

  client := http.Client{}
  return client.Do(req)
}
