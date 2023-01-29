package msgpack

type floatElement struct {
	baseElement
	value float64
}

// GetFloat implements Elementer
func (f *floatElement) GetFloat() (float64, error) {
	return f.value, nil
}

func newFloatElement(value float64) *floatElement {
	return &floatElement{
		baseElement: baseElement{
			type_: ElementTypeFloat,
		},
		value: value,
	}
}

var _ Elementer = &floatElement{}
