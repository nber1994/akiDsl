package runCxt

import (
    "fmt"
    "encoding/json"
)


type RunCxt struct {
    Vars map[string]interface{} //这块模拟内存
    Name string
}

func NewRunCxt() *RunCxt{
    return &RunCxt{}
}

//获取值
func (this *RunCxt) GetValue(varName string) interface{}{
    fmt.Println("+++ get var ", varName)
    fmt.Println("+++ now var ", this.ToString())
    if _, exist := this.Vars[varName]; !exist {
        panic("syntax error: not exist var")
    }
    return this.Vars[varName]
}

//设置值
func (this *RunCxt) SetValue(varName string, value interface{}) bool {
    fmt.Println("+++ set var ", varName, value)
    fmt.Println("+++ now var ", this.ToString())
    this.Vars[varName] = value
    return true
}

func (this *RunCxt) ToString() string {
    jsonStr, err := json.Marshal(this.Vars)
    if err != nil {
        panic("run time cxt string error")
    }
    return string(jsonStr)
}
