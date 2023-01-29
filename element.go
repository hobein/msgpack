package msgpack

//go:generate stringer -type=ElementType -trimprefix ElementType

type ElementType uint8

const (
	// ElementTypeNil nil value
	ElementTypeNil ElementType = iota + 1
	// ElementTypeBool bool value
	ElementTypeBool
	// ElementTypeFloat float value
	ElementTypeFloat
	// ElementTypeInt int value
	ElementTypeInt
	// ElementTypeUint uint value
	ElementTypeUint
	// ElementTypeString string value
	ElementTypeString
	// ElementTypeObject object value
	ElementTypeObject
	// ElementTypeArray array value
	ElementTypeArray
	// ElementTypeBin binary value
	ElementTypeBin
	// ElementTypeExtension extension value
	ElementTypeExtension
)

type Elementer interface {
	// Type returns the type of the underlying value
	Type() ElementType
	// GetBool returns the boolean value or an error if the
	// underlying value is not bool.
	GetBool() (bool, error)
	// GetFloat returns the float64 value or an error if the
	// underlying value is not float64.
	GetFloat() (float64, error)
	// GetInt returns the int64 value or an error if the
	// underlying value is not int64.
	GetInt() (int64, error)
	// GetUint returns the uint64 value or an error if the
	// underlying value is not uint64.
	GetUint() (uint64, error)
	// GetString returns the string value or an error if the
	// underlying value is not string.
	GetString() (string, error)
	// GetBytes returns the byte array value or an error if the
	// underlying value is not []byte.
	GetBytes() ([]byte, error)
	// GetExtensionRaw returns raw msgpack encoded value of
	// the underlying extension or an error if the
	// underlying value is not extension.
	GetExtensionRaw() ([]byte, error)
	// AsArray returns an array wrapper if the underlying
	// value is an array.
	AsArray() (ArrayWrapper, error)
	// AsObject returns an object wrapper if the underlying
	// value is an object.
	AsObject() (ObjectWrapper, error)
}

// ArrayWrapper exposes methods to access the items of an array.
type ArrayWrapper interface {
	// Size returns the size of the array.
	Size() int
	// Get returns the element at the index postion.
	Get(index int) Elementer
}

// ObjectWrapper exposes methods to access the items of an object.
type ObjectWrapper interface {
	// Size returns the number of key of the object.
	Size() int
	// Get returns the key and value element at the index position.
	Get(index int) (Elementer, Elementer)
}
