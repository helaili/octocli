package api

type ResponseHandler interface {
}

type GraphQLResponseHandler interface {
  ResponseHandler
  Print(jsonObj map[string]interface{})
  PageInfoPath() []string
}

type RestResponseHandler interface {
  ResponseHandler
  Print(jsonArray[]map[string]interface{})
}
