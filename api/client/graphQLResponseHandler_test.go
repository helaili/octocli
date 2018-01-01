package client

import (
  "testing"
  "encoding/json"
)

var payload = []byte(`{
  "data": {
    "user": {
      "organizations": {
        "pageInfo": {
          "endCursor": "Y3Vyc29yOnYyOpHOAgk2AQ==",
          "hasNextPage": false
        },
        "nodes": [
          {
            "name": "GitHub"
          },
          {
            "name": "Atom"
          }
        ]
      }
    }
  }
}`)

var nodes = []byte(`[
  {
    "name": "GitHub"
  },
  {
    "name": "Atom"
  }
]`)

func TestGetNodes(t *testing.T) {
  var jsonPayload map[string]interface{}
  json.Unmarshal(payload, &jsonPayload)

  var jsonNodes []interface{}
  json.Unmarshal(nodes, &jsonNodes)

  path := []string{"data", "user", "organizations"}

	cases := []struct {
		jsonPayload map[string]interface{}
    path []string
    want []interface{}
	}{
		{jsonPayload, path, jsonNodes},
	}
	for _, c := range cases {
    gqlrh := BasicGraphQLResponseHandler{}
		got := gqlrh.GetNodes(c.jsonPayload, c.path)
		if !isEqual(got, c.want) {
			t.Errorf("GetNodes(%q, %q) == %q, want %q", c.jsonPayload, c.path, got, c.want)
		}
	}
}

func isEqual(got, want []interface{}) bool {
  if len(got) != len(want) {
    return false
  } else {
    for idx, entry := range got {
      if entry.(map[string]interface{})["name"] != want[idx].(map[string]interface{})["name"] {
        return false
      }
    }
  }
  return true
}
