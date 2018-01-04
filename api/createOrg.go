package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

func CreateOrg(server, token, name, owner, profileName string) {
  apiURL := client.GetRestApiURL(server, "/admin/organizations")
  params := fmt.Sprintf("{ \"login\": \"%s\", \"admin\": \"%s\", \"profile_name\": \"%s\"}", name, owner, profileName)
  result := client.Post(apiURL, token, params)
  if result != nil {
    fmt.Printf("Congratulations!!! Organization %s was created susccesfully.\n", result["login"])
  }
}
