package msgpack

type arrayElement struct {
	baseElement
	values []Elementer
}

// AsArray implements Elementer
func (a *arrayElement) AsArray() (ArrayWrapper, error) {
	return a, nil
}

// Get implements ArrayElementer
func (a *arrayElement) Get(index int) Elementer {
	return a.values[index]
}

// Size implements ArrayElementer
func (a *arrayElement) Size() int {
	return len(a.values)
}

func newArrayElement(values []Elementer) *arrayElement {
	return &arrayElement{
		baseElement: baseElement{
			type_: ElementTypeArray,
		},
		values: values,
	}
}

// NewArrayElement creates an array from the given values.
func NewArrayElement(values []Elementer) *arrayElement {
	return newArrayElement(values)
}

var (
	_ Elementer    = &arrayElement{}
	_ ArrayWrapper = &arrayElement{}
)
