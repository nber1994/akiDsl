package compile

import (
	"fmt"
	"github.com/spf13/cast"
)

func BAdd(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l + cast.ToInt(r)
	case uint:
		return l + cast.ToUint(r)
	case int8:
		return l + cast.ToInt8(r)
	case int16:
		return l + cast.ToInt16(r)
	case int32:
		return l + cast.ToInt32(r)
	case int64:
		return l + cast.ToInt64(r)
	case uint8:
		return l + cast.ToUint8(r)
	case uint16:
		return l + cast.ToUint16(r)
	case uint32:
		return l + cast.ToUint32(r)
	case uint64:
		return l + cast.ToUint64(r)
	case float32:
		return l + cast.ToFloat32(r)
	case float64:
		return l + cast.ToFloat64(r)
	case string:
		return l + cast.ToString(r)

	default:
        panic(fmt.Sprintf("syntax error: bad binary add type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BSub(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l - cast.ToInt(r)
	case uint:
		return l - cast.ToUint(r)
	case int8:
		return l - cast.ToInt8(r)
	case int16:
		return l - cast.ToInt16(r)
	case int32:
		return l - cast.ToInt32(r)
	case int64:
		return l - cast.ToInt64(r)
	case uint8:
		return l - cast.ToUint8(r)
	case uint16:
		return l - cast.ToUint16(r)
	case uint32:
		return l - cast.ToUint32(r)
	case uint64:
		return l - cast.ToUint64(r)
	case float32:
		return l - cast.ToFloat32(r)
	case float64:
		return l - cast.ToFloat64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BMul(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l * cast.ToInt(r)
	case uint:
		return l * cast.ToUint(r)
	case int8:
		return l * cast.ToInt8(r)
	case int16:
		return l * cast.ToInt16(r)
	case int32:
		return l * cast.ToInt32(r)
	case int64:
		return l * cast.ToInt64(r)
	case uint8:
		return l * cast.ToUint8(r)
	case uint16:
		return l * cast.ToUint16(r)
	case uint32:
		return l * cast.ToUint32(r)
	case uint64:
		return l * cast.ToUint64(r)
	case float32:
		return l * cast.ToFloat32(r)
	case float64:
		return l * cast.ToFloat64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BQuo(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l / cast.ToInt(r)
	case uint:
		return l / cast.ToUint(r)
	case int8:
		return l / cast.ToInt8(r)
	case int16:
		return l / cast.ToInt16(r)
	case int32:
		return l / cast.ToInt32(r)
	case int64:
		return l / cast.ToInt64(r)
	case uint8:
		return l / cast.ToUint8(r)
	case uint16:
		return l / cast.ToUint16(r)
	case uint32:
		return l / cast.ToUint32(r)
	case uint64:
		return l / cast.ToUint64(r)
	case float32:
		return l / cast.ToFloat32(r)
	case float64:
		return l / cast.ToFloat64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BRem(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l % cast.ToInt(r)
	case uint:
		return l % cast.ToUint(r)
	case int8:
		return l % cast.ToInt8(r)
	case int16:
		return l % cast.ToInt16(r)
	case int32:
		return l % cast.ToInt32(r)
	case int64:
		return l % cast.ToInt64(r)
	case uint8:
		return l % cast.ToUint8(r)
	case uint16:
		return l % cast.ToUint16(r)
	case uint32:
		return l % cast.ToUint32(r)
	case uint64:
		return l % cast.ToUint64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BAnd(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l & cast.ToInt(r)
	case uint:
		return l & cast.ToUint(r)
	case int8:
		return l & cast.ToInt8(r)
	case int16:
		return l & cast.ToInt16(r)
	case int32:
		return l & cast.ToInt32(r)
	case int64:
		return l & cast.ToInt64(r)
	case uint8:
		return l & cast.ToUint8(r)
	case uint16:
		return l & cast.ToUint16(r)
	case uint32:
		return l & cast.ToUint32(r)
	case uint64:
		return l & cast.ToUint64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BOr(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l | cast.ToInt(r)
	case uint:
		return l | cast.ToUint(r)
	case int8:
		return l | cast.ToInt8(r)
	case int16:
		return l | cast.ToInt16(r)
	case int32:
		return l | cast.ToInt32(r)
	case int64:
		return l | cast.ToInt64(r)
	case uint8:
		return l | cast.ToUint8(r)
	case uint16:
		return l | cast.ToUint16(r)
	case uint32:
		return l | cast.ToUint32(r)
	case uint64:
		return l | cast.ToUint64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BXor(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case int:
		return l ^ cast.ToInt(r)
	case uint:
		return l ^ cast.ToUint(r)
	case int8:
		return l ^ cast.ToInt8(r)
	case int16:
		return l ^ cast.ToInt16(r)
	case int32:
		return l ^ cast.ToInt32(r)
	case int64:
		return l ^ cast.ToInt64(r)
	case uint8:
		return l ^ cast.ToUint8(r)
	case uint16:
		return l ^ cast.ToUint16(r)
	case uint32:
		return l ^ cast.ToUint32(r)
	case uint64:
		return l ^ cast.ToUint64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BShl(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case uint:
		return l << cast.ToUint(r)
	case uint8:
		return l << cast.ToUint8(r)
	case uint16:
		return l << cast.ToUint16(r)
	case uint32:
		return l << cast.ToUint32(r)
	case uint64:
		return l << cast.ToUint64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BShr(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case uint:
		return l >> cast.ToUint(r)
	case uint8:
		return l >> cast.ToUint8(r)
	case uint16:
		return l >> cast.ToUint16(r)
	case uint32:
		return l >> cast.ToUint32(r)
	case uint64:
		return l >> cast.ToUint64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BLss(l interface{}, r interface{}) bool {
	switch l := l.(type) {
	case int:
		return l < cast.ToInt(r)
	case uint:
		return l < cast.ToUint(r)
	case int8:
		return l < cast.ToInt8(r)
	case int16:
		return l < cast.ToInt16(r)
	case int32:
		return l < cast.ToInt32(r)
	case int64:
		return l < cast.ToInt64(r)
	case uint8:
		return l < cast.ToUint8(r)
	case uint16:
		return l < cast.ToUint16(r)
	case uint32:
		return l < cast.ToUint32(r)
	case uint64:
		return l < cast.ToUint64(r)
	case float32:
		return l < cast.ToFloat32(r)
	case float64:
		return l < cast.ToFloat64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret bool
    return ret
}

func BGtr(l interface{}, r interface{}) bool {
	switch l := l.(type) {
	case int:
		return l > cast.ToInt(r)
	case uint:
		return l > cast.ToUint(r)
	case int8:
		return l > cast.ToInt8(r)
	case int16:
		return l > cast.ToInt16(r)
	case int32:
		return l > cast.ToInt32(r)
	case int64:
		return l > cast.ToInt64(r)
	case uint8:
		return l > cast.ToUint8(r)
	case uint16:
		return l > cast.ToUint16(r)
	case uint32:
		return l > cast.ToUint32(r)
	case uint64:
		return l > cast.ToUint64(r)
	case float32:
		return l > cast.ToFloat32(r)
	case float64:
		return l > cast.ToFloat64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret bool
    return ret
}

func BLeq(l interface{}, r interface{}) bool {
	switch l := l.(type) {
	case int:
		return l <= cast.ToInt(r)
	case uint:
		return l <= cast.ToUint(r)
	case int8:
		return l <= cast.ToInt8(r)
	case int16:
		return l <= cast.ToInt16(r)
	case int32:
		return l <= cast.ToInt32(r)
	case int64:
		return l <= cast.ToInt64(r)
	case uint8:
		return l <= cast.ToUint8(r)
	case uint16:
		return l <= cast.ToUint16(r)
	case uint32:
		return l <= cast.ToUint32(r)
	case uint64:
		return l <= cast.ToUint64(r)
	case float32:
		return l <= cast.ToFloat32(r)
	case float64:
		return l <= cast.ToFloat64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret bool
    return ret
}

func BGeq(l interface{}, r interface{}) bool {
	switch l := l.(type) {
	case int:
		return l >= cast.ToInt(r)
	case uint:
		return l >= cast.ToUint(r)
	case int8:
		return l >= cast.ToInt8(r)
	case int16:
		return l >= cast.ToInt16(r)
	case int32:
		return l >= cast.ToInt32(r)
	case int64:
		return l >= cast.ToInt64(r)
	case uint8:
		return l >= cast.ToUint8(r)
	case uint16:
		return l >= cast.ToUint16(r)
	case uint32:
		return l >= cast.ToUint32(r)
	case uint64:
		return l >= cast.ToUint64(r)
	case float32:
		return l >= cast.ToFloat32(r)
	case float64:
		return l >= cast.ToFloat64(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret bool
    return ret
}

func BNeq(l interface{}, r interface{}) bool {
    if l == nil && r == nil {
        return false
    } else if  l != nil && r == nil {
        return true
    } else if l == nil && r != nil {
        return true
    } else if l != nil && r != nil {
        return l != r
    }
    return false
}

func BEql(l interface{}, r interface{}) bool {
    if l == nil && r == nil {
        return true
    } else if  l != nil && r == nil {
        return false
    } else if l == nil && r != nil {
        return false
    } else if l != nil && r != nil {
        return l == r
    }
    return false
}

func BInc(r interface{}) interface{} {
	switch r := r.(type) {
	case int:
		return  cast.ToInt(r) + cast.ToInt(1)
	case uint:
		return  cast.ToUint(r) + cast.ToUint(1)
	case int8:
		return  cast.ToInt8(r) + cast.ToInt8(1)
	case int16:
		return  cast.ToInt16(r) + cast.ToInt16(1)
	case int32:
		return  cast.ToInt32(r) + cast.ToInt32(1)
	case int64:
		return  cast.ToInt64(r) + cast.ToInt64(1)
	case uint8:
		return  cast.ToUint8(r) + cast.ToUint8(1)
	case uint16:
		return  cast.ToUint16(r) + cast.ToUint16(1)
	case uint32:
		return  cast.ToUint32(r) + cast.ToUint32(1)
	case uint64:
		return  cast.ToUint64(r) + cast.ToUint64(1)
	case float32:
		return  cast.ToFloat32(r) + cast.ToFloat32(1)
	case float64:
		return  cast.ToFloat64(r) + cast.ToFloat64(1)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", r))
	}
    var ret interface{}
    return ret
}

func BDec(r interface{}) interface{} {
	switch r := r.(type) {
	case int:
		return  cast.ToInt(r) - cast.ToInt(1)
	case uint:
		return  cast.ToUint(r) - cast.ToUint(1)
	case int8:
		return  cast.ToInt8(r) - cast.ToInt8(1)
	case int16:
		return  cast.ToInt16(r) - cast.ToInt16(1)
	case int32:
		return  cast.ToInt32(r) - cast.ToInt32(1)
	case int64:
		return  cast.ToInt64(r) - cast.ToInt64(1)
	case uint8:
		return  cast.ToUint8(r) - cast.ToUint8(1)
	case uint16:
		return  cast.ToUint16(r) - cast.ToUint16(1)
	case uint32:
		return  cast.ToUint32(r) - cast.ToUint32(1)
	case uint64:
		return  cast.ToUint64(r) - cast.ToUint64(1)
	case float32:
		return  cast.ToFloat32(r) - cast.ToFloat32(1)
	case float64:
		return  cast.ToFloat64(r) - cast.ToFloat64(1)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", r))
	}
    var ret interface{}
    return ret

}


func BLand(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case bool:
		return l && cast.ToBool(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}

func BLor(l interface{}, r interface{}) interface{} {
	switch l := l.(type) {
	case bool:
		return l || cast.ToBool(r)
	default:
        panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", l))
	}
    var ret interface{}
    return ret
}
