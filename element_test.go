package msgpack

import (
	"bytes"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
)

func TestElement_unmarshalMsgPack(t *testing.T) {
	/*
		{
		  "int": 1,
		  "float": 0.5,
		  "boolean": true,
		  "null": null,
		  "string": "foo bar",
		  "array": [
		    "foo",
		    "bar"
		  ],
		  "object": {
		    "foo": 1,
		    "baz": 0.5
		  }
		}
	*/
	payload := []byte{135, 163, 105, 110, 116, 1, 165, 102, 108, 111, 97, 116, 203, 63, 224, 0, 0, 0, 0, 0, 0, 167, 98, 111, 111, 108, 101, 97, 110, 195, 164, 110, 117, 108, 108, 192, 166, 115, 116, 114, 105, 110, 103, 167, 102, 111, 111, 32, 98, 97, 114, 165, 97, 114, 114, 97, 121, 146, 163, 102, 111, 111, 163, 98, 97, 114, 166, 111, 98, 106, 101, 99, 116, 130, 163, 102, 111, 111, 1, 163, 98, 97, 122, 203, 63, 224, 0, 0, 0, 0, 0, 0}
	dec := msgpack.NewDecoder(bytes.NewBuffer(payload))

	elt, err := UnmarshalFromDecoder(dec)
	if err != nil {
		t.Fatalf("fail to unmarshal msgpack. %v", err)
	}

	t.Logf("element type %s", elt.Type())

	want := ElementTypeObject
	got := elt.Type()
	if got != want {
		t.Fatalf("element type missmatch. want %v, got %v", want, got)
	}

	{
		v, err := elt.GetBool()
		if err == nil {
			t.Fatalf("GetBool should return an error")
		}
		if v != false {
			t.Fatalf("GetBool should false and an error")
		}
	}
	{
		v, err := elt.GetFloat()
		if err == nil {
			t.Fatalf("GetFloat should return an error")
		}
		if v != 0 {
			t.Fatalf("GetFloat should 0 and an error")
		}
	}
	{
		v, err := elt.GetInt()
		if err == nil {
			t.Fatalf("GetInt should return an error")
		}
		if v != 0 {
			t.Fatalf("GetInt should 0 and an error")
		}
	}
	{
		v, err := elt.GetUint()
		if err == nil {
			t.Fatalf("GetUint should return an error")
		}
		if v != 0 {
			t.Fatalf("GetUint should 0 and an error")
		}
	}
	{
		v, err := elt.GetString()
		if err == nil {
			t.Fatalf("GetString should return an error")
		}
		if v != "" {
			t.Fatalf("GetString should \"\" and an error")
		}
	}
	{
		v, err := elt.GetBytes()
		if err == nil {
			t.Fatalf("GetBytes should return an error")
		}
		if v != nil {
			t.Fatalf("GetBytes should nil and an error")
		}
	}
	{
		v, err := elt.GetExtensionRaw()
		if err == nil {
			t.Fatalf("GetExtensionRaw should return an error")
		}
		if v != nil {
			t.Fatalf("GetExtensionRaw should nil and an error")
		}
	}
	{
		a, err := elt.AsArray()
		if err == nil {
			t.Fatalf("AsArray should return an error")
		}
		if a != nil {
			t.Fatalf("AsArray should nil and an error")
		}
	}

	o, err := elt.AsObject()
	if err != nil {
		t.Fatalf("AsObject should not returned an error. %v", err)
	}
	if want, got := 7, o.Size(); want != got {
		t.Fatalf("Size missmatch. want %d, got %d", want, got)
	}

	t.Logf("element type %s", elt.Type())

	for i := 0; i < o.Size(); i++ {
		k, v := o.Get(i)
		if want, got := ElementTypeString, k.Type(); want != got {
			t.Fatalf("key #%d type missmatch. want %s, got %s", i, want, got)
		}
		switch i {
		case 0:
			if want, got := ElementTypeInt, v.Type(); want != got {
				t.Fatalf("value #%d type missmatch. want %s, got %s", i, want, got)
			}
			w, err := v.GetInt()
			if err != nil {
				t.Fatalf("cannot call getInt on a %s element. %v", v.Type(), err)
			}
			if want, got := int64(1), w; want != got {
				t.Fatalf("int value missmatch. want %v, got %v", want, got)
			}
		case 1:
			if want, got := ElementTypeFloat, v.Type(); want != got {
				t.Fatalf("value #%d type missmatch. want %s, got %s", i, want, got)
			}
			w, err := v.GetFloat()
			if err != nil {
				t.Fatalf("cannot call getInt on a %s element. %v", v.Type(), err)
			}
			if want, got := float64(0.5), w; want != got {
				t.Fatalf("float value missmatch. want %v, got %v", want, got)
			}
		}
	}
}

