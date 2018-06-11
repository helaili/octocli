package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

var listReposQuery = `query($login:String!, $count:Int!, $cursor:String) {
  organization(login: $login) {
  	name
    repositories(first: $count, after: $cursor) {
      pageInfo {
       endCursor
       hasNextPage
      }
      nodes {
        name
        description
      }
    }
  }
}`

type ListReposQueryResponseHandler struct {
  client.BasicGraphQLResponseHandler
}

func (this *ListReposQueryResponseHandler) TableHeader() []string {
  return []string{"name", "description"}
}

func (this *ListReposQueryResponseHandler) TableRows(jsonObj map[string]interface{}) [][]string {
  table := [][]string{}
  nodes := this.GetNodes(jsonObj, this.ResultPath())
  if nodes != nil {
    for _, org := range nodes {
      row := []string{
        fmt.Sprintf("%v", org.(map[string]interface{})["name"]),
        fmt.Sprintf("%v", org.(map[string]interface{})["description"])}
      table = append(table, row)
    }
    return table
  } else {
    return nil
  }
}

func (this *ListReposQueryResponseHandler) ResultPath() []string {
  return []string{"data", "organization", "repositories"}
}

func PrintRepos(org string) {
  params := map[string]interface{}{"login" : org}
  handler := ListReposQueryResponseHandler{}
  client.GraphQLQueryAndPrintTable(listReposQuery, params, &handler)
}
