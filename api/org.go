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

type orgResponse struct {
  Login string `json:"login"`
  Id int `json:"id"`
}

var linkListRegex = regexp.MustCompile("<([A-Za-z0-9\\:\\.\\/\\=\\?\\{\\}]*)>; rel=\"([A-Za-z]*)\"")

func GetAllOrgs(server, token string) {
  apiURL := GetRestApiURL(server, "organizations")
  getAllOrgs(apiURL, token)
}

func getAllOrgs(apiURL, token string) {
  resp, err := DoRestApiCall(apiURL, token)
  if err != nil {
    log.Fatal("Error while querying the server.", err)
    return
  } else if resp.StatusCode != http.StatusOK {
    log.Fatalf("Ooops... sorry, server sent a %d HTTP status code: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
  }

  // Close when method returns
  defer resp.Body.Close()

  var orgs []orgResponse
  // Decode the JSON array
  decodeError := json.NewDecoder(resp.Body).Decode(&orgs)
  if decodeError != nil {
    log.Fatal("Error while decoding the server response.", decodeError)
    return
  } else  {
    for _, org := range  orgs {
      fmt.Printf("Name: %s\n", org.Login)
    }

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
          getAllOrgs(linkElement[1], token)
          return
        }
	    }
    }
  }

}

func GetUserOrgs(server, token, user string) {
  getUserOrgs(server, token, user, "")
}


func getUserOrgs(server, token, user, cursor string) {
  params := userOrgsQueryParams{Login: user, Count: 100, Cursor: cursor}
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
        getUserOrgs(server, token, user, data.Data.User.Organizations.PageInfo.EndCursor)
    }
  }
}
