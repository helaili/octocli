package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

type OrgResponseHandler struct {
}

func (this *OrgResponseHandler) TableHeader() []string {
  return []string{"login", "description"}
}

func (this *OrgResponseHandler) TableRows(jsonArray []map[string]interface{}) [][]string {
  table := [][]string{}
  for _, org := range jsonArray {
    row := []string{fmt.Sprintf("%v", org["login"]), fmt.Sprintf("%v", org["description"])}
    table = append(table, row)
  }
  return table
}

func GetAllOrgs(server, token string) {
  apiURL := client.GetRestApiURL(server, "/organizations")
  orgHandler := OrgResponseHandler{}
  client.RestGet(apiURL, token, &orgHandler)
}
