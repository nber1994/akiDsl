package nodejson

import (
	"encoding/json"
	"errors"
	"reflect"
)

func (n Node) Int() int {
	v, _ := n.IntE()
	return v
}

func (n Node) Int64() int64 {
	v, _ := n.Int64E()
	return v
}

func (n Node) String() string {
	v, _ := n.StringE()
	return v
}

func (n Node) Bool() bool {
	v, _ := n.BoolE()
	return v
}

func (n Node) Float64() float64 {
	v, _ := n.Float64E()
	return v
}

func (n Node) Interface() interface{} {
	return n.data
}

func (n *Node) IntE() (int, error) {
    var nilRet int
	switch n.data.(type) {
	case json.Number:
		i, err := n.data.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(n.data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(n.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(n.data).Uint()), nil
	}
	return nilRet, errors.New("invalid value type")
}

func (n *Node) Float64E() (float64, error) {
    var nilRet float64
	switch n.data.(type) {
	case json.Number:
		return n.data.(json.Number).Float64()
	case float32, float64:
		return reflect.ValueOf(n.data).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(n.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(n.data).Uint()), nil
	}
	return nilRet, errors.New("invalid value type")
}

func (n *Node) Int64E() (int64, error) {
    var nilRet int64
	switch n.data.(type) {
	case json.Number:
		return n.data.(json.Number).Int64()
	case float32, float64:
		return int64(reflect.ValueOf(n.data).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(n.data).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(n.data).Uint()), nil
	}
	return nilRet, errors.New("invalid value type")
}

func (n *Node) StringE() (string, error) {
    var nilRet string
	if s, ok := (n.data).(string); ok {
		return s, nil
	}
	return nilRet, errors.New("type assertion to string failed")

}

func (n *Node) BoolE() (bool, error) {
    var nilRet bool
	if s, ok := (n.data).(bool); ok {
		return s, nil
	}
	return nilRet, errors.New("type assertion to bool failed")
}

