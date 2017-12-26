package api

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "regexp"
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

type UserOrgsQueryParams struct {
  Login string `json:"login"`
  Count int `json:"count"`
  Cursor string `json:"cursor,omitempty"`
}

type UserOrgsResponse struct {
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

type OrgResponse struct {
  Login string `json:"login"`
  Id int `json:"id"`
}

type OrgResponseDecoder struct {}

func (org OrgResponseDecoder) Decode(decoder *json.Decoder) error {
  var orgs []OrgResponse
  decodeError := decoder.Decode(&orgs)
  if decodeError == nil {
    for _, org := range orgs {
      fmt.Println(org)
    }
  }
  return decodeError
}

var linkListRegex = regexp.MustCompile("<([A-Za-z0-9\\:\\.\\/\\=\\?\\{\\}]*)>; rel=\"([A-Za-z]*)\"")

func GetAllOrgs(server, token string) {
  apiURL := GetRestApiURL(server, "organizations")
  getAllOrgs(apiURL, token)
}

func getAllOrgs(apiURL, token string) {
  orgDecoder := OrgResponseDecoder{}
  DoRestApiCall(apiURL, token, orgDecoder)
}

func GetUserOrgs(server, token, user string) {
  getUserOrgs(server, token, user, "")
}


func getUserOrgs(server, token, user, cursor string) {
  params := UserOrgsQueryParams{Login: user, Count: 100, Cursor: cursor}
  query := GraphQLQuery{userOrgsQuery, params}

  resp, err := DoGraphQLApiCall(server, token, query)
  if err != nil {
    log.Fatal("Error while querying the server.", err)
    return
  } else if resp.StatusCode != http.StatusOK {
    log.Fatalf("Ooops... sorry, server sent a %d HTTP status code: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
  }

  // Close when method returns
  defer resp.Body.Close()

  var data UserOrgsResponse
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
        getUserOrgs(server, token, user, data.Data.User.Organizations.PageInfo.EndCursor)
    }
  }
}
