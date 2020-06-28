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
    return &RunCxt{
        Vars: make(map[string]interface{}),
    }
}

//获取值
func (this *RunCxt) GetValue(varName string) interface{}{
    if _, exist := this.Vars[varName]; !exist {
        panic("syntax error: not exist var")
    }
    fmt.Println("[+]get var ", varName)
    fmt.Println("[+]now var ", this.ToString())
    return this.Vars[varName]
}

//设置值
func (this *RunCxt) SetValue(varName string, value interface{}) bool {
    this.Vars[varName] = value
    fmt.Println("[+]set var ", varName, value)
    fmt.Println("[+]now var ", this.ToString())
    return true
}

func (this *RunCxt) ToString() string {
    jsonStu, err := json.Marshal(this.Vars)
    if err != nil {
        fmt.Println("[-]err to json string")
    }
    return string(jsonStu)
}
