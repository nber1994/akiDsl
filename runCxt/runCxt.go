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
    rcxt := make(map[string]interface{})
    rcxt["nil"] = nil
    return &RunCxt{
        Vars: rcxt,
    }
}

//获取值
func (this *RunCxt) GetValue(varName string) interface{}{
    //fmt.Println("[+]get var ", varName)
    if _, exist := this.Vars[varName]; !exist {
        panic("syntax error: not exist var")
    }
    //fmt.Println("[+]now var ", this.ToString())
    return this.Vars[varName]
}

func (this *RunCxt) ValueExist(varName string) bool {
    _, exist := this.Vars[varName]
    return exist
}

//设置值
func (this *RunCxt) SetValue(varName string, value interface{}) bool {
    //fmt.Println("[+]set var ", varName, value)
    this.Vars[varName] = value
    //fmt.Println("[+]now var ", this.ToString())
    return true
}

func (this *RunCxt) ToString() string {
    jsonStu, err := json.Marshal(this.Vars)
    if err != nil {
        //fmt.Println("[-]err to json string")
    }
    return string(jsonStu)
}
