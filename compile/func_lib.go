package compile

import (
    "fmt"
    "time"
)

type FuncLib struct {

}

func NewFuncLib() *FuncLib {
    return &FuncLib{}
}

func (this *FuncLib) Append(target []interface{}, item interface{}) interface{} {
    return append(target, item)
}

func (this *FuncLib) Len(target []interface{}) int {
    return len(target)
}

func (this *FuncLib) Println(target ...interface{}) {
    //fmt.Println(target...)
}

func (this *FuncLib) Sprintf(spf string, v ...interface{}) string {
    //fmt.Println(">>sprintf ", spf, v)
    return fmt.Sprintf(spf, v...)
}

func (this *FuncLib) TimeParse(spf, timeStr string) string {
    ret, err := time.Parse(spf, timeStr)
    if nil != err {
        panic(err)
    }
    return ret.String()
}
