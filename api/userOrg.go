package api

import (
  "fmt"
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

type UserOrgsResponseHandler struct {
}

func (this *UserOrgsResponseHandler) Print(jsonObj map[string]interface{})  {
  for _, org := range jsonObj {
    fmt.Printf("Name: %s\n", org)
  }
}

func (this *UserOrgsResponseHandler) TableHeader() []string {
  return []string{"name"}
}

func (this *UserOrgsResponseHandler) TableRows(jsonObj map[string]interface{}) [][]string {
  table := [][]string{}
  nodes := getNodes(jsonObj, []string{"data", "user", "organizations"})
  for _, org := range nodes {
    row := []string{fmt.Sprintf("%v", org.(map[string]interface{})["name"])}
    table = append(table, row)
  }
  return table
}

func (this *UserOrgsResponseHandler) PageInfoPath() []string {
  return []string{"data", "user", "organizations"}
}


func GetUserOrgs(server, token, user string) {
  params := map[string]interface{}{"login" : user}
  orgHandler := UserOrgsResponseHandler{}
  DoGraphQLApiCall(server, token, userOrgsQuery, params, &orgHandler)
}
