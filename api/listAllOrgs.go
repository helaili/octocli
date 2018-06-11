package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

type OrgResponseHandler struct {
}

func (this *OrgResponseHandler) TableHeader() []string {
  return []string{"login", "email", "location", "description"}
}

func (this *OrgResponseHandler) TableRows(jsonArray []map[string]interface{}) [][]string {
  table := [][]string{}
  for _, org := range jsonArray {
    row := []string{
      fmt.Sprintf("%v", org["login"]),
      fmt.Sprintf("%v", org["email"]),
      fmt.Sprintf("%v", org["location"]),
      fmt.Sprintf("%v", org["description"])}
    table = append(table, row)
  }
  return table
}

func PrintAllOrgs() {
  apiURL := client.GetRestApiURL("/organizations")
  orgHandler := OrgResponseHandler{}
  client.RestGetAndPrintTable(apiURL, &orgHandler)
}
