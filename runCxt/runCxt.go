package runCxt

import (
    "fmt"
)


type RunCxt struct {
    Vars map[string]interface{} //这块模拟内存
    Name string
}

func NewRunCxt() *RunCxt{
    return &RunCxt{
        Vars: make(map[string]interface{}),
    }
}

//获取值
func (this *RunCxt) GetValue(varName string) interface{}{
    if _, exist := this.Vars[varName]; !exist {
        panic("syntax error: not exist var")
    }
    fmt.Println("+++ get var ", varName)
    print_json(this.Vars)
    return this.Vars[varName]
}

//设置值
func (this *RunCxt) SetValue(varName string, value interface{}) bool {
    this.Vars[varName] = value
    fmt.Println("+++ set var ", varName, value)
    print_json(this.Vars)
    return true
}

func print_json(m map[string]interface{}) {
    for k, v := range m {
        switch vv := v.(type) {
        case string:
            fmt.Println(k, "is string", vv)
        case float64:
            fmt.Println(k, "is float", int64(vv))
        case int:
            fmt.Println(k, "is int", vv)
        case []interface{}:
            fmt.Println(k, "is an array:")
            for i, u := range vv {
                fmt.Println(i, u)
            }
        case nil:
            fmt.Println(k, "is nil", "null")
        case map[string]interface{}:
            fmt.Println(k, "is an map:")
            print_json(vv)
        default:
            fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
        }
    }
}
