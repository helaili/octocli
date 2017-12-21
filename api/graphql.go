package api

type GraphQLQuery struct {
  Query string `json:"query"`
  Variables interface{} `json:"variables"`
}
