package msgpack

type objectElement struct {
	baseElement
	values []Elementer
}

// AsObject implements Elementer
func (o *objectElement) AsObject() (ObjectWrapper, error) {
	return o, nil
}

// Get implements ObjectElementer
func (o *objectElement) Get(index int) (Elementer, Elementer) {
	_ = o.values[index*2+1] // bounds check hint to compiler

	return o.values[index*2], o.values[index*2+1]
}

// Size implements ObjectElementer
func (o *objectElement) Size() int {
	return len(o.values) / 2
}

func newObjectElement(values []Elementer) *objectElement {
	return &objectElement{
		baseElement: baseElement{
			type_: ElementTypeObject,
		},
		values: values,
	}
}

// NewObjectElement creates an object from the given values.
func NewObjectElement(values []Elementer) *objectElement {
	return newObjectElement(values)
}

var (
	_ Elementer     = &objectElement{}
	_ ObjectWrapper = &objectElement{}
)
