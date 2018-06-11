package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

func CreateTeam(org, name, description, privacy string) {
  apiURL := client.GetRestApiURL(fmt.Sprintf("/orgs/%s/teams", org))
  params := fmt.Sprintf("{ \"name\": \"%s\", \"description\": \"%s\", \"privacy\": \"%s\"}", name, description, privacy)
  result := client.RestPostForObject(apiURL, params)
  if result != nil {
    fmt.Printf("Congratulations!!! Team %s was created susccesfully with slug %s.\n", result["name"], result["slug"])
  }
}
