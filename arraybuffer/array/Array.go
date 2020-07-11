package array

import (
	"errors"
	"syscall/js"
)

var ERR_ARRAY_UNSUPPORTED = errors.New("Array unsupported")

type ArrayType string

func (self ArrayType) String() string {
	return string(self)
}

const (
	ArrayTypeArray             ArrayType = "Array"
	ArrayTypeInt8Array         ArrayType = "Int8Array"
	ArrayTypeUint8Array        ArrayType = "Uint8Array"
	ArrayTypeUint8ClampedArray ArrayType = "Uint8ClampedArray"
	ArrayTypeInt16Array        ArrayType = "Int16Array"
	ArrayTypeUint16Array       ArrayType = "Uint16Array"
	ArrayTypeInt32Array        ArrayType = "Int32Array"
	ArrayTypeUint32Array       ArrayType = "Uint32Array"
	ArrayTypeFloat32Array      ArrayType = "Float32Array"
	ArrayTypeFloat64Array      ArrayType = "Float64Array"
	ArrayTypeBigInt64Array     ArrayType = "BigInt64Array"
	ArrayTypeBigUint64Array    ArrayType = "BigUint64Array"
)

var ArrayTypes = []ArrayType{
	ArrayTypeArray,
	ArrayTypeInt8Array,
	ArrayTypeUint8Array,
	ArrayTypeUint8ClampedArray,
	ArrayTypeInt16Array,
	ArrayTypeUint16Array,
	ArrayTypeInt32Array,
	ArrayTypeUint32Array,
	ArrayTypeFloat32Array,
	ArrayTypeFloat64Array,
	ArrayTypeBigInt64Array,
	ArrayTypeBigUint64Array,
}

type Array struct {
	JSValue js.Value
}

func NewArray(array_type ArrayType, value js.Value) (*Array, error) {
	found := false
	for _, i := range ArrayTypes {
		if i == array_type {
			found = true
			break
		}
	}
	if !found {
		return nil, errors("Invalid array type name")
	}
	
	js.
}
