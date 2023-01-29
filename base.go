package msgpack

type baseElement struct {
	type_ ElementType
}

// Type implements Elementer
func (b *baseElement) Type() ElementType {
	return b.type_
}

// GetBool implements Elementer
func (b *baseElement) GetBool() (bool, error) {
	const want = ElementTypeBool
	return false, &TypeError{want: want, got: b.type_}
}

// GetFloat implements Elementer
func (b *baseElement) GetFloat() (float64, error) {
	const want = ElementTypeFloat
	return 0, &TypeError{want: want, got: b.type_}
}

// GetInt implements Elementer
func (b *baseElement) GetInt() (int64, error) {
	const want = ElementTypeInt
	return 0, &TypeError{want: want, got: b.type_}
}

// GetUint implements Elementer
func (b *baseElement) GetUint() (uint64, error) {
	const want = ElementTypeUint
	return 0, &TypeError{want: want, got: b.type_}
}

// GetString implements Elementer
func (b *baseElement) GetString() (string, error) {
	const want = ElementTypeString
	return "", &TypeError{want: want, got: b.type_}
}

// GetBytes implements Elementer
func (b *baseElement) GetBytes() ([]byte, error) {
	const want = ElementTypeBin
	return nil, &TypeError{want: want, got: b.type_}
}

// GetExtensionRaw implements Elementer
func (b *baseElement) GetExtensionRaw() ([]byte, error) {
	const want = ElementTypeExtension
	return nil, &TypeError{want: want, got: b.type_}
}

// AsArray implements Elementer
func (b *baseElement) AsArray() (ArrayWrapper, error) {
	const want = ElementTypeArray
	return nil, &TypeError{want: want, got: b.type_}
}

// AsObject implements Elementer
func (b *baseElement) AsObject() (ObjectWrapper, error) {
	const want = ElementTypeObject
	return nil, &TypeError{want: want, got: b.type_}
}
