package msgpack

type nilElement struct {
	baseElement
}

func newNilElement() *nilElement {
	return &nilElement{
		baseElement: baseElement{
			type_: ElementTypeNil,
		},
	}
}

// NewNilElement creates a nil element.
func NewNilElement() *nilElement {
	return newNilElement()
}

var _ Elementer = &nilElement{}
