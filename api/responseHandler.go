package api

type ResponseHandler interface {
  TableHeader() []string
}

type GraphQLResponseHandler interface {
  ResponseHandler
  Print(jsonObj map[string]interface{})
  TableRows(jsonObj map[string]interface{}) [][]string
  PageInfoPath() []string
}

type RestResponseHandler interface {
  ResponseHandler
  Print(jsonArray []map[string]interface{})
  TableRows(jsonArray []map[string]interface{}) [][]string
}
