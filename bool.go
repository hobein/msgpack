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

var _ Elementer = &boolElement{}
