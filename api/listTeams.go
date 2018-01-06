package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

var listTeamsQuery = `query($login:String!, $count:Int!, $cursor:String) {
  organization(login: $login) {
  	name
    teams(first: $count, after: $cursor) {
      pageInfo {
       endCursor
       hasNextPage
      }
      nodes {
        name
        slug
        description
      }
    }
  }
}`

type ListTeamsQueryResponseHandler struct {
  client.BasicGraphQLResponseHandler
}

func (this *ListTeamsQueryResponseHandler) TableHeader() []string {
  return []string{"name", "slug", "description"}
}

func (this *ListTeamsQueryResponseHandler) TableRows(jsonObj map[string]interface{}) [][]string {
  table := [][]string{}
  nodes := this.GetNodes(jsonObj, this.ResultPath())
  for _, org := range nodes {
    row := []string{
      fmt.Sprintf("%v", org.(map[string]interface{})["name"]),
      fmt.Sprintf("%v", org.(map[string]interface{})["slug"]),
      fmt.Sprintf("%v", org.(map[string]interface{})["description"])}
    table = append(table, row)
  }
  return table
}

func (this *ListTeamsQueryResponseHandler) ResultPath() []string {
  return []string{"data", "organization", "teams"}
}

func PrintTeams(server, token, org string) {
  params := map[string]interface{}{"login" : org}
  handler := ListTeamsQueryResponseHandler{}
  client.GraphQLQueryAndPrintTable(server, token, listTeamsQuery, params, &handler)
}
