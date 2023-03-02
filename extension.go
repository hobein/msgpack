package msgpack

type extElement struct {
	baseElement
	value []byte
}

// GetExtensionRaw implements Elementer
func (e *extElement) GetExtensionRaw() ([]byte, error) {
	return e.value, nil
}

func newExtElement(value []byte) *extElement {
	return &extElement{
		baseElement: baseElement{
			type_: ElementTypeExtension,
		},
		value: value,
	}
}

// NewExtElement creates an extension element from the given value.
func NewExtElement(value []byte) *extElement {
	return newExtElement(value)
}

var _ Elementer = &extElement{}
