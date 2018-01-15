package api

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/helaili/octocli/api/client"
)


type TeamMembershipSuccess struct {
  URL string `json:"url"`
  Role string `json:"role"`
  State string `json:"state"`
}

func decodeSuccessfulResponse(resp *http.Response) (success TeamMembershipSuccess, err error) {
  var jsonObj TeamMembershipSuccess

  // Decode the JSON array
  decodeError := json.NewDecoder(resp.Body).Decode(&jsonObj)
  if decodeError != nil {
    return TeamMembershipSuccess{}, decodeError
  } else  {
    return jsonObj, nil
  }
}

func AddTeamMembers(server, token, org, teamSlug string, members []string, role string) {
  // Need to figure out the team ID
  teamObj := GetRestTeam(server, token, org, teamSlug)

  if teamObj != nil && teamObj["id"] != nil {
    for _, member := range members {
      apiURL := client.GetRestApiURL(server, fmt.Sprintf("/teams/%d/memberships/%s", int(teamObj["id"].(float64)), member ))
      resp, err := client.RestQuery(apiURL, token, "PUT", fmt.Sprintf("{\"role\" : \"%s\" }", role), nil)

      if err != nil {
        fmt.Println(err)
      } else if resp.StatusCode != http.StatusOK {
        responseErr, err := client.GetJSONError(resp)
        if err == nil {
          fmt.Printf("Error while adding %s to %s/%s : %s\n", member, org, teamSlug, responseErr.Message)
        }
      } else {
        successMessage, err := decodeSuccessfulResponse(resp)
        if err == nil {
          fmt.Printf("%s was succesfully added to %s/%s as a %s with status %s\n", member, org, teamSlug, successMessage.Role, successMessage.State)
        } else {
          fmt.Printf("Couldn't add user %s: %s\n", member, err)
        }
      }
    }
  } else {
    fmt.Println("Wrong org or team - please use the team's slug as seen in the URL (lowercase, no space...)")
  }
}
