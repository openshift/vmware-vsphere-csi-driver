// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package attribute // import "go.opentelemetry.io/otel/attribute"

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"go.opentelemetry.io/otel/internal"
	"go.opentelemetry.io/otel/internal/attribute"
)

//go:generate stringer -type=Type

// Type describes the type of the data Value holds.
type Type int

// Value represents the value part in key-value pairs.
type Value struct {
	vtype    Type
	numeric  uint64
	stringly string
	// TODO Lazy value type?

	array interface{}
}

const (
	// INVALID is used for a Value with no value set.
	INVALID Type = iota
	// BOOL is a boolean Type Value.
	BOOL
	// INT64 is a 64-bit signed integral Type Value.
	INT64
	// FLOAT64 is a 64-bit floating point Type Value.
	FLOAT64
	// STRING is a string Type Value.
	STRING
	// ARRAY is an array Type Value used to store 1-dimensional slices or
	// arrays of bool, int, int32, int64, uint, uint32, uint64, float,
	// float32, float64, or string types.
	ARRAY
)

// BoolValue creates a BOOL Value.
func BoolValue(v bool) Value {
	return Value{
		vtype:   BOOL,
		numeric: internal.BoolToRaw(v),
	}
}

<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// BoolSliceValue creates a BOOLSLICE Value.
func BoolSliceValue(v []bool) Value {
	cp := make([]bool, len(v))
	copy(cp, v)
	return Value{
		vtype: BOOLSLICE,
		slice: &cp,
	}
}

// IntValue creates an INT64 Value.
func IntValue(v int) Value {
	return Int64Value(int64(v))
}

// IntSliceValue creates an INTSLICE Value.
func IntSliceValue(v []int) Value {
	cp := make([]int64, 0, len(v))
	for _, i := range v {
		cp = append(cp, int64(i))
	}
	return Value{
		vtype: INT64SLICE,
		slice: &cp,
	}
}

=======
// BoolSliceValue creates a BOOLSLICE Value.
func BoolSliceValue(v []bool) Value {
	return Value{vtype: BOOLSLICE, slice: attribute.BoolSliceValue(v)}
}

// IntValue creates an INT64 Value.
func IntValue(v int) Value {
	return Int64Value(int64(v))
}

// IntSliceValue creates an INTSLICE Value.
func IntSliceValue(v []int) Value {
	var int64Val int64
	cp := reflect.New(reflect.ArrayOf(len(v), reflect.TypeOf(int64Val)))
	for i, val := range v {
		cp.Elem().Index(i).SetInt(int64(val))
	}
	return Value{
		vtype: INT64SLICE,
		slice: cp.Elem().Interface(),
	}
}

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// Int64Value creates an INT64 Value.
func Int64Value(v int64) Value {
	return Value{
		vtype:   INT64,
		numeric: internal.Int64ToRaw(v),
	}
}

<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// Int64SliceValue creates an INT64SLICE Value.
func Int64SliceValue(v []int64) Value {
	cp := make([]int64, len(v))
	copy(cp, v)
	return Value{
		vtype: INT64SLICE,
		slice: &cp,
	}
}

=======
// Int64SliceValue creates an INT64SLICE Value.
func Int64SliceValue(v []int64) Value {
	return Value{vtype: INT64SLICE, slice: attribute.Int64SliceValue(v)}
}

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// Float64Value creates a FLOAT64 Value.
func Float64Value(v float64) Value {
	return Value{
		vtype:   FLOAT64,
		numeric: internal.Float64ToRaw(v),
	}
}

<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// Float64SliceValue creates a FLOAT64SLICE Value.
func Float64SliceValue(v []float64) Value {
	cp := make([]float64, len(v))
	copy(cp, v)
	return Value{
		vtype: FLOAT64SLICE,
		slice: &cp,
	}
}

