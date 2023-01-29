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

var _ Elementer = &extElement{}
