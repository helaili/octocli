package api

import (
  "fmt"
  "github.com/helaili/octocli/api/client"
)

var getTeamQuery = `query($login:String!, $slug:String!) {
  organization(login: $login) {
    team(slug: $slug) {
      name
      slug
      id
      description
      privacy
      url
    }
  }
}`

// Get a team with a GraphQL call
func GetTeam(server, token, org, slug string) map[string]interface{} {
  params := map[string]interface{}{"login" : org, "slug" : slug}
  graphQLResponse := client.GraphQLQueryObject(server, token, getTeamQuery, params)
  handler := client.BasicGraphQLResponseHandler{}
  return handler.GetObject(graphQLResponse, []string{"data", "organization", "team"})
}

// Get a team with a Rest call - id through REST is different than the GraphQL id
func GetRestTeam(server, token, org, slug string) map[string]interface{} {
  apiURL := client.GetRestApiURL(server, fmt.Sprintf("/orgs/%s/teams", org))
  teams := client.RestGetForArray(apiURL, token)
  // You can't get one single team, so you need to find the one your're looking for
  for _, team := range teams {
    if team["slug"].(string) == slug {
      return team
    }
  }
  return nil
}
