package api

import (
  "fmt"
)

type OrgResponseHandler struct {
}

func (this *OrgResponseHandler) Print(jsonArray []map[string]interface{})  {
  for _, org := range jsonArray {
    fmt.Printf("%s | %d | %s", org["login"], org["id"], org["description"])
  }
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
  apiURL := GetRestApiURL(server, "organizations")
  orgHandler := OrgResponseHandler{}
  DoRestApiCall(apiURL, token, &orgHandler)
}
