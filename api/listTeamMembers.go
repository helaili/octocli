package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

var listTeamMembersQuery = `query($login:String!, $slug:String!, $count:Int!, $cursor:String) {
  organization(login: $login) {
    team(slug: $slug) {
      name
      members(first: $count, after: $cursor) {
      	pageInfo {
       		endCursor
       		hasNextPage
      	}
      	nodes {
        	login
          name
      	}
      }
    }
  }
}`

type ListTeamMembersQueryResponseHandler struct {
  client.BasicGraphQLResponseHandler
}

func (this *ListTeamMembersQueryResponseHandler) TableHeader() []string {
  return []string{"login", "name"}
}

func (this *ListTeamMembersQueryResponseHandler) TableRows(jsonObj map[string]interface{}) [][]string {
  table := [][]string{}
  nodes := this.GetNodes(jsonObj, this.ResultPath())
  if nodes != nil {
    for _, org := range nodes {
      row := []string{
        fmt.Sprintf("%v", org.(map[string]interface{})["login"]),
        fmt.Sprintf("%v", org.(map[string]interface{})["name"])}
      table = append(table, row)
    }
    return table
  } else {
    return nil
  }
}

func (this *ListTeamMembersQueryResponseHandler) ResultPath() []string {
  return []string{"data", "organization", "team", "members"}
}

func PrintTeamMembers(org, team string) {
  params := map[string]interface{}{"login" : org, "slug" : team}
  handler := ListTeamMembersQueryResponseHandler{}
  client.GraphQLQueryAndPrintTable(listTeamMembersQuery, params, &handler)
}
