package client

import (
  "os"
  "fmt"
  "bytes"
  "errors"
  "strings"
  "net/http"
  "encoding/json"
  "github.com/spf13/viper"
  "github.com/olekukonko/tablewriter"
)

type GQLQuery struct {
  Query string `json:"query"`
  Variables interface{} `json:"variables"`
}

func GetGraphQLApiURL(server string) (string) {
  if server == "github.com" {
    return fmt.Sprintf("https://api.%s/graphql", server)
  } else {
    return fmt.Sprintf("https://%s/api/graphql", server)
  }
}

func GraphQLQueryAndPrintTable(query string, params map[string]interface{}, responseHandler GraphQLResponseHandler) {
  table := tablewriter.NewWriter(os.Stdout)

  if viper.GetBool("markdown") {
    table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
    table.SetCenterSeparator("|")
  }

  table.SetHeader(responseHandler.TableHeader())
  err := paginatedGraphQLQueryAndPrintTable(query, params, table, responseHandler)
  if err != nil {
    fmt.Println(err)
  } else {
    table.Render()
  }
}

func paginatedGraphQLQueryAndPrintTable(query string, params map[string]interface{}, table *tablewriter.Table, responseHandler GraphQLResponseHandler) error {
  if params["count"] == nil {
    params["count"] = 100
  }
  graphQLQuery := GQLQuery{query, params}
  jsonValue, _ := json.Marshal(graphQLQuery)
  apiURL := GetGraphQLApiURL(viper.GetString("server"))
  req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonValue))
  if err != nil {
    return errors.New(fmt.Sprintf("Failed while building the HTTP client: %s\n", err))
  }

  // Provide authentication
  req.Header.Add("Authorization", fmt.Sprintf("bearer %s", viper.GetString("token")))

  client := http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return errors.New(fmt.Sprintf("Error while querying the server: %s\n", err))
  } else if resp.StatusCode != http.StatusOK {
    return errors.New(fmt.Sprintf("Ooops... sorry, server sent a %d HTTP status code: %s\n", resp.StatusCode, http.StatusText(resp.StatusCode)))
  }

  // Close when method returns
  defer resp.Body.Close()

  var jsonObj map[string]interface{}

  // Decode the JSON array
  decodeError := json.NewDecoder(resp.Body).Decode(&jsonObj)
  if decodeError != nil {
    return errors.New(fmt.Sprintf("Error while decoding the server response: %s\n", decodeError))
  } else {
    rows := responseHandler.TableRows(jsonObj)

    if rows == nil {
      if jsonObj["errors"] != nil && jsonObj["errors"].([]interface {})[0].(map[string]interface{})["message"] != nil {
        if resp.Header.Get("X-GitHub-SSO") != "" {
          // SSO enabled, the token needs to be whitlisted
          return errors.New(fmt.Sprintf("%s Navigate to the following URL to whitelist your personal access token: %s",
                            jsonObj["errors"].([]interface {})[0].(map[string]interface{})["message"],
                            strings.TrimPrefix(resp.Header.Get("X-GitHub-SSO"), "required; url=")))
        } else {
          return errors.New(jsonObj["errors"].([]interface {})[0].(map[string]interface{})["message"].(string))
        }
      } else {
        return errors.New("Data access problem. Please check your input values")
      }
    }

    table.AppendBulk(rows)
    hasNextPage, endCursor := getPageInfo(jsonObj, responseHandler.ResultPath())
    if hasNextPage {
      params["cursor"] = endCursor
      return paginatedGraphQLQueryAndPrintTable(query, params, table, responseHandler)
    } else {
      return nil
    }
  }
}

// Navigate the JSON response to retrive the 'pageInfo' object and return its prorperies (hasNextPage and endCursor)
func getPageInfo(jsonObj map[string]interface{}, path []string) (hasNextPage bool, endCursor string) {
  if(len(path) == 0) {
    pageInfo := jsonObj["pageInfo"].(map[string]interface{})
    if pageInfo["hasNextPage"].(bool) {
      return true, pageInfo["endCursor"].(string)
    } else {
      return false, ""
    }
  } else {
    if jsonObj[path[0]] != nil {
      return getPageInfo(jsonObj[path[0]].(map[string]interface{}), path[1:])
    } else {
      return false, ""
    }
  }
}


func GraphQLQuery(query string, params map[string]interface{}) (resp *http.Response, err error) {
  graphQLQuery := GQLQuery{query, params}
  jsonValue, _ := json.Marshal(graphQLQuery)
  apiURL := GetGraphQLApiURL(viper.GetString("server"))
  req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonValue))
  if err != nil {
    fmt.Printf("Failed while building the HTTP client: %s\n", err)
    return
  }

  // Provide authentication
  req.Header.Add("Authorization", fmt.Sprintf("bearer %s", viper.GetString("token")))

  client := http.Client{}
  return client.Do(req)
}


func GraphQLQueryObject(query string, params map[string]interface{}) map[string]interface{} {
  resp, err := GraphQLQuery(query, params)
  if err != nil {
    fmt.Printf("Error while querying the server: %s\n", err)
    return nil
  } else if resp.StatusCode != http.StatusOK {
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
  } else {
    return jsonObj
  }
}
