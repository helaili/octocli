package api

import (
  "fmt"
)

type OrgResponseHandler struct {
}

func (this *OrgResponseHandler) Print(jsonArray []map[string]interface{})  {
  for _, org := range jsonArray {
    fmt.Println(org)
  }
}

func GetAllOrgs(server, token string) {
  apiURL := GetRestApiURL(server, "organizations")
  orgHandler := OrgResponseHandler{}
  DoRestApiCall(apiURL, token, &orgHandler)
}
