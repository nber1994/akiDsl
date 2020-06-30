package compile

import (
    "reflect"
    "fmt"
)

func ReflectValueRealTypeValue(v reflect.Value) interface{} {
    fmt.Println(v.Kind())

    switch v.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return v.Int()
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        return v.Uint()
    case reflect.Float32, reflect.Float64:
        return v.Float()
    case reflect.Bool:
        return v.Bool()
    case reflect.String:
        return v.String()
    default:
        panic("syntax error: nonsupport return value")
    }
}