func TestElement_nil(t *testing.T) {
	b := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(b)
	assert.Nil(t, enc.EncodeNil())

	dec := msgpack.NewDecoder(b)
	elt, err := UnmarshalFromDecoder(dec)
	assert.Nil(t, err)
	assert.Equal(t, elt.Type(), ElementTypeNil)

	{
		_, err := elt.GetBool()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetFloat()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetInt()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetUint()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetString()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetBytes()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetExtensionRaw()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsArray()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsObject()
		assert.NotNil(t, err)
	}
}

func TestElement_bool(t *testing.T) {
	for _, v := range []bool{false, true} {
		b := bytes.NewBuffer(nil)
		enc := msgpack.NewEncoder(b)
		assert.Nil(t, enc.EncodeBool(v))

		dec := msgpack.NewDecoder(b)
		elt, err := UnmarshalFromDecoder(dec)

		assert.Nil(t, err)
		assert.Equal(t, elt.Type(), ElementTypeBool)

		{
			value, err := elt.GetBool()
			assert.Nil(t, err)
			assert.Equal(t, v, value)
		}
		{
			_, err := elt.GetFloat()
			assert.NotNil(t, err)
		}
		{
			_, err := elt.GetInt()
			assert.NotNil(t, err)
		}
		{
			_, err := elt.GetUint()
			assert.NotNil(t, err)
		}
		{
			_, err := elt.GetString()
			assert.NotNil(t, err)
		}
		{
			_, err := elt.GetBytes()
			assert.NotNil(t, err)
		}
		{
			_, err := elt.GetExtensionRaw()
			assert.NotNil(t, err)
		}
		{
			_, err := elt.AsArray()
			assert.NotNil(t, err)
		}
		{
			_, err := elt.AsObject()
			assert.NotNil(t, err)
		}
	}
}

func TestElement_float(t *testing.T) {
	b := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(b)
	assert.Nil(t, enc.EncodeFloat32(12.4))

	dec := msgpack.NewDecoder(b)
	elt, err := UnmarshalFromDecoder(dec)

	assert.Nil(t, err)
	assert.Equal(t, elt.Type(), ElementTypeFloat)

	{
		_, err := elt.GetBool()
		assert.NotNil(t, err)
	}
	{
		value, err := elt.GetFloat()
		assert.Nil(t, err)
		assert.InDelta(t, value, float64(12.4), 1e-6)
	}
	{
		_, err := elt.GetInt()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetUint()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetString()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetBytes()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetExtensionRaw()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsArray()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsObject()
		assert.NotNil(t, err)
	}
}

func TestElement_int(t *testing.T) {
	b := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(b)
	assert.Nil(t, enc.EncodeInt64(5734057394753894))

	dec := msgpack.NewDecoder(b)
	elt, err := UnmarshalFromDecoder(dec)

	assert.Nil(t, err)
	assert.Equal(t, elt.Type(), ElementTypeInt)

	{
		_, err := elt.GetBool()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetFloat()
		assert.NotNil(t, err)

	}
	{
		value, err := elt.GetInt()
		assert.Nil(t, err)
		assert.Equal(t, value, int64(5734057394753894))
	}
	{
		_, err := elt.GetUint()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetString()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetBytes()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetExtensionRaw()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsArray()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsObject()
		assert.NotNil(t, err)
	}
}

func TestElement_uint(t *testing.T) {
	b := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(b)
	assert.Nil(t, enc.EncodeUint64(5734057394753894))

	dec := msgpack.NewDecoder(b)
	elt, err := UnmarshalFromDecoder(dec)

	assert.Nil(t, err)
	assert.Equal(t, elt.Type(), ElementTypeUint)

	{
		_, err := elt.GetBool()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetFloat()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetInt()
		assert.NotNil(t, err)
	}
	{
		value, err := elt.GetUint()
		assert.Nil(t, err)
		assert.Equal(t, value, uint64(5734057394753894))

	}
	{
		_, err := elt.GetString()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetBytes()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetExtensionRaw()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsArray()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsObject()
		assert.NotNil(t, err)
	}
}

