package obj

// NewEntry returns a new *Entry instance initialized with a *Object of the
// given input JSON data.
func NewEntry(inputJSON []byte) (*Entry, error) {
	obj, err := NewObject(inputJSON)
	if err != nil {
		return nil, err
	}

	return &Entry{
		Object: obj,
		Index:  obj.Index(),
	}, err
}

// Entry is a parsed heap item object
type Entry struct {
	Object *Object
	Offset int64
	Index  string
}

// Address returns the Address property of the entry's Object.
func (s *Entry) Address() string {
	return s.Object.Address
}
