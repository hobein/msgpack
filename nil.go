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

var _ Elementer = &nilElement{}
