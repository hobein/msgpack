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

var _ Elementer = &binElement{}
