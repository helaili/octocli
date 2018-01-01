package client

type RestResponseHandler interface {
  ResponseHandler
  TableRows(jsonArray []map[string]interface{}) [][]string
}
