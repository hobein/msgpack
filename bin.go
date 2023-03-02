package msgpack

type binElement struct {
	baseElement
	value []byte
}

// GetBytes implements Elementer
func (b *binElement) GetBytes() ([]byte, error) {
	return b.value, nil
}

func newBinElement(value []byte) *binElement {
	return &binElement{
		baseElement: baseElement{
			type_: ElementTypeBin,
		},
		value: value,
	}
}

// NewBinElement creates a binary element from the given bytes.
func NewBinElement(value []byte) *binElement {
	return newBinElement(value)
}

var _ Elementer = &binElement{}
