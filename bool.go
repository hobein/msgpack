package msgpack

type boolElement struct {
	baseElement
	value bool
}

// GetBool implements Elementer
func (b *boolElement) GetBool() (bool, error) {
	return b.value, nil
}

func newBoolElement(value bool) *boolElement {
	return &boolElement{
		baseElement: baseElement{
			type_: ElementTypeBool,
		},
		value: value,
	}
}

// NewBoolElement creates a bool element from the given value.
func NewBoolElement(value bool) *boolElement {
	return newBoolElement(value)
}

var _ Elementer = &boolElement{}
