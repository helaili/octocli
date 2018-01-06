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
    fmt.Printf("Error while decoding the error object: %s\n", decodeError)
    return TeamMembershipSuccess{}, decodeError
  } else  {
    return jsonObj, nil
  }
}

func AddTeamMembers(server, token, org, teamSlug string, members []string) {
  // Need to figure out the team ID
  teamObj := GetRestTeam(server, token, org, teamSlug)

  if teamObj != nil && teamObj["id"] != nil {
    for _, member := range members {
      apiURL := client.GetRestApiURL(server, fmt.Sprintf("/teams/%d/memberships/%s", int(teamObj["id"].(float64)), member ))
      resp, err := client.RestQuery(apiURL, token, "PUT", "{\"role\" : \"member\" }", nil)

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
          fmt.Printf("%s was succesfully added to %s/%s with status %s\n", member, org, teamSlug, successMessage.State)
        }
      }
    }
  }
}
