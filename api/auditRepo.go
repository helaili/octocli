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
  showIsGitHubber bool
}

func (this *RepoAuditResponseHandler) TableHeader() []string {
  if this.showIsGitHubber {
    return []string{"login", "name", "email", "GitHubber"}
  } else {
    return []string{"login", "name", "email"}
  }
}

func (this *RepoAuditResponseHandler) TableRows(jsonObj map[string]interface{}) [][]string {
  table := [][]string{}
  nodes := this.GetNodes(jsonObj, this.ResultPath())

  if nodes != nil {
    var row []string
    for _, org := range nodes {
      if this.showIsGitHubber {
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

func PrintRepoAudit(user, name string, showIsGitHubber bool) {
  params := map[string]interface{}{"name" : name, "owner": user}
  handler := RepoAuditResponseHandler{}
  handler.showIsGitHubber = showIsGitHubber
  client.GraphQLQueryAndPrintTable(repoAuditQuery, params, &handler)
}
