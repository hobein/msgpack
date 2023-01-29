package msgpack

type uintElement struct {
	baseElement
	value uint64
}

// GetUint implements Elementer
func (u *uintElement) GetUint() (uint64, error) {
	return u.value, nil
}

func newUintElement(value uint64) *uintElement {
	return &uintElement{
		baseElement: baseElement{
			type_: ElementTypeUint,
		},
		value: value,
	}
}

var _ Elementer = &uintElement{}