=======
// Float64SliceValue creates a FLOAT64SLICE Value.
func Float64SliceValue(v []float64) Value {
	return Value{vtype: FLOAT64SLICE, slice: attribute.Float64SliceValue(v)}
}

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// StringValue creates a STRING Value.
func StringValue(v string) Value {
	return Value{
		vtype:    STRING,
		stringly: v,
	}
}

<<<<<<< HEAD
// IntValue creates an INT64 Value.
func IntValue(v int) Value {
	return Int64Value(int64(v))
}

// ArrayValue creates an ARRAY value from an array or slice.
// Only arrays or slices of bool, int, int64, float, float64, or string types are allowed.
// Specifically, arrays  and slices can not contain other arrays, slices, structs, or non-standard
// types. If the passed value is not an array or slice of these types an
// INVALID value is returned.
func ArrayValue(v interface{}) Value {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Slice:
		// get array type regardless of dimensions
		typ := reflect.TypeOf(v).Elem()
		kind := typ.Kind()
		switch kind {
		case reflect.Bool, reflect.Int, reflect.Int64,
			reflect.Float64, reflect.String:
			val := reflect.ValueOf(v)
			length := val.Len()
			frozen := reflect.Indirect(reflect.New(reflect.ArrayOf(length, typ)))
			reflect.Copy(frozen, val)
			return Value{
				vtype: ARRAY,
				array: frozen.Interface(),
			}
		default:
			return Value{vtype: INVALID}
		}
	}
	return Value{vtype: INVALID}
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// StringSliceValue creates a STRINGSLICE Value.
func StringSliceValue(v []string) Value {
	cp := make([]string, len(v))
	copy(cp, v)
	return Value{
		vtype: STRINGSLICE,
		slice: &cp,
	}
=======
// StringSliceValue creates a STRINGSLICE Value.
func StringSliceValue(v []string) Value {
	return Value{vtype: STRINGSLICE, slice: attribute.StringSliceValue(v)}
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// Type returns a type of the Value.
func (v Value) Type() Type {
	return v.vtype
}

// AsBool returns the bool value. Make sure that the Value's type is
// BOOL.
func (v Value) AsBool() bool {
	return internal.RawToBool(v.numeric)
}

<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// AsBoolSlice returns the []bool value. Make sure that the Value's type is
// BOOLSLICE.
func (v Value) AsBoolSlice() []bool {
	if s, ok := v.slice.(*[]bool); ok {
		return *s
	}
	return nil
}

=======
// AsBoolSlice returns the []bool value. Make sure that the Value's type is
// BOOLSLICE.
func (v Value) AsBoolSlice() []bool {
	if v.vtype != BOOLSLICE {
		return nil
	}
	return v.asBoolSlice()
}

func (v Value) asBoolSlice() []bool {
	return attribute.AsBoolSlice(v.slice)
}

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// AsInt64 returns the int64 value. Make sure that the Value's type is
// INT64.
func (v Value) AsInt64() int64 {
	return internal.RawToInt64(v.numeric)
}

<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// AsInt64Slice returns the []int64 value. Make sure that the Value's type is
// INT64SLICE.
func (v Value) AsInt64Slice() []int64 {
	if s, ok := v.slice.(*[]int64); ok {
		return *s
	}
	return nil
}

=======
// AsInt64Slice returns the []int64 value. Make sure that the Value's type is
// INT64SLICE.
func (v Value) AsInt64Slice() []int64 {
	if v.vtype != INT64SLICE {
		return nil
	}
	return v.asInt64Slice()
}

func (v Value) asInt64Slice() []int64 {
	return attribute.AsInt64Slice(v.slice)
}

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// AsFloat64 returns the float64 value. Make sure that the Value's
// type is FLOAT64.
func (v Value) AsFloat64() float64 {
	return internal.RawToFloat64(v.numeric)
}

<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// AsFloat64Slice returns the []float64 value. Make sure that the Value's type is
// FLOAT64SLICE.
func (v Value) AsFloat64Slice() []float64 {
	if s, ok := v.slice.(*[]float64); ok {
		return *s
	}
	return nil
}

=======
// AsFloat64Slice returns the []float64 value. Make sure that the Value's type is
// FLOAT64SLICE.
func (v Value) AsFloat64Slice() []float64 {
	if v.vtype != FLOAT64SLICE {
		return nil
	}
	return v.asFloat64Slice()
}

func (v Value) asFloat64Slice() []float64 {
	return attribute.AsFloat64Slice(v.slice)
}

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// AsString returns the string value. Make sure that the Value's type
// is STRING.
func (v Value) AsString() string {
	return v.stringly
}

<<<<<<< HEAD
// AsArray returns the array Value as an interface{}.
func (v Value) AsArray() interface{} {
	return v.array
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// AsStringSlice returns the []string value. Make sure that the Value's type is
// STRINGSLICE.
func (v Value) AsStringSlice() []string {
	if s, ok := v.slice.(*[]string); ok {
		return *s
	}
	return nil
=======
// AsStringSlice returns the []string value. Make sure that the Value's type is
// STRINGSLICE.
func (v Value) AsStringSlice() []string {
	if v.vtype != STRINGSLICE {
		return nil
	}
	return v.asStringSlice()
}

func (v Value) asStringSlice() []string {
	return attribute.AsStringSlice(v.slice)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

type unknownValueType struct{}

// AsInterface returns Value's data as interface{}.
func (v Value) AsInterface() interface{} {
	switch v.Type() {
	case ARRAY:
		return v.AsArray()
	case BOOL:
		return v.AsBool()
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case BOOLSLICE:
		return v.AsBoolSlice()
=======
	case BOOLSLICE:
		return v.asBoolSlice()
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case INT64:
		return v.AsInt64()
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case INT64SLICE:
		return v.AsInt64Slice()
=======
	case INT64SLICE:
		return v.asInt64Slice()
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case FLOAT64:
		return v.AsFloat64()
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case FLOAT64SLICE:
		return v.AsFloat64Slice()
=======
	case FLOAT64SLICE:
		return v.asFloat64Slice()
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case STRING:
		return v.stringly
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case STRINGSLICE:
		return v.AsStringSlice()
=======
	case STRINGSLICE:
		return v.asStringSlice()
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	}
	return unknownValueType{}
}

// Emit returns a string representation of Value's data.
func (v Value) Emit() string {
	switch v.Type() {
<<<<<<< HEAD
	case ARRAY:
		return fmt.Sprint(v.array)
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case BOOLSLICE:
		return fmt.Sprint(*(v.slice.(*[]bool)))
=======
	case BOOLSLICE:
		return fmt.Sprint(v.asBoolSlice())
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case BOOL:
		return strconv.FormatBool(v.AsBool())
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case INT64SLICE:
		return fmt.Sprint(*(v.slice.(*[]int64)))
=======
	case INT64SLICE:
		return fmt.Sprint(v.asInt64Slice())
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case INT64:
		return strconv.FormatInt(v.AsInt64(), 10)
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case FLOAT64SLICE:
		return fmt.Sprint(*(v.slice.(*[]float64)))
=======
	case FLOAT64SLICE:
		return fmt.Sprint(v.asFloat64Slice())
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case FLOAT64:
		return fmt.Sprint(v.AsFloat64())
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case STRINGSLICE:
		return fmt.Sprint(*(v.slice.(*[]string)))
=======
	case STRINGSLICE:
		return fmt.Sprint(v.asStringSlice())
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	case STRING:
		return v.stringly
	default:
		return "unknown"
	}
}

// MarshalJSON returns the JSON encoding of the Value.
func (v Value) MarshalJSON() ([]byte, error) {
	var jsonVal struct {
		Type  string
		Value interface{}
	}
	jsonVal.Type = v.Type().String()
	jsonVal.Value = v.AsInterface()
	return json.Marshal(jsonVal)
}
