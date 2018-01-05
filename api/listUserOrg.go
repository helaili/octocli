package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
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
  client.BasicGraphQLResponseHandler
}

func (this *UserOrgsResponseHandler) TableHeader() []string {
  return []string{"name"}
}

func (this *UserOrgsResponseHandler) TableRows(jsonObj map[string]interface{}) [][]string {
  table := [][]string{}
  nodes := this.GetNodes(jsonObj, this.ResultPath())
  for _, org := range nodes {
    row := []string{fmt.Sprintf("%v", org.(map[string]interface{})["name"])}
    table = append(table, row)
  }
  return table
}

func (this *UserOrgsResponseHandler) ResultPath() []string {
  return []string{"data", "user", "organizations"}
}

func GetUserOrgs(server, token, user string) {
  params := map[string]interface{}{"login" : user}
  orgHandler := UserOrgsResponseHandler{}
  client.GraphQLPost(server, token, userOrgsQuery, params, &orgHandler)
}
