package obj

import "encoding/json"

//go:generate easyjson -all object.go

// NewObject returns a new *Object instance with it's attributes populated from
// the given input JSON data.
func NewObject(inputJSON []byte) (*Object, error) {
	var obj Object
	err := json.Unmarshal(inputJSON, &obj)

	return &obj, err
}

// Object is a minimal representation of a Ruby heap object as exported from
// Ruby via `ObjectSpace.dump_all`.
type Object struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

// Index returns a unique index for the given Object.
func (s *Object) Index() string {
	return s.Address + ":" + s.Type
}
