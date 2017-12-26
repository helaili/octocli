package api

import (
  "encoding/json"
)

type ResponseDecoder interface {
  Decode(decoder *json.Decoder) error
}
