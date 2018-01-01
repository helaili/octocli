package client

type GraphQLResponseHandler interface {
  ResponseHandler
  TableRows(jsonObj map[string]interface{}) [][]string
  PageInfoPath() []string
}

type BasicGraphQLResponseHandler struct {
}

// Navigate the JSON response to retrive the 'nodes' array
func (this *BasicGraphQLResponseHandler) GetNodes(jsonObj map[string]interface{}, path []string) ([]interface{}) {
  if(len(path) == 0) {
    return jsonObj["nodes"].([]interface{})
  } else {
    return this.GetNodes(jsonObj[path[0]].(map[string]interface{}), path[1:])
  }
}
