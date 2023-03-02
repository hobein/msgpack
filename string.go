package msgpack

type stringElement struct {
	baseElement
	value string
}

// GetString implements Elementer
func (s *stringElement) GetString() (string, error) {
	return s.value, nil
}

func newStringElement(value string) *stringElement {
	return &stringElement{
		baseElement: baseElement{
			type_: ElementTypeString,
		},
		value: value,
	}
}

// NewStringElement creates a string element from the given value.
func NewStringElement(value string) *stringElement {
	return newStringElement(value)
}

var _ Elementer = &stringElement{}
