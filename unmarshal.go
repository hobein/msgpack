package msgpack

import (
	"bytes"
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
)

// Unmarshal decodes the MessagePack-encoded data and returns an interface Elementer.
func Unmarshal(data []byte) (Elementer, error) {
	dec := msgpack.GetDecoder()
	defer msgpack.PutDecoder(dec)
	dec.Reset(bytes.NewBuffer(data))
	return UnmarshalFromDecoder(dec)
}

// UnmarshalFromDecoder consumes the decoder and returns an interface Elementer.
func UnmarshalFromDecoder(dec *msgpack.Decoder) (Elementer, error) {
	code, err := dec.PeekCode()
	if err != nil {
		return nil, err
	}

	switch code {
	case msgpcode.Nil:
		err := dec.DecodeNil()
		if err != nil {
			return nil, err
		}
		return newNilElement(), nil

	case msgpcode.False, msgpcode.True:
		value, err := dec.DecodeBool()
		if err != nil {
			return nil, err
		}

		return newBoolElement(value), nil

	case msgpcode.Float, msgpcode.Double:
		value, err := dec.DecodeFloat64()
		if err != nil {
			return nil, err
		}

		return newFloatElement(value), nil

	case msgpcode.Uint8, msgpcode.Uint16, msgpcode.Uint32, msgpcode.Uint64:
		value, err := dec.DecodeUint64()
		if err != nil {
			return nil, err
		}

		return newUintElement(value), nil

	case msgpcode.Int8, msgpcode.Int16, msgpcode.Int32, msgpcode.Int64:
		value, err := dec.DecodeInt64()
		if err != nil {
			return nil, err
		}

		return newIntElement(value), nil
	}

	if msgpcode.IsFixedNum(code) {
		value, err := dec.DecodeInt64()
		if err != nil {
			return nil, err
		}

		return newIntElement(value), nil
	}

	if msgpcode.IsString(code) {
		value, err := dec.DecodeString()
		if err != nil {
			return nil, err
		}

		return newStringElement(value), nil
	}

	if msgpcode.IsBin(code) {
		value, err := dec.DecodeBytes()
		if err != nil {
			return nil, err
		}

		return newBinElement(value), nil
	}

	if isArray(code) {
		size, err := dec.DecodeArrayLen()
		if err != nil {
			return nil, err
		}

		arr := make([]Elementer, 0, size)

		for i := 0; i < size; i++ {
			item, err := UnmarshalFromDecoder(dec)
			if err != nil {
				return nil, err
			}
			arr = append(arr, item)
		}

		return newArrayElement(arr), nil
	}

	if isMap(code) {
		size, err := dec.DecodeMapLen()
		if err != nil {
			return nil, err
		}

		arr := make([]Elementer, 0, 2*size)

		for i := 0; i < size; i++ {
			keyItem, err := UnmarshalFromDecoder(dec)
			if err != nil {
				return nil, err
			}
			arr = append(arr, keyItem)

			valueItem, err := UnmarshalFromDecoder(dec)
			if err != nil {
				return nil, err
			}
			arr = append(arr, valueItem)
		}

		return newObjectElement(arr), nil
	}

	if msgpcode.IsExt(code) {
		value, err := dec.DecodeRaw()
		if err != nil {
			return nil, err
		}

		return newExtElement(value), nil
	}

	return nil, fmt.Errorf("msgpack: unknown code %x", code)
}

func isMap(code byte) bool {
	switch code {
	case msgpcode.Map16, msgpcode.Map32:
		return true
	}

	return msgpcode.IsFixedMap(code)
}

func isArray(code byte) bool {
	switch code {
	case msgpcode.Array16, msgpcode.Array32:
		return true
	}

	return msgpcode.IsFixedArray(code)
}