func TestElement_string(t *testing.T) {
	const str = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam massa tellus, aliquam eu dapibus at."
	b := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(b)
	assert.Nil(t, enc.EncodeString(str))

	dec := msgpack.NewDecoder(b)
	elt, err := UnmarshalFromDecoder(dec)

	assert.Nil(t, err)
	assert.Equal(t, elt.Type(), ElementTypeString)

	{
		_, err := elt.GetBool()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetFloat()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetInt()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetUint()
		assert.NotNil(t, err)
	}
	{
		value, err := elt.GetString()
		assert.Nil(t, err)
		assert.Equal(t, str, value)
	}
	{
		_, err := elt.GetBytes()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetExtensionRaw()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsArray()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsObject()
		assert.NotNil(t, err)
	}
}

func TestElement_object(t *testing.T) {
	o := map[string]interface{}{
		"a": "hello",
		"b": "world",
	}
	b := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(b)
	assert.Nil(t, enc.EncodeMap(o))

	dec := msgpack.NewDecoder(b)
	elt, err := UnmarshalFromDecoder(dec)

	assert.Nil(t, err)
	assert.Equal(t, elt.Type(), ElementTypeObject)

	{
		_, err := elt.GetBool()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetFloat()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetInt()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetUint()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetString()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetBytes()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetExtensionRaw()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsArray()
		assert.NotNil(t, err)
	}
	{
		value, err := elt.AsObject()
		assert.Nil(t, err)
		assert.Equal(t, 2, value.Size())
	}
}

func TestElement_bin(t *testing.T) {
	byteArray := make([]byte, 514)
	{
		_, err := rand.Read(byteArray)
		assert.Nil(t, err)
	}
	b := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(b)
	assert.Nil(t, enc.EncodeBytes(byteArray))

	dec := msgpack.NewDecoder(b)
	elt, err := UnmarshalFromDecoder(dec)

	assert.Nil(t, err)
	assert.Equal(t, elt.Type(), ElementTypeBin)

	{
		_, err := elt.GetBool()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetFloat()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetInt()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetUint()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetString()
		assert.NotNil(t, err)
	}
	{
		value, err := elt.GetBytes()
		assert.Nil(t, err)
		assert.Equal(t, value, byteArray)
	}
	{
		_, err := elt.GetExtensionRaw()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsArray()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsObject()
		assert.NotNil(t, err)
	}
}

func TestElement_extension(t *testing.T) {
	now := time.Now()
	b := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(b)
	assert.Nil(t, enc.EncodeTime(now))

	dec := msgpack.NewDecoder(b)
	elt, err := UnmarshalFromDecoder(dec)

	assert.Nil(t, err)
	assert.Equal(t, elt.Type(), ElementTypeExtension)

	{
		_, err := elt.GetBool()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetFloat()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetInt()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetUint()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetString()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetBytes()
		assert.NotNil(t, err)
	}
	{
		extRaw, err := elt.GetExtensionRaw()
		assert.Nil(t, err)
		dec.Reset(bytes.NewBuffer(extRaw))
		decodedTime, err := dec.DecodeTime()
		assert.Nil(t, err)
		assert.Equal(t, now.UnixNano(), decodedTime.UnixNano())
	}
	{
		_, err := elt.AsArray()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.AsObject()
		assert.NotNil(t, err)
	}
}

func TestElement_array(t *testing.T) {
	byteArray := make([]byte, 514)
	{
		_, err := rand.Read(byteArray)
		assert.Nil(t, err)
	}
	b := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(b)
	assert.Nil(t, enc.EncodeArrayLen(len(byteArray)))
	for _, c := range byteArray {
		assert.Nil(t, enc.EncodeUint8(c))
	}

	dec := msgpack.NewDecoder(b)
	elt, err := UnmarshalFromDecoder(dec)

	assert.Nil(t, err)
	assert.Equal(t, elt.Type(), ElementTypeArray)

	{
		_, err := elt.GetBool()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetFloat()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetInt()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetUint()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetString()
		assert.NotNil(t, err)
	}
	{
		_, err := elt.GetBytes()
		assert.NotNil(t, err)
	}
	{
		value, err := elt.AsArray()
		assert.Nil(t, err)
		assert.Equal(t, len(byteArray), value.Size())
		for i := 0; i < value.Size(); i++ {
			c := value.Get(i)
			assert.Equal(t, c.Type(), ElementTypeUint)
			v, err := c.GetUint()
			assert.Nil(t, err)
			assert.Equal(t, uint8(v), byteArray[i])
		}
	}
	{
		_, err := elt.AsObject()
		assert.NotNil(t, err)
	}
}
