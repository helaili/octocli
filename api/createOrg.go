package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

func CreateOrg(name, owner, profileName string) {
  apiURL := client.GetRestApiURL("/admin/organizations")
  params := fmt.Sprintf("{ \"login\": \"%s\", \"admin\": \"%s\", \"profile_name\": \"%s\"}", name, owner, profileName)
  result := client.RestPostForObject(apiURL, params)
  if result != nil {
    fmt.Printf("Congratulations!!! Organization %s was created susccesfully.\n", result["login"])
  }
}
