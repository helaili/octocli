package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

var repoAuditQuery = `query($name:String!, $owner:String!, $count:Int!, $cursor:String) {
	repository(name: $name, owner: $owner) {
    mentionableUsers(first: $count, after: $cursor) {
      pageInfo {
       endCursor
       hasNextPage
      }
      nodes {
       login,
       name,
       email,
       isEmployee
     }
   }
 }
}`

type RepoAuditResponseHandler struct {
  client.BasicGraphQLResponseHandler
  isGitHubber bool
}

func (this *RepoAuditResponseHandler) TableHeader() []string {
  return []string{"login", "name", "email", "GitHubber"}
}

func (this *RepoAuditResponseHandler) TableRows(jsonObj map[string]interface{}) [][]string {
  table := [][]string{}
  nodes := this.GetNodes(jsonObj, this.ResultPath())

  if nodes != nil {
    var row []string
    for _, org := range nodes {
      if this.isGitHubber {
        row = []string{
          fmt.Sprintf("%v", org.(map[string]interface{})["login"]),
          fmt.Sprintf("%v", org.(map[string]interface{})["name"]),
          fmt.Sprintf("%v", org.(map[string]interface{})["email"]),
          fmt.Sprintf("%v", org.(map[string]interface{})["isEmployee"])}
      } else {
        row = []string{
          fmt.Sprintf("%v", org.(map[string]interface{})["login"]),
          fmt.Sprintf("%v", org.(map[string]interface{})["name"]),
          fmt.Sprintf("%v", org.(map[string]interface{})["email"])}
      }
      table = append(table, row)
    }
    return table
  } else {
    return nil
  }
}

func (this *RepoAuditResponseHandler) ResultPath() []string {
  return []string{"data", "repository", "mentionableUsers"}
}

func PrintRepoAudit(server, token, user, name string, isGitHubber bool) {
  params := map[string]interface{}{"name" : name, "owner": user}
  handler := RepoAuditResponseHandler{}
  handler.isGitHubber = isGitHubber
  client.GraphQLQueryAndPrintTable(server, token, repoAuditQuery, params, &handler)
}
