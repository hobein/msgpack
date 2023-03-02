package msgpack

type intElement struct {
	baseElement
	value int64
}

// GetInt implements Elementer
func (i *intElement) GetInt() (int64, error) {
	return i.value, nil
}

func newIntElement(value int64) *intElement {
	return &intElement{
		baseElement: baseElement{
			type_: ElementTypeInt,
		},
		value: value,
	}
}

// NewIntElement creates an integer element from the given value.
func NewIntElement(value int64) *intElement {
	return newIntElement(value)
}

var _ Elementer = &intElement{}
