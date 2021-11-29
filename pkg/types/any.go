package types

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"
)

func UnmarshalAny(v interface{}) (*anypb.Any, error) {
	switch v := v.(type) {
	case []byte:
		return &anypb.Any{}, nil //TODO add an unmarshal mechanism
	case json.RawMessage:
		return &anypb.Any{}, nil
	default:
		return &anypb.Any{}, fmt.Errorf("%T is not json.RawMessage", v)
	}
}
