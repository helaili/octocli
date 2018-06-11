package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

type TeamRepositoriesResponseHandler struct {
}

func (this *TeamRepositoriesResponseHandler) TableHeader() []string {
  return []string{"name", "description", "private", "fork", "permission"}
}

func (this *TeamRepositoriesResponseHandler) TableRows(jsonArray []map[string]interface{}) [][]string {
  table := [][]string{}
  for _, repo := range jsonArray {
    permission := "read"
    permissions := repo["permissions"].(map[string]interface {})
    if permissions["admin"].(bool) {
      permission = "admin"
    } else if permissions["push"].(bool) {
      permission = "write"
    }
    row := []string{
      fmt.Sprintf("%v", repo["name"]),
      fmt.Sprintf("%v", repo["description"]),
      fmt.Sprintf("%v", repo["private"]),
      fmt.Sprintf("%v", repo["fork"]),
      permission}
    table = append(table, row)
  }
  return table
}

func PrintTeamRepositories(org, teamSlug string) {
  // Need to figure out the team ID
  teamObj := GetRestTeam(org, teamSlug)

  if teamObj == nil || teamObj["id"] == nil {
    fmt.Printf("Couldn't find a team with slug %s\n", teamSlug)
    return
  }

  apiURL := client.GetRestApiURL(fmt.Sprintf("/teams/%d/repos", int(teamObj["id"].(float64))))
  handler := TeamRepositoriesResponseHandler{}
  client.RestGetAndPrintTable(apiURL, &handler)
}
