package runCxt


type RunCxt struct {
    Vars map[string]interface{} //这块模拟内存的堆
    Name string
}

func NewRunCxt() *RunCxt{
    return &RunCxt{}
}

//获取值
func (this *RunCxt) GetValue(varName string) (interface{}, bool){
    if _, exist := this.Vars[varName]; exist {
        return this.Vars[varName], true
    }
    return nil, false
}

//设置值
func (this *RunCxt) SetValue(varName string, value interface{}) bool {
    this.Vars[varName] = value
    return true
}
