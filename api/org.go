package api

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)


var userOrgsQuery = `query($login:String!, $count:Int!, $cursor:String) {
  user(login: $login) {
    organizations(first: $count, after: $cursor) {
      pageInfo {
       endCursor
       hasNextPage
      }
      nodes {
        name
      }
    }
  }
}`

type userOrgsQueryParams struct {
  Login string `json:"login"`
  Count int `json:"count"`
  Cursor string `json:"cursor,omitempty"`
}

type userOrgsResponse struct {
  Data struct {
    User struct {
       Organizations struct {
         PageInfo struct {
           EndCursor string `json:"endCursor"`
           HasNextPage bool `json:"hasNextPage"`
         } `json:"pageInfo"`
         Nodes []struct {
           Name string `json:"name"`
         } `json:"nodes"`
       } `json:"organizations"`
     } `json:"user"`
  } `json:"data"`
  Errors []struct {
    Message string `json:"message"`
  } `json:"errors"`
}

func GetUserOrgs(server, token, user, cursor string) {
  params := userOrgsQueryParams{Login: user, Count: 100, Cursor: cursor}
  query := GraphQLQuery{userOrgsQuery, params}

  resp, err := DoApiCall(server, token, query)
  if err != nil {
    log.Fatal("Error while querying the server.", err)
    return
  } else if resp.StatusCode != http.StatusOK {
    log.Fatalf("Ooops... sorry, server sent a %d HTTP status code: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
  }

  // Close when method returns
  defer resp.Body.Close()

  var data userOrgsResponse
  // Decode the JSON array
  decodeError := json.NewDecoder(resp.Body).Decode(&data)
  if decodeError != nil {
    log.Fatal("Error while decoding the server response.", decodeError)
    return
  } else if data.Errors != nil {
    // Getting an error within the GraphQL response body.
    log.Println("Server sent back some errors!")
    for _, error := range  data.Errors {
      log.Printf(" - %s", error.Message)
    }
    return
  } else  {
    for _, org := range  data.Data.User.Organizations.Nodes {
      fmt.Printf("Name: %s\n", org.Name)
    }
    if data.Data.User.Organizations.PageInfo.HasNextPage {
        GetUserOrgs(server, token, user, data.Data.User.Organizations.PageInfo.EndCursor)
    }
  }
}
